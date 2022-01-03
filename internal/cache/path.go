package cache

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/titolins/jira_scrapper/config"
)

func pathInfo(path string) (os.FileInfo, bool) {
	fi, err := os.Stat(path)
	return fi, !os.IsNotExist(err)
}

func savePath(p string) string {
	p = path.Join(config.Path(), p)

	dir := path.Dir(p)
	if _, exists := pathInfo(dir); !exists {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("failed to create dir %q: %v", dir, err)
		}
	}
	return p
}

func cachePath(p, f string) string {
	f = fmt.Sprintf("%s.json", f)
	p = path.Join("cache", p, f)
	return savePath(p)
}
