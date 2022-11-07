package mk

import (
	"os"
	"path/filepath"
)

// Type File is a building block for scaffolding a file.
//
// These parameters can be set to customize the scaffolding process:
//   - Name specifies the file name.
//   - Content specifies the file content in []byte.
type File struct {
	Name string

	Content []byte
}

// Scaffold scaffolds a file on the top of basePath.
func (file File) Scaffold(basePath string) error {
	return os.WriteFile(filepath.Join(basePath, file.Name), file.Content, os.ModePerm)
}
