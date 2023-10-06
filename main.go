package main

import (
	"meme-saviour/app"
	"meme-saviour/storage"
	"meme-saviour/telegram"
	"os"
)

func main() {
	botConfig := telegram.BotConfig{
		Token: os.Getenv("MEMESAVE_TELEGRAM_TOKEN"),
	}
	mockSaviour := storage.NewMockSaviour()
	memeSaviour := app.NewMemeSaviour()
	memeSaviour.RegisterStorage(mockSaviour)
	telegram.Run(botConfig, *memeSaviour)
}
