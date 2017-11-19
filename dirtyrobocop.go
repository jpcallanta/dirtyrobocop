package main

import ()

func main() {
	twitter := initTwitter()

	// Main loop forever!
	for {
		snore(sleepMin, sleepMax)
		tweetSomething(twitter)
	}
}
