package mk

import (
	"os"
	"path/filepath"
)

type Dir struct {
	Name string

	Content Entry
}

func (dir Dir) Scaffold(basePath string) error {
	path := filepath.Join(basePath, dir.Name)

	err := os.Mkdir(path, 0700)
	if err != nil {
		return err
	}

	if dir.Content != nil {
		dir.Content.Scaffold(path)
	}
	return nil
}
