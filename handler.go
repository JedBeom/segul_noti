package main

import (
	"fmt"
	"log"

	disgo "github.com/bwmarrin/discordgo"
)

type Request struct {
	ChannelID string
	AuthorID  string
}

var (
	requests []Request
)

func handler(s *disgo.Session, m *disgo.MessageCreate) {
	if m.Content == "~sub" {

		requests = append(requests, Request{m.ChannelID, m.Author.ID})
		reply(s, m, "알겠어요! 새 게시물이 있으면 쥬니올이 가져오게 할게요!")
		fmt.Println("Subs:", requests)

	} else if m.Content == "~posts" {
		printPosts(s, m)
	}
}

func printPosts(s *disgo.Session, m *disgo.MessageCreate) {
	reply(s, m, "잠시만 기다려주세요!")
	posts, err := getItems()
	if err != nil {
		reply(s, m, "죄송해요. 쥬니올이 게시물들을 물어오지 못했어요.")
		return
	}

	var embed disgo.MessageEmbed
	embed.Author = &author

	for _, post := range posts {
		addField(&embed.Fields, post.Title, "[바로가기]("+post.Link+")")
	}

	send := disgo.MessageSend{
		Content: "<@" + m.Author.ID + "> 쥬니올이 새 게시물을 가져왔어요!",
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
