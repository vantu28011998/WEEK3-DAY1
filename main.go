package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	os.Setenv("ChannelAccessToken", "/u21Pj0z1BxGz3vfQCAHOtkG86rAEdJbJ58XorCN/7XSM6SEeETxTrSUvbv0ogXZL1NgEs0BPvV4N5R9ShbmeKOphbpuME2T62dccRm7oaOfzDtRvHExqxTlIR/HKAmmbd4B73OumlrTEIxhDUjlwwdB04t89/1O/w1cDnyilFU=")
	os.Setenv("ChannelSecret", "f3a1d4ecd5bcd7bd69b834472564071e")
	bot, err := linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("CONNECTION OK")
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		fmt.Print("REQUEST OK")
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				case *linebot.StickerMessage:
					replyMessage := fmt.Sprintf(
						"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}

}
