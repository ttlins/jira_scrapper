package cache

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Cache struct {
	path string
}

func New(p, f string) *Cache {
	return &Cache{cachePath(p, f)}
}

func (c *Cache) Load(v interface{}) error {
	d, err := ioutil.ReadFile(c.path)
	if err != nil {
		return fmt.Errorf("failed to read cache file at %q: %v", c.Path(), err)
	}

	if err := json.Unmarshal(d, v); err != nil {
		return fmt.Errorf("failed to unmarshal cache at %q: %v", c.Path(), err)
	}

	return nil
}

func (c *Cache) Save(v interface{}) error {
	d, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal cache file %q: %v", c.Path(), err)
	}

	if err = ioutil.WriteFile(c.path, d, 0644); err != nil {
		return fmt.Errorf("failed to write cache file at %q: %v", c.Path(), err)
	}

	return nil
}

func (c *Cache) Path() string {
	return c.path
}

func (c *Cache) Exists() (os.FileInfo, bool) {
	return pathInfo(c.Path())
}
