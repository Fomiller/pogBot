package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2/clientcredentials"

	twitch "github.com/gempir/go-twitch-irc/v3"
	"github.com/joho/godotenv"
)

var (
	oauth2Config *clientcredentials.Config
	Channel      string = "swolenesss"
	PogCount     *int
)

// type PogCountModifier interface {
// 	IncrementPogCount()
// }
// type PogCount struct {
// 	PogCountTotal     int
// 	PogCountPerMinute int
// }
//
// func (p *PogCount) IncrementPogCount() {
// 	p.PogCountTotal++
// }
type PogMessage struct {
	User    string
	Message string
	Time    time.Time
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file. ERR: %s", err)
	}
	ClientPassword := os.Getenv("CLIENT_PASSWORD")
	BotUsername := os.Getenv("BOTUSERNAME")
	client := twitch.NewClient(BotUsername, ClientPassword)
	var p int
	var pogList = []PogMessage{}
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if strings.Contains(strings.ToLower(message.Message), "pog") {
			x := PogMessage{message.User.Name, message.Message, message.Time}
			pogList = append(pogList, x)
			fmt.Println(len(pogList))
			fmt.Println(pogList)
			p++
			fmt.Println(p)

			client.Say(Channel, fmt.Sprintf("Pog has been said %v times", p))
		}
	})
	client.Join("swolenesss")
	err = client.Connect()
	if err != nil {
		panic(err)
	}
}
