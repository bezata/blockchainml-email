package r2

import (
    "context"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
    client     *s3.Client
    bucketName string
}

func NewClient(endpoint, bucketName, accessKey, secretKey string) (*Client, error) {
    cfg := s3.Options{
        BaseEndpoint: &endpoint,
        Region:      "auto",
        Credentials: NewStaticCredentialsProvider(accessKey, secretKey),
    }
    
    client := s3.New(cfg)
    
    return &Client{
        client:     client,
        bucketName: bucketName,
    }, nil
}

func (c *Client) Upload(ctx context.Context, key string, data []byte, contentType string) error {
    // Implement upload logic
    return nil
}

func (c *Client) Download(ctx context.Context, key string) ([]byte, error) {
    // Implement download logic
    return nil, nil
}
