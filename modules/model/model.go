package model

import (
	"fmt"
)

func init () {
	fmt.Println("model package is imported")
}

type User struct {
	Name string
	Age int
}

func (user *User) UpdateName(newName string) {
	oldName := user.Name
	user.Name = newName
	fmt.Printf("Update Name; %v --> %v\n", oldName, newName)
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}


type Dimensions struct {
	Width int `json: "width"`
	Height int `json: "height"`
}

type Data struct {
	Species string `json: "species"`
	Description string `json: "description"`
	Dimensions Dimensions `json: "dimensions"`
}

