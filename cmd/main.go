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

const (
	botUsername               = "poggersbot69"
	clientUsername            = "swolenesss"
	justinFan                 = "justinfan123123"
	clientAuthenticationToken = "oauth:123123123"
	myCreds                   = "oauth:g9yzwusiujjvdoxb0am9n0i5so1nqw"
)

var (
	oauth2Config *clientcredentials.Config
	Channel      string = "swolenesss"
	BotUserName  string = "swolenesss"
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
	client := twitch.NewClient(botUsername, ClientPassword)
	// var pogCount PogCount
	var p int
	var pogList = []PogMessage{}
	// client.OnPrivateMessage(func(message twitch.PrivateMessage) {
	// 	if strings.Contains(strings.ToLower(message.Message), "ping") {
	// 		log.Println(message.User.Name, "PONG", message.Message)
	// 		// pogCount.IncrementPogCount()
	// 		// p++
	// 		// fmt.Println(p)
	// 		// log.Println(pogCount.PogCountTotal)
	//
	// 		client.Say(Channel, "poggers")
	// 	}
	// })
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if strings.Contains(strings.ToLower(message.Message), "pog") {
			// log.Println(message.User.Name, "POG", message.Message)
			// log.Println(message.Message)
			// log.Println(message.Time)
			x := PogMessage{message.User.Name, message.Message, message.Time}
			pogList = append(pogList, x)
			fmt.Println(len(pogList))
			fmt.Println(pogList)
			// pogCount.IncrementPogCount()
			p++
			fmt.Println(p)
			// log.Println(pogCount.PogCountTotal)

			client.Say(Channel, fmt.Sprintf("Pog has been said %v times", p))
		}
	})
	// client.OnPrivateMessage(func(message twitch.PrivateMessage) {
	// 	log.Println(message)
	// 	log.Println(message.User.Name)
	// 	log.Println(message.Message)
	// })
	client.Join("swolenesss")
	err = client.Connect()
	if err != nil {
		panic(err)
	}
}
