package telegram

import (
	"log"
	"log/slog"
	"meme-saviour/app"
	"os"
	"path/filepath"
	"slices"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotConfig struct {
	Token          string
	ValidUsernames []string
}

func (c *BotConfig) validate() {
	if len(c.Token) == 0 {
		slog.Error("Telegram token not set.")
		log.Fatal("Telegram token not set.")
	}
	slog.Info("Valid usernames set.", "usernames", c.ValidUsernames, "count", len(c.ValidUsernames))
	if len(c.ValidUsernames) == 0 {
		slog.Error("No valid Telegram username take")
		log.Fatal("No valid Telegram username take.")
	}
}

func Run(c BotConfig, ms app.MemeSaviour) {
	c.validate()
	bot, err := tgbotapi.NewBotAPI(c.Token)
	if err != nil {
		slog.Error("Couldnot request new bot API", "error", err)
		os.Exit(1)
	}
	slog.Info("Bot inititated", "name", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if !slices.Contains(c.ValidUsernames, update.Message.From.UserName) {
			slog.Warn("Wrong user identified")
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, your username is not allowed."))
			continue
		}
		if update.Message != nil && update.Message.Photo != nil {
			slog.Info("Meme identified")
			handleMeme(bot, c, update, ms)
		} else if update.Message != nil {
			slog.Info("Message is not a photo.")
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Please send a message to the chat in order to save a memee."))
		}
	}
}

func handleMeme(bot *tgbotapi.BotAPI, c BotConfig, update tgbotapi.Update, ms app.MemeSaviour) error {
	photo := update.Message.Photo[len(update.Message.Photo)-1]
	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: photo.FileID})
	if err != nil {
		log.Fatal(err)
	}
	fileExt := strings.ToLower(filepath.Ext(file.FilePath))
	// Create a new Meme
	mName := strings.ReplaceAll(update.Message.Caption, " ", "_")
	meme, err := app.NewMeme(mName, fileExt, file.Link(c.Token), []byte{})
	if err != nil {
		return err
	}
	// Save the Meme struct using the Saviour interface
	err = ms.SaveMemeTo("Local", *meme)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
