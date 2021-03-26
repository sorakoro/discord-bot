package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	channelID string
	message   string
)

func init() {
	flag.StringVar(&channelID, "c", "", "channel ID")
	flag.StringVar(&message, "m", "", "message")
	flag.Parse()
}

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	_, err = sess.ChannelMessageSend(channelID, message)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
