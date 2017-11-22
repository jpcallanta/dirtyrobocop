package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: dirtyrobocop <config_path> <runonce|runforever>")
		os.Exit(1)
	}

	pullConfig(os.Args[1])

	twitter := initTwitter()

	switch os.Args[2] {
	case "runonce":
		tweetSomething(twitter)

	case "runforever":
		// Main loop forever!
		for {
			snore(config.SleepMin, config.SleepMax)
			tweetSomething(twitter)
		}
	}
}
