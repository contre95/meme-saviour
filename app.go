package main

import (
	"errors"
	"log/slog"
)

func SaveMeme(s Saviour, m Meme) error {
	if m.Size > s.MaxSize() {
		slog.Error("Meme size is to big", "size", m.Size)
		return errors.New("Meme size is to big")
	}

	err := s.Save(m)
	if err != nil {
		slog.Error("Meme could not be saved", "meme", m.Name, "saviour", s.Name())
		return nil
	}
	slog.Info("Meme save.", "path", m.Path)
	return nil
}

func validateSize(s Saviour) (bool, error) {
	slog.Info("Validating size.", "size", s.MaxSize(), "saviour", s.Name())
	return true, nil
}
func GetRandomMeme(s Saviour) (*Meme, error) {
	return nil, nil
}
