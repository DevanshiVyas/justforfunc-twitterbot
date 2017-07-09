package main

import (
	"net/url"

	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
)

var (
	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("could not get value for" + name)
	}
	return v
}

func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	log := &logger{logrus.New()}
	api.SetLogger(log)

	stream := api.PublicStreamFilter(url.Values{
		"track": []string{"#golang"},
	})

	defer stream.Stop()
	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			log.Warningf("encountered unexpected value of type %T \n", v)
			continue
		}

		if t.RetweetedStatus != nil {
			continue
		}

		_, err := api.Retweet(t.Id, false)
		if err != nil {
			log.Errorf("Could not retweet %d , with error %v", t.Id, err)
			continue
		}
		log.Infof("Retweeted %d", t.Id)
	}
}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{}) {
	log.Error(args...)
}

func (log *logger) Criticalf(s string, args ...interface{}) {
	log.Errorf(s, args...)
}

func (log *logger) Notice(args ...interface{}) {
	log.Info(args...)
}

func (log *logger) Noticef(s string, args ...interface{}) {
	log.Infof(s, args...)
}
