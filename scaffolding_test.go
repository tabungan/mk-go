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

	t.Run("calls scaffold and pass base path", func(t *testing.T) {
		var entry mockEntry
		Scaffolding(basePath, &entry)

		assert.Equal(t, basePath, entry.basePathInput)
	})

	t.Run("forwards error from scaffold", func(t *testing.T) {
		entry := mockEntry{errOutput: os.ErrExist}
		assert.Error(t, Scaffolding(basePath, &entry), os.ErrExist)
	})
}

func TestScaffoldOrPanic(t *testing.T) {
	basePath := t.TempDir()

	t.Run("calls scaffold and pass base path", func(t *testing.T) {
		var entry mockEntry
		ScaffoldOrPanic(basePath, &entry)

		assert.Equal(t, basePath, entry.basePathInput)
	})

	t.Run("panics if got error from scaffold", func(t *testing.T) {
		entry := mockEntry{errOutput: os.ErrExist}
		assert.PanicsWithError(t, os.ErrExist.Error(), func() {
			ScaffoldOrPanic(basePath, &entry)
		})
	})
}
