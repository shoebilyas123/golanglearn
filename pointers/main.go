package main

import "fmt"


func main(){
	fmt.Println("Welcome to Pointers");

	var ptr *int;

	fmt.Println("This is a pointer value, ", ptr);


	number := 23;
	var ptr2 = &number;

	fmt.Println("This is a poiinter address: ", ptr2);
	fmt.Println("This is a poiinter address: ", *ptr2);

	*ptr2 = *ptr2 +2;
	fmt.Println(*ptr2)
	fmt.Println("Number: ", number);
}