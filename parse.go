package main

import (
	"errors"
	"github.com/anaskhan96/soup"
	"strings"
)

const (
	siteBase = "http://wangun.ms.jne.kr/user/"
)

type Post struct {
	Title   string
	HasFile bool
	Link    string
}

func removeMultiSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func getPosts() (posts []Post, err error) {
	resp, err := soup.Get("http://wangun.ms.jne.kr/user/indexSub.action?codyMenuSeq=129362&siteId=wangun_ms&menuUIType=top")
	if err != nil {
		return
	}
	doc := soup.HTMLParse(resp)
	rows := doc.FindAll("tr")

	var post Post
	for _, row := range rows[2:] {
		fields := row.FindAll("td")
		if len(fields) != 6 {
			err = errors.New("length of fields is not 6")
			return
		}

		a := fields[1].Find("a")
		post.Title = removeMultiSpaces(a.Text())
		post.Link = siteBase + a.Attrs()["href"]

		post.HasFile = false
		if len(fields[5].Children()) == 3 {
			post.HasFile = true
		}

		posts = append(posts, post)
	}

	return

}
