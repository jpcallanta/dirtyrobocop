package main

import (
	"fmt"
	"math/rand"
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
