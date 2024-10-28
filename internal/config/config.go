package config

import (
	"encoding/json"
	"os"
	"time"
)

type Config struct {
	Server     ServerConfig     `json:"server"`
	MongoDB    MongoDBConfig    `json:"mongodb"`
	Redis      RedisConfig      `json:"redis"`
	R2         R2Config         `json:"r2"`
	JWT        JWTConfig        `json:"jwt"`
	Monitoring MonitoringConfig `json:"monitoring"`
	Search     SearchConfig     `json:"search"`
	Cache      CacheConfig      `json:"cache"`
	Realtime   RealtimeConfig   `json:"realtime"`
	Cloudflare CloudflareConfig `json:"cloudflare"`
}

type ServerConfig struct {
	Port            int    `json:"port"`
	Host            string `json:"host"`
	ReadTimeout     int    `json:"readTimeout"`
	WriteTimeout    int    `json:"writeTimeout"`
	MaxRequestSize  int64  `json:"maxRequestSize"`
}

type MongoDBConfig struct {
	URI             string `json:"uri"`
	Database        string `json:"database"`
	MaxPoolSize     uint64 `json:"maxPoolSize"`
	MinPoolSize     uint64 `json:"minPoolSize"`
	MaxConnIdleTime int    `json:"maxConnIdleTime"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type R2Config struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
	Bucket          string `json:"bucket"`
	Region          string `json:"region"`
}

type PrometheusConfig struct {
	Port int    `json:"port"`
	Path string `json:"path"`
}

type TracingConfig struct {
	Enabled bool   `json:"enabled"`
	Endpoint string `json:"endpoint"`
}

type LoggingConfig struct {
	Level string `json:"level"`
	Format string `json:"format"`
}

type MonitoringConfig struct {
	Prometheus PrometheusConfig `json:"prometheus"`
	Tracing    TracingConfig    `json:"tracing"`
	Logging    LoggingConfig    `json:"logging"`
}

type JWTConfig struct {
	Secret    string `json:"secret"`
	ExpiresIn int    `json:"expiresIn"` // in minutes
}

type SearchConfig struct {
	ElasticsearchURLs []string `json:"elasticsearchUrls"`
	Username          string   `json:"username"`
	Password          string   `json:"password"`
	IndexPrefix       string   `json:"indexPrefix"`
}

type CacheConfig struct {
	DefaultTTL  time.Duration `json:"defaultTTL"`
	MaxEntries  int64         `json:"maxEntries"`
	MaxMemory   string        `json:"maxMemory"`
}

type RealtimeConfig struct {
	EnableWebSocket bool   `json:"enableWebSocket"`
	PingInterval   int    `json:"pingInterval"`
	WriteTimeout   int    `json:"writeTimeout"`
	ReadTimeout    int    `json:"readTimeout"`
}

type CloudflareConfig struct {
	AccountID      string `json:"accountId"`
	ZoneID        string `json:"zoneId"`
	APIToken      string `json:"apiToken"`
	WorkersDomain string `json:"workersDomain"`
	R2Bucket      string `json:"r2Bucket"`
}

type SecurityConfig struct {
	CloudflareAccess struct {
		AUD         string   `json:"aud"`
		TeamDomain string   `json:"teamDomain"`
		AllowedIPs []string `json:"allowedIps"`
	} `json:"cloudflareAccess"`
}

// LoadConfig loads config from file and environment variables
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	// Override with environment variables
	if mongoURI := os.Getenv("MONGODB_URI"); mongoURI != "" {
		config.MongoDB.URI = mongoURI
	}

	return &config, nil
}
