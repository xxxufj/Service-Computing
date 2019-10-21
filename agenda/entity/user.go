package entity

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type User struct {
	Username    string
	Password    string
	Email       string
	PhoneNumber string
}

var users []User

func QueryUserByName(username string) bool {
	for _, user := range users {
		if user.Username == username {
			return true
		}
	}
	return false
}

func QueryUserByNameAndPassword(username string, password string) bool {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

func AddUser(newUser User) {
	users = append(users, newUser)
}

func GetExsitingUsers() {
	file, err1 := os.Open("curUser")
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Fail to open \"curUser\" and read existing users")
	}
	decoder := json.NewDecoder(file)
	err2 := decoder.Decode(&users)
	if err2 != io.EOF && err2 != nil {
		fmt.Fprintf(os.Stderr, "fail to decode \"curUser\"")
	}
	file.Close()
}

func SaveUsers() {
	file, err1 := os.Create("curUser")
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Fail to open \"curUser\" and write existing users")
	}
	encoder := json.NewEncoder(file)
	err2 := encoder.Encode(&users)
	if err2 != nil {
		fmt.Fprintf(os.Stderr, "fail to encode existing users and write to \"curUser\"")
	}
	file.Close()
}
