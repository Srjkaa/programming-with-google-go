package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func printNames(users []Name) {
	for _, user := range users {
		fmt.Printf("first name: %s, last name: %s\n", user.fname, user.lname)
	}
}

func splitFileToLines(file []byte, err error) ([]string, error) {
	fileLines := strings.Split(string(file), "\n")
	return fileLines, err
}

func readFile(fileName string) ([]byte, error) {
	file, err := ioutil.ReadFile(fileName)
	return file, err
}

func readUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	var fileName string
	if scanner.Scan() {
		fileName = scanner.Text()
	}
	return fileName
}

func transformLinesToNameStructure(fileLines []string) []Name {
	names := make([]Name, 0)
	for _, line := range fileLines {
		splittedLine := strings.Split(line, " ")
		fname := splittedLine[0]
		lname := splittedLine[1]
		names = append(names, Name{ fname: fname, lname: lname })
	}
	return names
}

func main() {
	names := make([]Name, 0)
	fmt.Print("Enter name of the text file: ")
	fileName := readUserInput()

	fileLines, err := splitFileToLines(readFile(fileName))
	if err != nil {
		fmt.Println(err)
	}

	names = transformLinesToNameStructure(fileLines)
	printNames(names)
}
