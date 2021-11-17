package storage

import (
	"fmt"
	"net/url"

	"github.com/superseriousbusiness/gotosocial/internal/config"
	"github.com/superseriousbusiness/gotosocial/internal/storage/local"
	"github.com/superseriousbusiness/gotosocial/internal/storage/s3"
)

type Storage interface {
	Get(name string) ([]byte, error)
	Put(name string, data []byte) error
	Delete(name string) error

	URL(name string) (*url.URL, error)
}

func New(c *config.StorageConfig) (Storage, error) {
	if c.Backend == config.StorageBackendLocal {
		return local.New(c)
	} else if c.Backend == config.StorageBackendS3 {
		return s3.New(c)
	}

	return nil, fmt.Errorf("invalid storage backend type: %v", c.Backend)
}
