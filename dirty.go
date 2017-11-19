package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"math/rand"
	"os/exec"
	"time"
)

func snore(min int, max int) {
	rand.Seed(time.Now().UnixNano())

	seconds := rand.Intn(max-min) + min
	msgBuf := fmt.Sprintf("Sleeping for %d seconds...", seconds)

	logOut(msgBuf)
	time.Sleep(time.Duration(seconds) * time.Second)
}

func logOut(msg string) {
	fmt.Printf("[%v] - %v\n", time.Now().Format(time.RFC850), msg)
}

func initTwitter() *twitter.Client {
	configAuth := oauth1.NewConfig(config.ConsumerKey, config.ConsumerSecret)
	token := oauth1.NewToken(config.AccessToken, config.AccessSecret)
	httpClient := configAuth.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}

func tweetSomething(client *twitter.Client) {
	for {
		out, _ := exec.Command("fortune", "-o").Output()

		if len(string(out)) < 140 {
			logOut(string(out))
			client.Statuses.Update(string(out), nil)
			return
		}
	}
}
