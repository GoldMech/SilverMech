package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	"./currency"
	"./info"
	"./moderation"
)

/**

	Based Off: https://github.com/bwmarrin/discordgo/blob/master/examples/airhorn/main.go

**/
func main() {

	fmt.Println()

	/*
		Enviromental File
	*/
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file failed to load!")
	}

	token := os.Getenv("DISCORD_TOKEN")

	/*
		Create Discord Bot
	*/
	silverMech, err := discordgo.New("Bot " + token)
	if err != nil {
		return
	}
	fmt.Println("Bot Connection Created Sucessfully.")

	/*
		Startup Data
	*/
	info.RefreshFileList()

	/*
		Handlers
	*/
	silverMech.AddHandler(moderation.Ready)
	silverMech.AddHandler(messageCreate)

	/*
		Open Discord Bot Connection
	*/
	err = silverMech.Open()
	if err != nil {
		fmt.Println("Bot Failed to Open")
	} else {
		fmt.Println("Bot Connection Opened Sucessfully.")
	}

	/*
		Memory clean and closing control
	*/
	fmt.Println()
	fmt.Println("SilverMech Running Now.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	silverMech.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content != "" {
		if (m.Content[0] == '^' || m.Content[0] == '~') && m.Type == 0 {

			/*
				Retrieve Channel Information
			*/
			var useChannel string

			postChannel, err := s.UserChannelCreate(m.Author.ID)
			if err != nil {
				fmt.Println("Couldn't get the channel struct")
				return
			}

			privateChannel, err := s.Channel(m.ChannelID)
			if err != nil {
				fmt.Println(err)
				return
			}

			if m.Content[0] == '~' {
				useChannel = privateChannel.ID
			} else {
				useChannel = postChannel.ID
			}

			/*
				Check if the ^ command is being used in a channel on the guild ID specified in .env
			*/
			if postChannel.GuildID != os.Getenv("GUILD_ID") && postChannel.GuildID != "" {
				s.ChannelMessageSend(postChannel.ID, postChannel.GuildID+" is not a valid guild id! Make sure you are in our discord, and not using ^ commands in private chat!")

			} else {

				fmt.Println(m.Author.Username + ": " + m.Content)
				switch c := strings.Split(m.Content[1:], " "); strings.ToLower(c[0]) {

				case "hello":
					s.ChannelMessageSend(postChannel.ID, "Hello "+m.Author.Mention()+"!")

				case "info":
					info.SafeInfoFIle(s, useChannel, c[1])

				case "help":
					info.InfoFile(s, useChannel, "help")

				case "rules":
					info.InfoFile(s, useChannel, "rules")

				case "welcome":
					info.InfoFile(s, useChannel, "welcome")

				case "balance":
					currency.Balance(s, useChannel, "1000", m.Author.Username)

				case "test":
					info.InfoFile(s, useChannel, "rules")

				default:
					s.ChannelMessageSend(useChannel, "Unknown Command: "+c[0])

				}

			}

		}
	}

}
