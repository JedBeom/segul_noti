package main

import (
	disgo "github.com/bwmarrin/discordgo"
	"log"
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
		reply(s, m, "알겠어요.")
	}
}

/*
func printPosts(s *disgo.Session, m *disgo.MessageCreate) {
	reply(s, m, "잠시만 기다려주세요!")
	posts, err  := getPosts()
	if err != nil {
		reply(s, m, "죄송해요. 쥬니올이 게시물들을 물어오지 못했어요.")
		return
	}

	answer := "쥬니올이 게시물들을 가져왔어요!\n"
	for _, post := range posts[1:6] {
		answer += fmt.Sprintf("%s %s\n", post.Title, post.Link)
	}
	reply(s, m, answer)
}
*/

// 부른 사람에게 멘션하기
func reply(s *disgo.Session, m *disgo.MessageCreate, content string) {
	message := "<@" + m.Author.ID + "> " + content
	msg, err := s.ChannelMessageSend(m.ChannelID, message)
	if err != nil {
		log.Println("Error Replying\nMsg:", msg, "\nErr:", err)
	}
}
