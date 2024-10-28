package config

type SecurityConfig struct {
	EncryptionKey string                   `json:"encryptionKey"`
	RateLimit     RateLimitConfig          `json:"rateLimit"`
	JWT           JWTConfig                `json:"jwt"`
	Cloudflare    CloudflareSecurityConfig `json:"cloudflare"`
}

type RateLimitConfig struct {
	RequestsPerMinute int `json:"requestsPerMinute"`
	BurstSize         int `json:"burstSize"`
}

type CloudflareSecurityConfig struct {
	AccessAUD  string   `json:"accessAud"`
	TeamDomain string   `json:"teamDomain"`
	AllowedIPs []string `json:"allowedIps"`
}
