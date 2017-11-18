package main

import (
	"fmt"
	"math/rand"
	"time"
)

func snore(min int, max int) {
	rand.Seed(time.Now().UnixNano())

	seconds := rand.Int(max-min) + min

	time.Sleep(seconds * time.Second)
}
