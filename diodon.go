package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/keftcha/markovchaingo"
)

var discordToken string
var talkiness float64
var mcg *markovchaingo.MarkovChainGo

func init() {
	discordToken = os.Getenv("DISCORD")
	if discordToken == "" {
		panic("No discord token found in envoronment variable `DISCORD_TOKEN`.")
	}

	var err error
	talkiness, err = strconv.ParseFloat(os.Getenv("TALKINESS"), 64)
	if err != nil {
		panic(err)
	}

	connectionString := os.Getenv("CONNECTION_STRING")
	if connectionString == "" {
		panic("No connection string provided")
	}
	mcg = markovchaingo.New(connectionString)
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
	err := mcg.Learn(m.Content)
	if err != nil {
		fmt.Println(err)
	}
}

func talk(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Check if we are mentioned
	shouldAnswer := mentioned(m.Mentions, s.State.User) || rand.Float64() < talkiness

	if shouldAnswer && m.Author.ID != s.State.User.ID {
		if sentence, err := mcg.Talk(); err != nil {
			fmt.Println(err)
			s.ChannelMessageSend(m.ChannelID, "An error occurred.")
		} else {
			s.ChannelMessageSend(m.ChannelID, sentence)
		}
	}
}

func mentioned(users []*discordgo.User, user *discordgo.User) bool {
	for _, mentionedUsers := range users {
		if mentionedUsers.ID == user.ID {
			return true
		}
	}
	return false
}
