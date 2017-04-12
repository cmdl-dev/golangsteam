package main

import (
	"io/ioutil"
	"log"

	"fmt"
	steam "github.com/Philipp15b/go-steam"
	"github.com/Philipp15b/go-steam/protocol/steamlang"
	"github.com/Philipp15b/go-steam/socialcache"
	//"github.com/Philipp15b/go-steam/steamid"
)

const USERNAME = "golangbot"
const PASSWORD = "572676cah1"
const CODE = "HHWJY"

type Client struct {
	Username   string
	Password   string
	FriendList *socialcache.FriendsList
	client     *steam.Client
}

func main() {

	c := newClient(USERNAME, PASSWORD)
	myLoginInfo := new(steam.LogOnDetails)
	myLoginInfo.Username = c.Username
	myLoginInfo.Password = c.Password
	myLoginInfo.AuthCode = CODE
	myLoginInfo.TwoFactorCode = CODE
	myLoginInfo.SentryFileHash, _ = ioutil.ReadFile("sentry")
	//friendList := socialcache.NewFriendsList()

	err := steam.InitializeSteamDirectory()

	if err != nil {
		log.Panic(err.Error())
	}
	c.client.Connect()

	for event := range c.client.Events() {
		switch e := event.(type) {
		case *steam.ConnectedEvent:
			log.Println("Connecting SteamBot")
			c.client.Auth.LogOn(myLoginInfo)
		case *steam.MachineAuthUpdateEvent:
			log.Println("Wrote Sentry")
			ioutil.WriteFile("sentry", e.Hash, 0666)
		case *steam.LoggedOnEvent:
			log.Println("SteamBot logged in")
			c.client.Social.SetPersonaState(steamlang.EPersonaState_Online)
		case *steam.LogOnFailedEvent:
			log.Println("Failed because: ", e)
		case *steam.ChatMsgEvent:
			go c.chatMsgEvent(e)
		case *steam.ChatInviteEvent:
			fmt.Println("invite chat")
		case *steam.ChatEnterEvent:
			fmt.Println("enter chat")
		case *steam.ChatActionResultEvent:
			fmt.Println("chat action")
		case *steam.ChatMemberInfoEvent:
			fmt.Println("Chatmember")
		case *steam.FriendStateEvent:
			fmt.Println("Friend")
			go c.friendStateEvent(e)
		case *steam.FriendAddedEvent:
			fmt.Println("added friend")
			go friendAddedEvent(e, c.client)
		case *steam.DisconnectedEvent:
			log.Print("Disconnected")
			log.Print("attempting to recconnoct")

		case error:
			log.Println(e)
			//default:
			//log.Println(e)
		}
	}
}
