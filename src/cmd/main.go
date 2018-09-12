package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/damienfamed75/dungeon-crawlers/src/common"
)

func main() {
	common.Clear()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Hello, Adventurer!\n" +
		"What is your name?")
	fmt.Print("> ")

	username, _ := reader.ReadString('\n')

	common.Clear()

	fmt.Println("Hello," + username)
}
