/*

mybot - Illustrative Slack bot in Go

Copyright (c) 2015 RapidLoop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: mybot slack-bot-token\n")
		os.Exit(1)
	}

	// start a websocket-based Real Time API session
	ws, id := slackConnect(os.Args[1])
	fmt.Println("mybot ready, ^C exits")

	for {
		// read each incoming message
		m, err := getMessage(ws)
		if err != nil {
			log.Fatal(err)
		}

		// see if we're mentioned
		if m.Type == "message" && strings.HasPrefix(m.Text, "<@"+id+">") {
			// if so try to parse if
			parts := strings.Fields(m.Text)

			switch parts[1] {
			case "stock":
				if len(parts) == 3 {
					go func(m Message) {
						m.Text = getQuote(parts[2])
						postMessage(ws, m)
					}(m)
				} else {
					// huh?
					m.Text = fmt.Sprintf("sorry, that does not compute\n")
					postMessage(ws, m)
				}
			case "news":
				go func(m Message) {
					var searchText string
					for i := 2; i < len(parts); i++ {
						searchText = searchText + "%20" + parts[i]
					}
					m.Text = GetGoogleNews(searchText)
					postMessage(ws, m)
				}(m)
			case "twitterPost":
				go func(m Message) {
					var tweetText string
					for i := 2; i < len(parts); i++ {
						tweetText = tweetText + " " + parts[i]
					}
					m.Text = SendTweet(tweetText)
					postMessage(ws, m)
				}(m)
			case "twitterSearch":
				go func(m Message) {
					var searchText string
					for i := 2; i < len(parts); i++ {
						searchText = searchText + " " + parts[i]
					}
					m.Text = SearchTweet(searchText)
					postMessage(ws, m)
				}(m)
			}

		}
	}
}
