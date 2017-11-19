package main

import (
	"os"
)

func main() {
	pullConfig(os.Args[1])
	twitter := initTwitter()

	// Main loop forever!
	for {
		snore(config.SleepMin, config.SleepMax)
		tweetSomething(twitter)
	}
}
