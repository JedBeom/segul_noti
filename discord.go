package main

import (
	"log"

	disgo "github.com/bwmarrin/discordgo"
)

var (
	dg *disgo.Session
)

func init() {
	discordInit()
}

func discordInit() {
	var err error
	dg, err = disgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal("Discord Session:", err)
		return
	}

	dg.AddHandler(handler)

	err = dg.Open()
	if err != nil {
		log.Fatal("Open Session:", err)
		return
	}

	log.Println("Bot Opened.")

	err = dg.UpdateListeningStatus("~sub or ~posts")
	if err != nil {
		log.Println("Update Listening Status:", err)
	}

}

func addField(fields *[]*disgo.MessageEmbedField, name, value string) {
	*fields = append(*fields, &disgo.MessageEmbedField{
		Name:   name,
		Value:  value,
		Inline: false,
	})
}
