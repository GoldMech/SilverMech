package info

import (
	"fmt"
	"io/ioutil"

	"github.com/bwmarrin/discordgo"
)

var fileList map[string]bool
var fileListString string

func RefreshFileList() {
	files, err := ioutil.ReadDir("./config/messages")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileList = make(map[string]bool)

	fileListString = "**List of info topics:** \n" + "```\n"

	for _, file := range files {
		fmt.Println("Loaded file: " + file.Name())
		fileListString += file.Name()[:len(file.Name())-4] + "\n"
		fileList[file.Name()[:len(file.Name())-4]] = true
	}

	fileListString = fileListString + "```"
}

func InfoFile(s *discordgo.Session, c string, file string) {

	b, err := ioutil.ReadFile("./config/messages/" + file + ".txt")

	if err != nil {
		s.ChannelMessageSend(c, "Error finding info, try again later.")
		return
	}

	fileOutput := string(b)

	s.ChannelMessageSend(c, fileOutput)

}

func List(s *discordgo.Session, c string) {
	RefreshFileList()

	s.ChannelMessageSend(c, fileListString)
}

func SafeInfoFIle(s *discordgo.Session, c string, file string) {

	fmt.Println("FILE: !" + file + "!")
	if file == "list" {
		List(s, c)
	} else if fileList[file] == true {
		InfoFile(s, c, file)
	} else {
		fmt.Println(fileList[file])
		s.ChannelMessageSend(c, "Invalid topic")
	}

}
