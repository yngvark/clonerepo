// Package storage can store temporary stuff.
package store

import (
	"io/ioutil"

	"github.com/spf13/afero"
)

// TemporaryStorage wraps a filesystem.
type TemporaryStorage struct {
	BasePath string
	Fs       afero.Fs
}

// NewTemporaryStorage creates a new temporary storage.
func NewTemporaryStorage() (*TemporaryStorage, error) {
	dir, err := ioutil.TempDir("", "gclone")
	if err != nil {
		return nil, err
	}

	return &TemporaryStorage{
		BasePath: dir,
		Fs:       afero.NewBasePathFs(afero.NewOsFs(), dir),
	}, nil
}
