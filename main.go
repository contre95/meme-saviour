package main

import (
	"log/slog"
	"meme-saviour/app"
	"meme-saviour/storage"
	"meme-saviour/telegram"
	"os"
	"strings"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		// Level: slog.LevelWarn,
	})))
	botConfig := telegram.BotConfig{
		Token:          os.Getenv("BOT_MEMESAVE_TELEGRAM_TOKEN"),
		ValidUsernames: strings.Split(os.Getenv("BOT_ALLOWED_USERNAMES"), ","),
	}
	mockSaviour := storage.NewMockStorage()
	memeSaviour := app.NewMemeSaviour()
	memeSaviour.RegisterStorage(mockSaviour)
	telegram.Run(botConfig, *memeSaviour)
}
