package main

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
)

const (
	consumerKey       = "WuWtVBneDIFi8PetwWejZjw5C"
	consumerSecret    = "tCV2bjdlvFNowl8e9mwMYb4UrNj7LjXqpltERZZn3JZSOLfrsM"
	accessToken       = "705491955774595072-pIWKpOYm7iK8fzhqqLtv0h5ZlPNUl18"
	accessTokenSecret = "BYaoKxaXl60rdcO98XpXzUKDmj6fefJQGrvDTdxKkqXuk"
)

func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	res, _ := api.GetSearch("golang", nil)
	for _, tweet := range res.Statuses {
		fmt.Print(tweet.Text)
	}
}
