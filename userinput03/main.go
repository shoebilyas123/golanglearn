package main

import (
	"bufio"
	"fmt"
	"os"
)

// Comma OK Syntax;

func main(){	
	welcome := "Welcome to userinput";

	fmt.Println(welcome);

	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter the rating for our pizza: ");

	input, _ := reader.ReadString('\n');

	fmt.Printf("Thanks for rating our pizza: %s \n", input);

}