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

// RegisterStorage register a new Storage with it's name as a key.
func (m *MemeSaviour) RegisterStorage(s Storage) {
	slog.Info("New storage registered", "stoarge", s.GetName())
	m.storage[s.GetName()] = s
}

func (m *MemeSaviour) SaveMemeTo(storagekey string, meme Meme) error {
	s, keyExists := m.storage[storagekey]
	if keyExists {
		err := s.Save(meme)
		if err != nil {
			slog.Error("Meme could not be saved", "name", meme.Name, "storage", s.GetName(), "error", err)
			return nil
		}
	}
	slog.Info("Meme save.", "name", meme.Name)
	return nil
}

func validateSize(s Storage) (bool, error) {
	slog.Info("Validating size.", "storage", s.GetName())
	return true, nil
}

func GetRandomMeme(s Storage) (*Meme, error) {
	return nil, nil
}
