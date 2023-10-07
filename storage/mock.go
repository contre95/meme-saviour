package storage

import (
	"log/slog"
	"meme-saviour/app"
)

type MockStorage struct {
	name string
}

func NewMockStorage() *MockStorage {
	return &MockStorage{
		name: "Mock",
	}
}

func (s MockStorage) Save(m app.Meme) error {
	slog.Warn("Using Mock storage.", "warn", "nothign will be saved.")
	slog.Info("Saving file", "storage", s.GetName(), "name", m.Name, "size", m.Size, "link", m.Link)
	return nil
}

func (s MockStorage) GetName() string {
	return s.name
}

func (s MockStorage) GetRandomMeme() (*app.Meme, error) {
	panic("GetRandomMeme method is not implemented")
}
