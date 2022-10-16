package mk

import (
	"os"
	"path/filepath"
)

type File struct {
	Name string
}

func (file File) Scaffold(basePath string) error {
	return os.WriteFile(filepath.Join(basePath, file.Name), []byte{}, 0600)
}
