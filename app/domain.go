package app

type Size float32

type Meme struct {
	Link string
	Size Size
	Name string
}

type Storage interface {
	Save(m Meme) error
	GetName() string
	MaxSize() Size
	GetRandomMeme() (*Meme, error)
}
