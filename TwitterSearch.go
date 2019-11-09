package main

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Credentials struct does something. idk
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

//SearchResponse is response object for search tweet
type SearchResponse struct {
	Status   int
	Text     string
	FullText string
}

//SendTweetResponse is response for sending tweet
type SendTweetResponse struct {
	Status int
	ID     int64
}

func getClient(creds *Credentials) (*twitter.Client, error) {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}
	return client, nil
}
func getCredentials() Credentials {
	return Credentials{
		AccessToken:       "125277653-iwDs38SZcAwYVQI76zn3O1MAHDesk0shEgSFvzKZ",
		AccessTokenSecret: "lgRCdt5ftq5xrZq3CvKjsYhkPglkkW9sXD2Le2tnZXNpF",
		ConsumerKey:       "xpedwliEFGhYQKNZFtG6Mm7BR",
		ConsumerSecret:    "texBxnSkmddS0a32iQLXmLttVJkXnrfmiB5YlmXg8BPcMUjChb",
	}
}

//SendTweet function sends tweet for authenticated user
func SendTweet(tweetText string) string {
	creds := getCredentials()
	var clnt, err = getClient(&creds)
	if err != nil {
		return fmt.Sprintf("error retrieving client %s", err)
	}
	tweet, resp, err := clnt.Statuses.Update(tweetText, nil)
	if err != nil || resp.StatusCode != 200 {
		return fmt.Sprintf("Tweeting failed %s", err)
	}
	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)
	return fmt.Sprintf("Successfully tweeted %+v\n %+v\n", resp, tweet)
}

//SearchTweet searches a tweet or hashtag
func SearchTweet(searchText string) string {
	creds := getCredentials()
	var clnt, err = getClient(&creds)
	if err != nil {
		return fmt.Sprintf("error retrieving client %s", err)
	}
	var test = twitter.SearchTweetParams{
		Query:      searchText,
		Count:      1,
		ResultType: "popular",
		Since:      "2010-01-01",
	}
	search, resp, err := clnt.Search.Tweets(&test)
	if err != nil {
		log.Println(err)
	}
	var searchResponse = SearchResponse{
		FullText: search.Statuses[0].FullText,
		Status:   resp.StatusCode,
		Text:     search.Statuses[0].Text,
	}
	return fmt.Sprintf("%s %s", searchResponse.Text, searchResponse.FullText)
}
