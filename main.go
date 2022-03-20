package main

import "fmt"

func main() {
	var cmd string
	fmt.Println("Welcome to Menu!")
	for {
		fmt.Scanln(&cmd)
		switch cmd {
		case "version":
			fmt.Println("Menu 1.0")
		case "hello":
			fmt.Println("world")
		case "help":
			fmt.Println("Try hello")
		case "quit":
			fmt.Println("bye")
			return
		default:
			fmt.Println("emmmm...I can't understand.")
		}
	}
}
