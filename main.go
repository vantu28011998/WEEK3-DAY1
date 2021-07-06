package main

import (
	"fmt"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	os.Setenv("ChannelAccessToken", "WaWrpzl9GkEBbMl0nMgAi7e+hqiUwvlYC6x1Uv1fbqMOBqEINdbcB4elBy8PJAfN2V+gs2eEkZtPvrrbuhrNm8TrmDf+l1BE2jdIBDPzVJMtcooXSRbax1PBqvScLAmAbxN/abn+NKUJ92nxUHN2yAdB04t89/1O/w1cDnyilFU=")
	os.Setenv("ChannelSecret", "9208ccd3e9780d150c0aa8f009fff9a2")
	os.Setenv("PORT", "8080")
	bot, err := linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("SUCCESSFULLY CREATED LINE BOT")

	if _, err := bot.PushMessage("Ua3c33fd6a14f54421fbd1835cba1a487", linebot.NewTextMessage("Hello Da Nang")).Do(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("API OK")

}
