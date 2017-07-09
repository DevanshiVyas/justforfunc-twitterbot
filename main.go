package main

import (
	"net/url"

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
		"track": []string{"#golang"},
	})

	defer stream.Stop()
	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			logrus.Warningf("encountered unexpected value of type %T \n", v)
			continue
		}

		if t.RetweetedStatus != nil {
			continue
		}

		_, err := api.Retweet(t.Id, false)
		if err != nil {
			logrus.Errorf("Could not retweet %d , with error %v", t.Id, err)
			continue
		}
		logrus.Infof("Retweeted %d", t.Id)
	}
}
