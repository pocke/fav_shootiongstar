package main

import (
	"github.com/mrjones/oauth"
	"os"
)

var tokens = make(map[string]*oauth.RequestToken)

var twitter = func() *oauth.Consumer {
	ck := os.Getenv("consumer_key")
	cs := os.Getenv("consumer_secret")

	return oauth.NewConsumer(
		ck,
		cs,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)
}()
