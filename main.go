package main

import (
	"fmt"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	os.Setenv("ChannelAccessToken", "FE+w3+zruEZ0e91TbusoHnOY0pUfa4kY5u4CGpqjoDybwRFy3JphtzNt7S6ciV2NCvuhNktPcZjOa6SrO/6dihlQ+Sd1b8dfdxkvsCEn2Y3P/HdyemMawGln8WBZWY9F7MPw/El3G5EQ3HqMTd2G2wdB04t89/1O/w1cDnyilFU=")
	os.Setenv("ChannelSecret", "5e6ce8d78756dc46b6f75f5aefbad7b0")
	os.Setenv("PORT", "8080")
	bot, err := linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("SUCCESSFULLY CREATED LINE BOT")

	if _, err := bot.PushMessage("U1bf6bc60d2e9e9fd5b71cd01e7db10a6", linebot.NewTextMessage("Hello !!! This is Go team's bot.")).Do(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("API OK")

}
