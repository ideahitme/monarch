package storage

// Storage defines simple key-value storage
type Storage interface {
	SetValue(key, value string) error
	ReadKey(key string) (string, error)
	Keys() ([]string, error)
}
