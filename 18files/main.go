package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)


func main(){

	fmt.Println("Welcome to files in golang")
	content := "Working with files in golang"

	file, err := os.Create("./mygo.txt")
	
	checkNilError(err)


	len, errFile := io.WriteString(file, content)
	checkNilError(errFile)

	fmt.Println("The content has been written into the file,",len)
	readFile("mygo.txt")
	file.Close()
 	
}


func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilError(err)

	fmt.Println("The file content is: \n",string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}