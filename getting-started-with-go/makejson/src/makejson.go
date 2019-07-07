package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func promptUserData() (string, string) {
	var (
		name    string
		address string
	)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a name: ")
	if scanner.Scan() {
		name = scanner.Text()
	}
	fmt.Print("Enter an address: ")
	if scanner.Scan() {
		address = scanner.Text()
	}
	return name, address
}

func main() {
	name, address := promptUserData()
	u := make(map[string]string)
	u["name"] = name
	u["address"] = address
	userJson, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(userJson))
}
