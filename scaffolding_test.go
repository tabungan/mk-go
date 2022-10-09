package mk

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockEntry struct {
	basePathInput string
	errOutput     error
}

func (entry *mockEntry) Scaffold(basePath string) error {
	entry.basePathInput = basePath
	return entry.errOutput
}

func TestScaffolding(t *testing.T) {
	basePath := t.TempDir()

	t.Run("call scaffold and pass base path", func(t *testing.T) {
		var entry mockEntry
		Scaffolding(basePath, &entry)

		assert.Equal(t, basePath, entry.basePathInput)
	})

	t.Run("forward error from scaffold", func(t *testing.T) {
		entry := mockEntry{errOutput: os.ErrExist}
		assert.Error(t, Scaffolding(basePath, &entry), os.ErrExist)
	})
}

func TestScaffoldingOrPanic(t *testing.T) {
	basePath := t.TempDir()

	t.Run("call scaffold and pass base path", func(t *testing.T) {
		var entry mockEntry
		ScaffoldingOrPanic(basePath, &entry)

		assert.Equal(t, basePath, entry.basePathInput)
	})
}
