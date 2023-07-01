package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Welcome to time.");


	presentTime := time.Now();

	fmt.Println(presentTime.UTC().Date());
}