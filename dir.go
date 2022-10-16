package mk

import (
	"os"
	"path/filepath"
)

type Dir struct {
	Name string
}

func (dir Dir) Scaffold(basePath string) error {
	os.Mkdir(filepath.Join(basePath, dir.Name), 0700)
	return nil
}
