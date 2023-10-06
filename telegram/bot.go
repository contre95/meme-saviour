package telegram

import (
	"log"
	"log/slog"
	"meme-saviour/app"
	"slices"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotConfig struct {
	Token          string
	ValidUsernames []string
}

func Run(c BotConfig, ms app.MemeSaviour) {
	bot, err := tgbotapi.NewBotAPI(c.Token)
	if err != nil {
		log.Fatal(err)
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

func handleMeme(bot *tgbotapi.BotAPI, c BotConfig, update tgbotapi.Update, ms app.MemeSaviour) {
	photo := update.Message.Photo[len(update.Message.Photo)-1]
	// Download the photo
	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: photo.FileID})
	if err != nil {
		log.Fatal(err)
	}
	// Create a new Meme struct

	meme := app.Meme{
		Size: app.Size(photo.Width),
		Link: file.Link(c.Token),
		Name: strings.ReplaceAll(update.Message.Caption, " ", "_"),
	}
	// Save the Meme struct using the Saviour interface
	err = ms.SaveMemeTo("mock", meme)
	if err != nil {
		log.Fatal(err)
	}

}
