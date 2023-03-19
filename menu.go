package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var cmd string
	for true {
		fmt.Printf("Please enter your command:\n")
		fmt.Scanln(&cmd)
		if strings.EqualFold(cmd, "help") {
			fmt.Println("This is help cmd")
		} else if strings.EqualFold(cmd, "quit") {
			os.Exit(0)
		} else {
			fmt.Println("Wrong cmd!")
		}
	}
}
