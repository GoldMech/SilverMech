package currency

import "github.com/bwmarrin/discordgo"

func Balance(s *discordgo.Session, c string, userID string, user string) {

	
		s.ChannelMessageSend(c, "**" + user + "'s balance:** \n" +
								":gem: 5,000 :gem:")

	
}