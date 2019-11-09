package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type searchResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []article `json:"articles"`
}

type article struct {
	Sources     source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//GetGoogleNews returns the search result on basis of keyword
func GetGoogleNews(searchText string) string {
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&sortBy=popularity&apiKey=1c858444e19547d9b8c5904c50f64816", searchText)
	response, err := http.Get(url)
	var searchRes searchResponse
	if err != nil {
		return fmt.Sprintf("error retrieving data %s", err)
	} else {
		//read response
		data, _ := ioutil.ReadAll(response.Body)
		//parse response to object
		json.Unmarshal(data, &searchRes)
		title := searchRes.Articles[0].Title
		url := searchRes.Articles[0].URL
		return fmt.Sprintf("Here is the first result of your serach with title: %s. You can follow the news at: %s", title, url)
	}

}
