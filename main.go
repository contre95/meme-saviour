package main

import (
	"log/slog"
	"meme-saviour/app"
	"meme-saviour/storage"
	"meme-saviour/telegram"
	"os"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	})))
	botConfig := telegram.BotConfig{
		Token: os.Getenv("MEMESAVE_TELEGRAM_TOKEN"),
	}
	mockSaviour := storage.NewMockStorage()
	memeSaviour := app.NewMemeSaviour()
	memeSaviour.RegisterStorage(mockSaviour)
	telegram.Run(botConfig, *memeSaviour)
}
