package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/keftcha/markovchaingo"
)

var discordToken string
var mcg *markovchaingo.MarkovChainGo

func init() {
	discordToken = os.Getenv("DISCORD")

	if discordToken == "" {
		panic("No discord token found in envoronment variable `DISCORD_TOKEN`.")
	}

	mcg = markovchaingo.New("file:///data.json")
}

func main() {
	// Create the bot instance
	bot, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.Fatal(err)
	}

	// Register callback function
	bot.AddHandler(learn)
	bot.AddHandler(talk)

	// Open a websocket connection to Discord and begin listening.
	err = bot.Open()
	if err != nil {
		log.Fatal(err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("I'm logged in ! (Press CTRL-C to exit.)")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()
}

func learn(s *discordgo.Session, m *discordgo.MessageCreate) {
	mcg.Learn(m.Content)
}

func talk(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Check if we are mentioned
	shouldAnswer := false
	for _, mentionedUsers := range m.Mentions {
		if mentionedUsers.ID == s.State.User.ID {
			shouldAnswer = true
			break
		}
	}
	if shouldAnswer && m.Author.ID != s.State.User.ID {
		sentence := mcg.Talk()
		s.ChannelMessageSend(m.ChannelID, sentence)
	}
}
