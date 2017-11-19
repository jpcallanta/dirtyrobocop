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
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}

func searchTweets(query string, limit int, client *twitter.Client) {
	search, _, _ := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: query,
		Count: limit,
	})

	fmt.Println("Search results...")

	for _, tweets := range search.Statuses {
		fmt.Printf("[%v]: %v\n", tweets.User.ScreenName, tweets.Text)
		fmt.Println()
	}
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
