package storage

import "meme-saviour/app"

type MockSaviour struct {
	name string
}

func NewMockSaviour() *MockSaviour {
	return &MockSaviour{
		name: "Mock",
	}
}

func (s MockSaviour) Save(m app.Meme) error {
	panic("Save method is not implemented")
}

func (s MockSaviour) GetName() string {
	return s.name
}

func (s MockSaviour) MaxSize() app.Size {
	panic("MaxSize method is not implemented")
}

func (s MockSaviour) GetRandomMeme() (*app.Meme, error) {
	panic("GetRandomMeme method is not implemented")
}
