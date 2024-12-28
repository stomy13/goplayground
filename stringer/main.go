package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	user := User{
		UUID:     "DD878E0D-7BC7-4198-84A1-B9B7A1812032",
		ID:       1,
		Name:     "John Doe",
		Password: "password",
	}
	fmt.Println(user)
	fmt.Printf("%v\n", user)
	fmt.Printf("%#v\n", user)
	fmt.Printf("%+v\n", user)
	fmt.Println(user.Password)

	bytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

type User struct {
	UUID     string
	ID       int
	Name     string
	Password Password
}

type Password string

func (p Password) String() string {
	return "********"
}

func (p Password) GoString() string {
	return "********"
}
