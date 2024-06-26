package main

import (
	"github.com/somye-v01/discord-gpt-bot/config"
	"github.com/somye-v01/discord-gpt-bot/bot"
	"fmt"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()
	// client := cohereclient.NewClient(cohereclient.WithToken(config.AuthToken))
	
	<-make(chan struct{})
	return
}
