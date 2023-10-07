package app

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type Size float32

type Meme struct {
	Link string
	File []byte
	Size Size
	Name string
}

func NewMeme(name, extension, link string, file []byte) (*Meme, error) {
	if link == "" && len(file) == 0 {
		return nil, errors.New("Your need to provide either the link or the []bytes.")
	}
	if len(extension) == 0 || extension[0] != '.' {
		return nil, errors.New("Your need to provide and extension in the form of '.xxx' (eg. .jpg, .png)")
	}
	if len(name) != 0 {
		name = "_" + name
	}
	meme := Meme{
		Link: link,
		File: file,
		Size: 0,
		Name: time.Now().Format("20060102150405") + name + extension,
	}
	if len(file) == 0 {
		if err := meme.downloadFromLink(); err != nil {
			return nil, err
		}

	}
	return &meme, nil
}

// DownloadFromLink fills the Meme.File field if not present with the link present at m.Link
func (m *Meme) downloadFromLink() error {
	if m.File == nil {
		slog.Warn("Skiping file download. Already downloaded.", "len", len(m.File), "name", m.Name)
		return nil
	}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, m.Link, nil)
	if err != nil {
		return errors.New("Couldn't get file from link.")
	}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("Couldn't request link.")
	}
	defer resp.Body.Close()
	file, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Invalid file downloaded from link.")
	}
	m.File = file
	return nil
}

type Storage interface {
	Save(m Meme) error
	GetName() string
	GetRandomMeme() (*Meme, error)
}
