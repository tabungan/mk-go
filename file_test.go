package mk

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile_Scaffold(t *testing.T) {
	basePath := t.TempDir()

	t.Run("creates file", func(t *testing.T) {
		ScaffoldOrPanic(basePath, File{
			Name: "file",

			Content: []byte("content"),
		})

		path := filepath.Join(basePath, "file")

		t.Run("with name as specified", func(t *testing.T) {
			assert.FileExists(t, path)
		})

		t.Run("with content as specified", func(t *testing.T) {
			if content, err := os.ReadFile(path); assert.NoError(t, err) {
				assert.Equal(t, []byte("content"), content)
			}
		})
	})

	t.Run("fails if cannot create file", func(t *testing.T) {
		ScaffoldOrPanic(basePath, Dir{Name: "dir"})

		assert.Error(t, Scaffold(basePath, File{Name: "dir"}))
	})
}
