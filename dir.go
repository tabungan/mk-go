package mk

import (
	"os"
	"path/filepath"
)

type Dir struct {
	Name string
}

func (dir Dir) Scaffold(basePath string) error {
	return os.Mkdir(filepath.Join(basePath, dir.Name), 0700)
}
