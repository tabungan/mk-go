package mk

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile_Scaffold(t *testing.T) {
	basePath := t.TempDir()

	t.Run("creates file", func(t *testing.T) {
		ScaffoldingOrPanic(basePath, File{
			Name: "file",
		})

		t.Run("with name as specified", func(t *testing.T) {
			assert.FileExists(t, filepath.Join(basePath, "file"))
		})
	})

	t.Run("fails if cannot create file", func(t *testing.T) {
		ScaffoldingOrPanic(basePath, Dir{
			Name: "dir",
		})

		assert.Error(t, Scaffolding(basePath, File{
			Name: "dir",
		}))
	})
}
