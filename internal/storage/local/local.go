package local

import (
	"fmt"
	"net/url"
	"path"

	"codeberg.org/gruf/go-store/kv"

	"github.com/superseriousbusiness/gotosocial/internal/config"
)

type localStorage struct {
	*kv.KVStore
	basePath string
}

func New(c *config.StorageConfig) (*localStorage, error) {
	s, err := kv.OpenFile(c.Local.BasePath, nil)
	if err != nil {
		return nil, fmt.Errorf("error opening local storage backend: %s", err)
	}

	return &localStorage{
		KVStore:  s,
		basePath: c.Local.BasePath,
	}, nil
}

func (s *localStorage) URL(name string) (*url.URL, error) {
	return url.Parse(path.Join(s.basePath, name))
}
