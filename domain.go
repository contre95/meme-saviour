package main

type Size float32

type Meme struct {
	// Image []byte // Not sure if we are gonna need this.
	Size Size
	Path string
	Name string
}

type Saviour interface {
	Save(m Meme) error
	Name() string
	MaxSize() Size
	GetRandomMeme() (*Meme, error)
}
