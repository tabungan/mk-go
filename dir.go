package mk

import (
	"os"
	"path/filepath"
)

// Type Dir is a building block for scaffolding a directory.
//
// These parameters can be set to customize the scaffolding process:
//   - Name specifies the directory name.
//   - Content specifies the scaffolding to be scaffolded on the top of the directory.
type Dir struct {
	Name string

	Content Entry
}

// Scaffold scaffolds a directory on the top of basePath.
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
