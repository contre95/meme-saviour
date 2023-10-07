package storage

import (
	"log"
	"log/slog"
	"meme-saviour/app"
	"os"
)

type Local struct {
	maxSize float32
	name    string
	path    string
}

func NewLocalStorage(path string, maxSize float32) *Local {
	return &Local{
		name:    "Local",
		path:    path,
		maxSize: maxSize,
	}
}

func (s Local) Save(m app.Meme) error {
	slog.Info("Saving file", "storage", s.GetName(), "name", m.Name, "size", m.Size, "link", m.Link)
	file, err := os.Create(s.path + "/" + m.Name)
	if err != nil {
		log.Println("hi1")
		return err
	}
	defer file.Close()
	_, err = file.Write(m.File)
	if err != nil {
		return err
	}
	return nil
}

func (s Local) GetName() string {
	return s.name
}

func (s Local) GetRandomMeme() (*app.Meme, error) {
	panic("GetRandomMeme method is not implemented")
}
