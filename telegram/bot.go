package telegram

import (
	"log"
	"meme-saviour/app"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotConfig struct {
	Token string
}

func Run(c BotConfig, ms app.MemeSaviour) {
	bot, err := tgbotapi.NewBotAPI(c.Token)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// updates, err := bot.GetUpdates(context.Background(), u)
    uConfig := tgbotapi.UpdateConfig{
    	Offset:         0,
    	Limit:          0,
    	Timeout:        0,
    	AllowedUpdates: []string{},
    }
	updates, err := bot.GetUpdates()
	if err != nil {
		log.Fatal(err)
	}

	for _, update := range updates {
		if update.Message != nil && update.Message.Photo != nil {
			// Get the largest photo size
			photo := update.Message.Photo[len(update.Message.Photo)-1]

			// Download the photo
			// file, err := bot.GetFile(context.Background(), photo.FileID)
			fConfig := tgbotapi.FileConfig{
				FileID: "",
			}
			file, err := bot.GetFile(fConfig)
			if err != nil {
				log.Fatal(err)
			}

			// Create a new Meme struct
			meme := app.Meme{
				Size: app.Size(photo.Width),
				Path: file.FilePath,
				Name: photo.FileID + ".jpg",
			}

			// Save the Meme struct using the Saviour interface
			err = ms.SaveMemeTo("mock", meme)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Meme saved to %s", meme.Name)
		}
	}
}
