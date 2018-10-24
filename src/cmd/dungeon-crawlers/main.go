package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

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

	var age int
	fmt.Println("What is your age?")
	fmt.Scan(&age)

	fmt.Printf("You're %v years old", age)

	time.Sleep(time.Second * 1)
}
