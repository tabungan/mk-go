package mk

import (
	"os"
	"path/filepath"
)

type File struct {
	Name string
}

func (file File) Scaffold(basePath string) error {
	os.WriteFile(filepath.Join(basePath, file.Name), []byte{}, 0600)
	return nil
}
