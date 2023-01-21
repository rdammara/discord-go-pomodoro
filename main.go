package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"time"
)

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + os.Getenv("MTA2NjI1MDMzODk2Njk3MDM3MQ.G9Ax7c.ZHU-8w5k3Vvp3s02TDJ5k_R_2ko3qslP-86JWM"))
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// Simple way to keep program running until CTRL-C is pressed.

	//Create a new channel
	stopChan := make(chan bool)
	//Create a goroutine that will run the timer

	<-make(chan struct{})
	return
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, stopChan chan bool, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "pomodoro" start the timer
	if m.Content == "!pomodoro" {
		s.ChannelMessageSend(m.ChannelID, "Starting 25-minutes Pomodoro timer...")
		stopChan := make(chan bool)
		go timer(s, stopChan, m)
		//time.Sleep(25 * time.Minute)

	}
	//If the message is "stop", stop the timer
	if m.Content == "!stop" {
		s.ChannelMessageSend(m.ChannelID, "Stopping timer...")
		stopChan <- true
	}

}

// timer function
func timer(s *discordgo.Session, stopChan chan bool, m *discordgo.MessageCreate) {
	for {
		select {
		case <-stopChan:
			s.ChannelMessageSend(m.ChannelID, "Timer stopped.")
			return
		case <-time.After(25 * time.Minute):
			s.ChannelMessageSend(m.ChannelID, "Time is up! Take a break.")
		}
	}
}
