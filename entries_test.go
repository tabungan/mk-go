package mk

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntries_Scaffold(t *testing.T) {
	basePath := t.TempDir()

	t.Run("scaffolds entries", func(t *testing.T) {
		ScaffoldOrPanic(basePath, Entries{
			Dir{Name: "dir"},
			File{Name: "file"},
		})

		t.Run("given 1st entry", func(t *testing.T) {
			assert.DirExists(t, filepath.Join(basePath, "dir"))
		})

		t.Run("given n-entries", func(t *testing.T) {
			assert.FileExists(t, filepath.Join(basePath, "file"))
		})
	})

	t.Run("fails if one entry cannot be scaffolded", func(t *testing.T) {
		ScaffoldOrPanic(basePath, File{Name: "file"})

		assert.Error(t, Scaffolding(basePath, Entries{
			Dir{Name: "file"},
		}))
	})
}
