package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {
	var Nameforfile string
	var Infoinfile string
	var startinput string
	fmt.Println("Hello USERS!")
	for {
		fmt.Println("Enter start to start((\n And any word to exit")

		fmt.Scan(&startinput)
		startinput2 := strings.ToLower(startinput)
		if startinput2 == "start" {
			// add file code:
			fmt.Println("Enter name for file:")
			fmt.Scan(&Nameforfile)
			fmt.Println("Name add succesfuly")
			fmt.Println("Enter infomation into youre file:")
			// special scan for info in our file
			reader := bufio.NewReader(os.Stdin)
			Infoinfile, _ = reader.ReadString('\n')
			Infoinfile = strings.TrimSpace(Infoinfile)
			fmt.Println("Information add succesfuly")
			// informing the user about the date the file was added.
			log.Println("Data of added file,", Nameforfile)
			// add file into downloads directory
			FILE(Nameforfile, Infoinfile)
			fmt.Println("Thanks for using our program!!!")
		} else {
			// exit program
			os.Exit(0)
		}
	}
}

func FILE(FILEName string, INFO string) {
	userInfo, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return
	}
	// join into Downloads directory
	downloadsPath := filepath.Join(userInfo.HomeDir, "Downloads")

	filepath := filepath.Join(downloadsPath, FILEName)

	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error creating file!", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(INFO)
	if err != nil {
		fmt.Println("Error in writing info in file!", err)
		return
	}
}
