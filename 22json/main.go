package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name string `json:"courseName"`
	Price int
	Platform string
	Password string
	Tags []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON Lecture")
	EncodeJSON()
 
}


func handleError(err error) {

	if err != nil {
		panic(err)
	}
}

func EncodeJSON() {
	lcoCourses := []courses{
		{"ReactJS", 299,"learncodeonline.in", "password", []string{"Web","Javascript"}},
		{"NodeJS", 299,"learncodeonline.in", "password3", []string{"Backend","Javascript"}},
		{"Docker", 299,"learncodeonline.in", "passwor2", nil},
	}


	finalJSON, err := json.MarshalIndent(lcoCourses, "","\t")

	handleError(err)

	fmt.Printf("%s\n",finalJSON)
}