package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type User struct {
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Active    bool      `json:"active"`
	LastLogin time.Time `json:"lastLogin"`
	Children  []User    `json:"children"`
}

func GetUsers() []User {
	bytes, err := os.ReadFile("../users.json")
	check(err)

	var users []User
	err = json.Unmarshal(bytes, &users)
	check(err)

	return users
}

func main() {
	fmt.Println("Users:", GetUsers())
}
