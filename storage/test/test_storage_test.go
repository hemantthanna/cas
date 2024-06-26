package storagetest

import (
	"testing"

	"github.com/hemantthanna/cas/storage"
)

func TestMemory(t *testing.T) {
	RunTests(t, func(_ testing.TB) (storage.Storage, func()) {
		return storage.NewInMemory(), func() {}
	})
}
