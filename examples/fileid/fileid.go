package main

import (
	"github.com/xsavix/gogram/telegram"
)

const (
	appID    = 6
	appHash  = "YOUR_APP_HASH"
	phoneNum = "YOUR_PHONE_NUMBER"
)

func main() {
	// Create a new client
	client, _ := telegram.NewClient(telegram.ClientConfig{
		AppID:    appID,
		AppHash:  appHash,
		LogLevel: telegram.LogInfo,
	})

	client.Conn() // Connect to the telegram server

	// Authentication
	client.LoginBot("<BOT_TOKEN_HERE>")

	client.On(telegram.OnMessage, func(msg *telegram.NewMessage) error {
		if msg.Media() != nil {
			fileID := msg.File.FileID // or telegram.PackBotFileID(msg.Media())

			msg.Respond("File ID: " + fileID)

			fileRetained, err := telegram.ResolveBotFileID(fileID)
			if err != nil {
				return err
			}

			msg.RespondMedia(fileRetained)
		}

		return nil
	})
}
