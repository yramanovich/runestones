package manager

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/yramanovich/runestones/storage"
)

type Storage interface {
	SaveRunestone(ctx context.Context, key string, content []byte) error
	FindRunestone(ctx context.Context, key string) ([]byte, error)
}

func NewManager() *Manager {
	return &Manager{storage: &storage.Filesystem{}}
}

type Manager struct {
	storage Storage
}

func (m *Manager) CreateRunestone(ctx context.Context, content []byte) (string, error) {
	key := m.generateKey()
	if err := m.storage.SaveRunestone(ctx, key, content); err != nil {
		return "", err
	}
	ctx.Value("logger").(*slog.Logger).InfoContext(ctx, "new runestone created", "len", len(content), "id", key)
	return key, nil
}

func (m *Manager) GetRunestone(ctx context.Context, key string) ([]byte, error) {
	return m.storage.FindRunestone(ctx, key)
}

func (m *Manager) generateKey() string {
	return uuid.NewString()
}
