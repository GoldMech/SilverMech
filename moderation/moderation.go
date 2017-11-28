package moderation

import "github.com/bwmarrin/discordgo"

/*
	Runs on silverMech ready event
*/
func Ready(s *discordgo.Session, event *discordgo.Ready) {
	
	s.UpdateStatus(0, "^help for more info!")

}