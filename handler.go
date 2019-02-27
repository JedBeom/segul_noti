package main

import (
	"fmt"
	"log"

	disgo "github.com/bwmarrin/discordgo"
)

var requests = make(map[string][]string)

func handler(s *disgo.Session, m *disgo.MessageCreate) {
	if m.Content == "~sub" {

		if users, ok := requests[m.ChannelID]; ok {

			for _, user := range users {
				if user == m.Author.ID {
					reply(s, m, "이미 오야붕에게 가져가게 되어있어! 그만 가져오게 하고 싶으면 `~unsub`라고 입력하면 돼!")
					return
				}
			}

			requests[m.ChannelID] = append(users, m.Author.ID)
		} else {
			requests[m.ChannelID] = []string{m.Author.ID}
		}

		reply(s, m, "알겠어, 오야붕! 새 게시물이 있으면 줄게! 그만 받고 싶으면 `~unsub`라고 말해줘!")
		fmt.Println("Subs:", requests)

	} else if m.Content == "~posts" {
		printPosts(s, m)
	} else if m.Content == "~unsub" {

		if users, ok := requests[m.ChannelID]; ok {

			for i, user := range users {
				if user == m.Author.ID {

					requests[m.ChannelID] = append(users[:i], users[i+1:]...)

					if len(requests[m.ChannelID]) == 0 {
						delete(requests, m.ChannelID)
					}

					reply(s, m, "오케이, 이제 그만 가져올게!")
					return
				}
			}

		}

		reply(s, m, "원래 오야붕에게는 안 주는 걸... 받고 싶으면 `~sub`라고 말해봐.")
	}

}

func printPosts(s *disgo.Session, m *disgo.MessageCreate) {
	reply(s, m, "잠시만 기다려줘!")
	posts, err := getItems()
	if err != nil {
		reply(s, m, "미안, 못 가져왔어!")
		return
	}

	var embed disgo.MessageEmbed
	embed.Author = &author

	for _, post := range posts {
		addField(&embed.Fields, post.Title, "[바로가기]("+post.Link+")")
	}

	send := disgo.MessageSend{
		Content: "<@" + m.Author.ID + "> 여기 지금 게시물!",
		Embed:   &embed,
	}

	msg, err := dg.ChannelMessageSendComplex(m.ChannelID, &send)
	if err != nil {
		log.Println("Error Replying\nMsg:", msg, "\nErr:", err)
	}
}

// 부른 사람에게 멘션하기
func reply(s *disgo.Session, m *disgo.MessageCreate, content string) {
	message := "<@" + m.Author.ID + "> " + content
	msg, err := s.ChannelMessageSend(m.ChannelID, message)
	if err != nil {
		log.Println("Error Replying\nMsg:", msg, "\nErr:", err)
	}
}
