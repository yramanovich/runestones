package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Filesystem struct {
}

func (fs *Filesystem) SaveRunestone(ctx context.Context, key string, content []byte) error {
	f, err := os.Create(generatePath(key))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, bytes.NewBuffer(content))
	return err
}

func (fs *Filesystem) FindRunestone(ctx context.Context, key string) ([]byte, error) {
	return os.ReadFile(generatePath(key))
}

func generatePath(key string) string {
	return filepath.Join(os.TempDir(), fmt.Sprintf("runestone-%s.txt", key))
}
