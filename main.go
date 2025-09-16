package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// File path to write
	path := "output.txt"

	for {
		fmt.Print("kkDB >> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == ".exit" {
			fmt.Println("Bye 👋")
			break
		}

		err := SaveData1(path, []byte(input))
		if err != nil {
			fmt.Println("Error saving data: ", err)
		} else {
			fmt.Println("Data saved successfully")
		}

	}

}

func SaveData1(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		return err
	}
	return fp.Sync()
}
