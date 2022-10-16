package mk

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntries_Scaffold(t *testing.T) {
	basePath := t.TempDir()

	t.Run("scaffolds entries", func(t *testing.T) {
		ScaffoldingOrPanic(basePath, Entries{
			Dir{Name: "dir"},
		})

		t.Run("given 1st entry", func(t *testing.T) {
			assert.DirExists(t, filepath.Join(basePath, "dir"))
		})
	})
}
