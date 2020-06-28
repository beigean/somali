package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

type auth struct {
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}

func main() {
	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile("./.confidencial/channel.json")
	if err != nil {
		log.Fatal(err)
	}
	// JSONデコード
	var auth auth
	if err := json.Unmarshal(bytes, &auth); err != nil {
		log.Fatal(err)
	}
	// デコードしたデータを表示
	fmt.Printf("secret: %s \n", auth.Secret)
	fmt.Printf("access_token: %s\n", auth.AccessToken)

	bot, err := linebot.New(auth.Secret, auth.AccessToken)

	fmt.Println(bot)
}
