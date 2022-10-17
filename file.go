package mk

import (
	"os"
	"path/filepath"
)

type File struct {
	Name string

	Content []byte
}

func (file File) Scaffold(basePath string) error {
	return os.WriteFile(filepath.Join(basePath, file.Name), file.Content, 0600)
}
