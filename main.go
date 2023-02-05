package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"os"
	"time"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json : "Token"`
	BotPrefix string `json : "BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, err := os.Open("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(fileBytes))

	err = json.Unmarshal(fileBytes, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil

}

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("messageHandler function called")
	if m.Author.ID == BotId {
		return
	}

	if m.Content == BotPrefix+"start" {
		s.ChannelMessageSend(m.ChannelID, "Starting a 25 minute Pomodoro Timer!")
		ticker := time.NewTicker(25 * time.Minute)
		go func() {
			for range ticker.C {
				s.ChannelMessageSend(m.ChannelID, "Time's up! Take a break.")
			}
		}()
	}
}

func main() {
	err := ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Start()

	<-make(chan struct{})
	return
}
