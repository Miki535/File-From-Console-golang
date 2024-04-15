package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {
	var Nameforfile string
	var Infoinfile string

	fmt.Println("Enter name for file:")
	fmt.Scan(&Nameforfile)
	fmt.Println("Name add succesfuly")
	fmt.Println("Enter infomation into youre file:")
	reader := bufio.NewReader(os.Stdin)
	Infoinfile, _ = reader.ReadString('\n')
	Infoinfile = strings.TrimSpace(Infoinfile)
	fmt.Println("Information add succesfuly")
	FILE(Nameforfile, Infoinfile)
	fmt.Println("Thanks for using our program!!!")
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
		fmt.Println(err)
		return
	}
}
