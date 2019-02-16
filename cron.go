package main

import "gopkg.in/robfig/cron.v2"

func jobInit() (c *cron.Cron) {
	c = cron.New()
	_, err := c.AddFunc("0 */30 * * * *", notification)
	if err != nil {
		panic(err)
	}
	return
}
