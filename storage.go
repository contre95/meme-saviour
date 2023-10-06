package main

type Meme struct {
	Size float32
	Path string
}

type Saviour interface {
	Save(m Meme) error
	Name() string
	GetRandomMeme() (*Meme, error)
}
