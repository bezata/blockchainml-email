package cache

// Cache defines the interface for caching operations
type Cache interface {
    Get(key string) (interface{}, error)
    Set(key string, value interface{}) error
    Delete(key string) error
}
