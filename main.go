package main

import (
	"net/url"

	"fmt"

	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
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

	stream := api.PublicStreamFilter(url.Values{
		"track": []string{"#love"},
	})

	defer stream.Stop()
	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			logrus.Warningf("encountered unexpected value of type %T \n", v)
			continue
		}
		fmt.Printf("%s \n", t.Text)
	}
}
