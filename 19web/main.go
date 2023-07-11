package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


const url = "https://lco.dev"

func main(){
	fmt.Println("Web Requests");
	
	res, err := http.Get(url)


	checkForError(err)

	fmt.Printf("The response is: %T\n",res)

	defer res.Body.Close();

	databytes, errBody := ioutil.ReadAll(res.Body)

	checkForError(errBody);

	fmt.Printf("The response body: %v\n",string(databytes))
}


func checkForError(err error){
	if err != nil {
		panic(err)
	}
}