// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kynwu/trafficbot"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

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
				trafficEvents, _ := trafficbot.GetTrafficEvents()
				switch message.Text {

				case "time":

					homeToAirport := trafficbot.GetDurationFromHomeToAirport();
					airportToHome := trafficbot.GetDurationFromAirportToHome();
					s := fmt.Sprintf("from home to airport: %s \n from airport to home: %s", homeToAirport, airportToHome);
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(s).Do(); err != nil {
						log.Print(err)
					}
				case "1001":
					var responseText string
					for _, element := range trafficEvents.FormData {
						if element.Highway == "1001" && element.Region == "N" {
							responseText += element.Comment + "\n"
						}
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(responseText)).Do(); err != nil {
						log.Print(err)
					}
				case "1002":
					var responseText string
					for _, element := range trafficEvents.FormData {
						if element.Highway == "1002" && element.Region == "N" {
							responseText += element.Comment + "\n"
						}
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(responseText)).Do(); err != nil {
						log.Print(err)
					}
				case "1003":
					var responseText string
					for _, element := range trafficEvents.FormData {
						if element.Highway == "1003" && element.Region == "N" {
							responseText += element.Comment + "\n"
						}
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(responseText)).Do(); err != nil {
						log.Print(err)
					}
				default:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("請輸入國道編號(例如：國道一號為1001)")).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}
