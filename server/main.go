package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	PerformPostJsonRequest("http://localhost:8000/post")
	// fmt.Printf("Request Response : %v\n",rer)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func PerformWebRequest(url string) string {
	response, err := http.Get(url);
	handleError(err)
	
	var responseString strings.Builder
	databyte, _ := ioutil.ReadAll(response.Body)

	defer response.Body.Close();

	fmt.Printf("Status : %v\n", response.Status)
	fmt.Printf("Content Length : %v\n", response.ContentLength)

	byteCount, _ := responseString.Write(databyte)

	fmt.Printf("Byte Count : %v\n",byteCount)
	fmt.Println(responseString.String())

	return string(databyte);
	
}


func PerformPostJsonRequest(url string) {
	requestBody := strings.NewReader(`
	{
		"coursename":"Let's gow tih golang",
		"price":23,
		"platform":"LearnCodeOnline.in"
	}
	`)

	response, err := http.Post(url, "application/json", requestBody)
	defer response.Body.Close()

	handleError(err);

	databyte, _ :=  ioutil.ReadAll(response.Body);
	var stringBuilder strings.Builder
	byteCount, _ := stringBuilder.Write(databyte)
	fmt.Println(byteCount)
	fmt.Println(stringBuilder.String());
}