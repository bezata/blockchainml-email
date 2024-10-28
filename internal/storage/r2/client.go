package r2

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/zap"
)

type Client struct {
	client     *s3.Client
	bucketName string
	logger     *zap.Logger
}

type Config struct {
	AccountID  string
	AccessKey  string
	SecretKey  string
	BucketName string
	Endpoint   string
}

func NewClient(cfg Config, logger *zap.Logger) (*Client, error) {
	creds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
		cfg.AccessKey,
		cfg.SecretKey,
		"",
	))

	client := s3.New(s3.Options{
		Credentials: creds,
		Region:      "auto",
		BaseEndpoint: aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", cfg.AccountID)),
	})

	return &Client{
		client:     client,
		bucketName: cfg.BucketName,
		logger:     logger,
	}, nil
}

// Upload uploads data to R2
func (c *Client) Upload(ctx context.Context, key string, data []byte, contentType string) error {
	input := &s3.PutObjectInput{
		Bucket:      aws.String(c.bucketName),
		Key:         aws.String(key),
		Body:        io.NopCloser(bytes.NewReader(data)),
		ContentType: aws.String(contentType),
	}

	_, err := c.client.PutObject(ctx, input)
	if err != nil {
		c.logger.Error("failed to upload to R2",
			zap.String("key", key),
			zap.Error(err),
		)
		return fmt.Errorf("failed to upload to R2: %w", err)
	}

	return nil
}

// Download downloads data from R2
func (c *Client) Download(ctx context.Context, key string) ([]byte, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(key),
	}

	result, err := c.client.GetObject(ctx, input)
	if err != nil {
		c.logger.Error("failed to download from R2",
			zap.String("key", key),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to download from R2: %w", err)
	}
	defer result.Body.Close()

	return io.ReadAll(result.Body)
}

// Delete deletes an object from R2
func (c *Client) Delete(ctx context.Context, key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(key),
	}

	_, err := c.client.DeleteObject(ctx, input)
	if err != nil {
		c.logger.Error("failed to delete from R2",
			zap.String("key", key),
			zap.Error(err),
		)
		return fmt.Errorf("failed to delete from R2: %w", err)
	}

	return nil
}

// ListObjects lists objects in R2 with a prefix
func (c *Client) ListObjects(ctx context.Context, prefix string) ([]string, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(c.bucketName),
		Prefix: aws.String(prefix),
	}

	var keys []string
	paginator := s3.NewListObjectsV2Paginator(c.client, input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			c.logger.Error("failed to list R2 objects",
				zap.String("prefix", prefix),
				zap.Error(err),
			)
			return nil, fmt.Errorf("failed to list R2 objects: %w", err)
		}

		for _, obj := range page.Contents {
			keys = append(keys, *obj.Key)
		}
	}

	return keys, nil
}
