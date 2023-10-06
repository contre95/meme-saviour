package app

type Size float32

type Meme struct {
	// Image []byte // Not sure if we are gonna need this.
	Size Size
	Path string
	Name string
}

type Storage interface {
	Save(m Meme) error
	GetName() string
	MaxSize() Size
	GetRandomMeme() (*Meme, error)
}
