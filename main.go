package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("kkDB >> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == ".exit" {
			fmt.Println("Bye 👋")
			break
		}
		fmt.Println("You typed:", input)
	}

}
