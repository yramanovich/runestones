package contents

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func NewFilesystem(dir string) *Filesystem {
	if dir == "" {
		dir = os.TempDir()
	}
	return &Filesystem{dir: dir}
}

type Filesystem struct {
	dir string
}

func (fs *Filesystem) SaveContent(_ context.Context, key string, content []byte) error {
	f, err := os.Create(fs.generatePath(key))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, bytes.NewBuffer(content))
	return err
}

func (fs *Filesystem) FindContent(_ context.Context, key string) ([]byte, error) {
	return os.ReadFile(fs.generatePath(key))
}

func (fs *Filesystem) generatePath(key string) string {
	return filepath.Join(fs.dir, fmt.Sprintf("runestone-%s.txt", key))
}
