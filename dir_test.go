package mk

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDir_Scaffold(t *testing.T) {
	basePath := t.TempDir()

	t.Run("creates directory", func(t *testing.T) {
		ScaffoldOrPanic(basePath, Dir{
			Name: "dir",

			Content: File{Name: "file"},
		})

		t.Run("with name as specified", func(t *testing.T) {
			assert.DirExists(t, filepath.Join(basePath, "dir"))
		})

		t.Run("with content as specified", func(t *testing.T) {
			assert.FileExists(t, filepath.Join(basePath, "dir", "file"))
		})
	})

	t.Run("fails if cannot create directory", func(t *testing.T) {
		ScaffoldOrPanic(basePath, File{Name: "file"})

		assert.Error(t, Scaffolding(basePath, Dir{Name: "file"}))
	})
}
