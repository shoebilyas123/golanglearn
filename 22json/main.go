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
	// EncodeJSON()
	DecodeJSON()
 
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

func DecodeJSON() {
	jsonDataFromWeb :=  []byte(`{
		"courseName": "Docker",
		"Price": 299,
		"Platform": "learncodeonline.in",
		"Password": "passwor2",
		"tags":["web","json"]
}`)

		var lcoCourse courses
		
		isJSONValid := json.Valid(jsonDataFromWeb)

		if isJSONValid {
			fmt.Println("JSON is Valid")
			
			json.Unmarshal(jsonDataFromWeb, &lcoCourse)
			fmt.Printf("%+v\n",lcoCourse)

			var myData map[string]interface {}
			
			json.Unmarshal(jsonDataFromWeb,&myData)
			fmt.Printf("%#v\n",myData)
			
		} else {
			fmt.Println("JSON is not Valid")
			
		}

}