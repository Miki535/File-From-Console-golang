package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	var enterscan string
	var FileName string
	var Info string
	fmt.Println("Hello USERS!\n Write start to start(((")
	fmt.Scan(&enterscan)
	if enterscan == "start" {
		fmt.Println("Enter name of youre file:")
		fmt.Scan(&FileName)
		fmt.Println("Your file name add successfully!")

		fmt.Println("Now write INFO in youre file:")
		fmt.Scan(&Info)
		FILE(FileName, Info)
		fmt.Println("To restart enter start, to exit enter exit")
	} else if enterscan == "exit" {
		os.Exit(0)
	}
}

func FILE(FILEName string, INFO string) {
	userInfo, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return
	}

	downloadsPath := filepath.Join(userInfo.HomeDir, "Downloads")

	filepath := filepath.Join(downloadsPath, FILEName)

	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(INFO)
	if err != nil {
		return
	}
}
