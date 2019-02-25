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

	old = items

	return

}

var (
	author = discordgo.MessageEmbedAuthor{
		URL:     "https://github.com/JedBeom/segul_noti",
		Name:    "오오가미 타마키",
		IconURL: "https://raw.githubusercontent.com/JedBeom/segul_noti/master/tamaki.png",
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
	embed.Title = "순천왕운중학교 알림판 새 게시물"
	embed.Color = 0xed90ba

	for _, post := range posts {
		addField(&embed.Fields, post.Title, "[바로가기]("+post.Link+")")
	}

	for channelID, users := range requests {

		var mention string
		for _, user := range users {
			mention += "<@" + user + "> "
		}

		content := fmt.Sprintf("%s오야붕, 여기 새 게시물 %d개야!", mention, len(posts))

		send := discordgo.MessageSend{
			Content: content,
			Embed:   &embed,
		}
		msg, err := dg.ChannelMessageSendComplex(channelID, &send)
		if err != nil {
			log.Println("Error Replying\nMsg:", msg, "\nErr:", err)
		}

	}
}
