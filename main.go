package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discordgo.New("")
	fmt.Println("Hello")
}

// "Rewriting" it from Python to Go since it is way better for free/cheap hosting since it is much more performant + uses less storage and ram
// However, barrier to contribution is a bit higher so rip.
