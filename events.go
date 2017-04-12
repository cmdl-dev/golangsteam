package main

import (
	//"io/ioutil"
	"log"

	"fmt"
	steam "github.com/Philipp15b/go-steam"
	//"github.com/Philipp15b/go-steam/steamid"
	"github.com/Philipp15b/go-steam/protocol/steamlang"
	"github.com/Philipp15b/go-steam/socialcache"
)

func newClient(username string, password string) *Client {
	return &Client{
		username,
		password,
		socialcache.NewFriendsList(),
		steam.NewClient(),
	}
}

func (c *Client) friendStateEvent(e *steam.FriendStateEvent) {
	switch e.Relationship {
	case steamlang.EFriendRelationship_None:
		fmt.Println("helo1")
	case steamlang.EFriendRelationship_RequestRecipient:
		c.client.Social.AddFriend(e.SteamId)
		fmt.Println("helo2")
	case steamlang.EFriendRelationship_Friend:
		fmt.Println("helo3")
	}
}
func friendAddedEvent(e *steam.FriendAddedEvent, client *steam.Client) {
	client.Social.SendMessage(e.SteamId, steamlang.EChatEntryType_ChatMsg, "Sup Its me your brother")
}
func (c *Client) chatMsgEvent(e *steam.ChatMsgEvent) {
	if e.IsMessage() {
		log.Println(e.Message)
		steamID := e.ChatterId
		fmt.Println(steamID)

	}
	c.client.Social.JoinChat(e.ChatterId)
	c.client.Social.SendMessage(e.ChatterId, steamlang.EChatEntryType_ChatMsg, c.client.Social.GetPersonaName())
	c.client.Social.RemoveFriend(e.ChatterId)
	c.client.Social.LeaveChat(e.ChatterId)

}
