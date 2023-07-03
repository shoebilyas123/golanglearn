package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Print("Welcome to Time");


	presentTime := time.Now();

	fmt.Print("Current Time: ")
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	createdTime := time.Date(2020, time.January, 6, 12, 5, 0, 0, time.UTC);

	fmt.Print("Printing Created Time for 2020: ");
	fmt.Println(createdTime);
}