package app

import (
	"log/slog"
)

// This is the interactor for our use cases.
type MemeSaviour struct {
	storage map[string]Storage
}

func NewMemeSaviour() *MemeSaviour {
	return &MemeSaviour{
		storage: map[string]Storage{},
	}
}

func (m *MemeSaviour) RegisterStorage(s Storage) {
	slog.Info("New storage registered", "stoarge", s.GetName())
	m.storage[s.GetName()] = s
}

func (m *MemeSaviour) SaveMemeTo(storagekey string, meme Meme) error {
	s, keyExists := m.storage[storagekey]
	if keyExists {
		err := s.Save(meme)
		if err != nil {
			slog.Error("Meme could not be saved", "meme", meme.Name, "storage", s.GetName())
			return nil
		}
	}
	slog.Info("Meme save.", "name", meme.Name)
	return nil
}

func validateSize(s Storage) (bool, error) {
	slog.Info("Validating size.", "size", s.MaxSize(), "storage", s.GetName())
	return true, nil
}

func GetRandomMeme(s Storage) (*Meme, error) {
	return nil, nil
}
