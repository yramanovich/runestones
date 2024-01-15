package runestones

import (
	"context"

	"github.com/google/uuid"

	"github.com/yramanovich/runestones/log"
)

// ContentsManager allows to save and find content.
type ContentsManager interface {
	SaveContent(ctx context.Context, url string, content []byte) error
	FindContent(ctx context.Context, url string) ([]byte, error)
}

// Repository manages runestone additional information.
type Repository interface {
	CreateRunestone(ctx context.Context, url string) (string, error)
	FindRunestone(ctx context.Context, id string) (Runestone, error)
}

// NewService returns new Service instance.
func NewService(cm ContentsManager, repo Repository) *Service {
	return &Service{contentsManager: cm, repository: repo}
}

// Service is the main component in teh system which manages interaction with runestone items.
type Service struct {
	contentsManager ContentsManager
	repository      Repository
}

// CreateRunestone saves runestones in the system and returns its id.
func (m *Service) CreateRunestone(ctx context.Context, content []byte) (string, error) {
	url := m.generateURL()
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

// GetRunestone returns runestone associated with the given id.
func (m *Service) GetRunestone(ctx context.Context, id string) ([]byte, error) {
	runestone, err := m.repository.FindRunestone(ctx, id)
	if err != nil {
		return nil, err
	}

	contents, err := m.contentsManager.FindContent(ctx, runestone.URL)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func (m *Service) generateURL() string {
	return uuid.NewString()
}
