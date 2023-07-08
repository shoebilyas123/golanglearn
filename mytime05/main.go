package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Welcome to time.");


	presentTime := time.Now();

	fmt.Println(presentTime.UTC().Date());

	fmt.Println(presentTime.Format("2006-01-02 15:04:05"));
}