package main

import (
	"log/slog"
	"meme-saviour/app"
	"meme-saviour/storage"
	"meme-saviour/telegram"
	"os"
	"strings"
	"time"

	"github.com/lmittmann/tint"
)

func main() {
	// Standarize logging configuration
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	// Use cases
	memeSaviour := app.NewMemeSaviour()

	// Local Storage
	if os.Getenv("LOCAL_STORAGE") == "1" {
		path := os.Getenv("LOCAL_STORAGE_PATH")
		if path == "" {
			path = "/data"
		}
		localStorage := storage.NewLocalStorage(path, 50.0)
		memeSaviour.RegisterStorage(localStorage)
	}

	// Mock storage
	// mockStorage := storage.NewMockStorage()
	// memeSaviour.RegisterStorage(mockStorage)

	// Bot configuration
	usernames := os.Getenv("TELEGRAM_ALLOWED_USERNAMES")
	if usernames == "" {
		slog.Error("TELEGRAM_ALLOWED_USERNAMES not set.")
		os.Exit(1)
	}
	botConfig := telegram.BotConfig{
		Token:          os.Getenv("TELEGRAM_TOKEN"),
		ValidUsernames: strings.Split(usernames, ","),
	}
	if memeSaviour.HasRegisteredStorages() {
		telegram.Run(botConfig, *memeSaviour)
	} else {
		slog.Error("Please set up at least one Storage for Meme Saviour to work.")
		os.Exit(1)

	}
}
