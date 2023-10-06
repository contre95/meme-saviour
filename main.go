package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Replace with your Telegram Bot API token
	botToken := os.Getenv("MEMESAVE_TELEGRAM_TOKEN")

	// Initialize the bot with your API token
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	// Set up updates channel to receive messages and files
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	// Handle incoming updates
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Photo != nil {
			// Handle incoming photos
			photo := (*update.Message.Photo)[len(*update.Message.Photo)-1]
			fileID := photo.FileID

			// Get the photo file
			fileConfig := tgbotapi.NewGetFile(fileID)
			file, err := bot.GetFile(fileConfig)
			if err != nil {
				log.Println(err)
				continue
			}

			// Download the photo
			fileURL := file.Link(botToken)
			response, err := http.Get(fileURL)
			if err != nil {
				log.Println(err)
				continue
			}
			defer response.Body.Close()

			// Create a folder to save the photos
			saveFolder := "photos"
			os.MkdirAll(saveFolder, os.ModePerm)

			// Create a file in the folder and save the photo
			fileName := filepath.Join(saveFolder, fmt.Sprintf("%s.jpg", fileID))
			outputFile, err := os.Create(fileName)
			if err != nil {
				log.Println(err)
				continue
			}
			defer outputFile.Close()

			// Copy the downloaded photo data to the output file
			_, err = io.Copy(outputFile, response.Body)
			if err != nil {
				log.Println(err)
				continue
			}

			log.Printf("Saved photo: %s\n", fileName)
		}
	}
}
