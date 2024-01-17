/*
Copyright © 2024 Pavlo Tarasiuk <pasha.tarasyuk@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

// startCmd represents the mavbot command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Run the main functionality of MAVBot",
	Long: `The mavbot command executes the main functionality of MAVBot,
	including interaction with Slack and other bot features.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("MAVBot %s started", appVersion)

		// Load Env variables from .env file
		godotenv.Load(".env")

		token := os.Getenv("SLACK_AUTH_TOKEN")
		channelId := os.Getenv("SLACK_CHANEL_ID")

		// Create a new client to slack by giving token
		// Set debug to true while developing
		client := slack.New(token, slack.OptionDebug(true))
		attachment := slack.Attachment{
			Pretext: "Super Bot Message",
			Text:    "some text",
			Color:   "4af030",
			Fields: []slack.AttachmentField{
				{
					Title: "Date",
					Value: time.Now().String(),
				},
			},
		}

		_, timestamp, err := client.PostMessage(
			channelId,

			slack.MsgOptionAttachments(attachment),
		)

		if err != nil {
			panic(err)
		}
		fmt.Printf("Message sent at %s", timestamp)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
