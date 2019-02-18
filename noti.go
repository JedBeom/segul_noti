package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/mmcdole/gofeed"
)

var old, _ = getItems()

func getItems() (items []*gofeed.Item, err error) {

	p := gofeed.NewParser()
	feed, err := p.ParseURL("http://wangun.ms.jne.kr/rssList.jsp?siteId=wangun_ms&boardId=775784")
	if err != nil {
		log.Println("Feed Parsing:", err)
		return
	}

	items = feed.Items
	return

}

func checkNew() (posts []*gofeed.Item) {
	items, err := getItems()
	if err != nil {
		return
	}

	for _, item := range items {
		if old[0].Title == item.Title {

			if len(posts) != 0 {
				log.Println("Post finding break")
			}
			break
		}

		posts = append(posts, item)

	}

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
		log.Println("No New Posts")
		return
	}

	var embed discordgo.MessageEmbed
	embed.Author = &author

	for _, post := range posts {
		addField(&embed.Fields, post.Title, "[바로가기]("+post.Link+")")
	}

	for _, request := range requests {

		fmt.Println(request)
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
