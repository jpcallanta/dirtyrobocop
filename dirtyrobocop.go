package main

import (
	"fmt"
)

func main() {
	// Main loop forever!
	for {
		snore(300, 1800)
		fmt.Println("Insert tweet here...")
	}
}
