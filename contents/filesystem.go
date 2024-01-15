package contents

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/yramanovich/runestones/log"
)

// NewFilesystem returns new Filesystem instance. Empty dir defaulted to os.TempDir().
func NewFilesystem(dir string) *Filesystem {
	if dir == "" {
		dir = os.TempDir()
	}
	return &Filesystem{dir: dir}
}

// Filesystem saves/reads content to/from filesystem.
type Filesystem struct {
	dir string
}

// SaveContent writes content to the file.
func (fs *Filesystem) SaveContent(ctx context.Context, key string, content []byte) error {
	f, err := os.Create(fs.generatePath(key))
	if err != nil {
		return err
	}
	defer log.Close(ctx, f, "close content file")

	_, err = io.Copy(f, bytes.NewBuffer(content))
	return err
}

// FindContent reads the file and returns the contents.
func (fs *Filesystem) FindContent(_ context.Context, key string) ([]byte, error) {
	return os.ReadFile(fs.generatePath(key))
}

func (fs *Filesystem) generatePath(key string) string {
	return filepath.Join(fs.dir, fmt.Sprintf("runestone-%s.txt", key))
}
