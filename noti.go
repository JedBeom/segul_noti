package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	old, _ = getPosts()
)

func checkNew() (newPosts []Post) {
	posts, err := getPosts()
	if err != nil {
		log.Println("Parsing:", err)
		return
	}

	if len(old) != len(posts) {
		log.Println("length of old and posts' is different.")
		old = posts
		return
	}

	for i, post := range posts {
		if old[i].Title == post.Title {
			break
		}

		newPosts = append(newPosts, post)
	}

	old = posts

	return

}

var (
	author = discordgo.MessageEmbedAuthor{
		URL:     "https://github.com/JedBeom/segul_noti",
		Name:    "하코자키 세리카",
		IconURL: "https://raw.githubusercontent.com/JedBeom/choicebot_discord/master/serika.png",
	}
)

func notification() {
	posts := checkNew()
	if len(posts) == 0 {
		return
	}

	var embed discordgo.MessageEmbed
	embed.Author = &author

	for _, post := range posts {
		addField(&embed.Fields, post.Title, "[바로가기]("+post.Link+")")
	}

	for _, request := range requests {

		send := discordgo.MessageSend{
			Content: "<@" + request.AuthorID + "> 쥬니올이 새 게시물을 가져왔어요!",
			Embed:   &embed,
		}
		msg, err := dg.ChannelMessageSendComplex(request.ChannelID, &send)
		if err != nil {
			log.Println("Error Replying\nMsg:", msg, "\nErr:", err)
		}

	}
}
