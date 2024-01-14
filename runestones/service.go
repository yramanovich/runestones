package runestones

import (
	"context"

	"github.com/google/uuid"

	"github.com/yramanovich/runestones/log"
)

type ContentsManager interface {
	SaveContent(ctx context.Context, url string, content []byte) error
	FindContent(ctx context.Context, url string) ([]byte, error)
}

type Repository interface {
	CreateRunestone(ctx context.Context, url string) (string, error)
	FindRunestone(ctx context.Context, id string) (Runestone, error)
}

func NewManager(cm ContentsManager, repo Repository) *Manager {
	return &Manager{contentsManager: cm, repository: repo}
}

type Manager struct {
	contentsManager ContentsManager
	repository      Repository
}

func (m *Manager) CreateRunestone(ctx context.Context, content []byte) (string, error) {
	url := m.generateUrl()
	if err := m.contentsManager.SaveContent(ctx, url, content); err != nil {
		return "", err
	}

	id, err := m.repository.CreateRunestone(ctx, url)
	if err != nil {
		return "", err
	}
	log.FromContext(ctx).DebugContext(ctx, "new runestone created", "len", len(content), "id", id)

	return id, nil
}

func (m *Manager) GetRunestone(ctx context.Context, id string) ([]byte, error) {
	runestone, err := m.repository.FindRunestone(ctx, id)
	if err != nil {
		return nil, err
	}

	contents, err := m.contentsManager.FindContent(ctx, runestone.Url)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func (m *Manager) generateUrl() string {
	return uuid.NewString()
}
