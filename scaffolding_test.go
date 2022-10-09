package mk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockEntry struct {
	basePathInput string
}

func (entry *mockEntry) Scaffold(basePath string) error {
	entry.basePathInput = basePath
	return nil
}

func TestScaffolding(t *testing.T) {
	basePath := t.TempDir()

	t.Run("call scaffold and pass base path", func(t *testing.T) {
		var entry mockEntry
		Scaffolding(basePath, &entry)

		assert.Equal(t, basePath, entry.basePathInput)
	})
}
