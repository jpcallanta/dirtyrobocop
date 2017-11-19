package main

import ()

func main() {
	pullConfig()
	twitter := initTwitter()

	// Main loop forever!
	for {
		snore(config.SleepMin, config.SleepMax)
		tweetSomething(twitter)
	}
}
