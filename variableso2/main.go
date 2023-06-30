package main

import "fmt"

// This is how we define a public variable; By declaring a constant or variable with first letter capital.
const PublicVariable  string = "taksdhkasjhdkjashdkjh";

func main(){
	var username string = "myuser";
	fmt.Println(username);
	fmt.Printf("The variable is of type: %T \n", username)


	
	var isLoggedIn bool = true;
	fmt.Println(isLoggedIn);
	fmt.Printf("The variable is of type: %T \n", isLoggedIn); // true

	var emptyString string;
	fmt.Println(emptyString);
	fmt.Printf("The type of variable is: %T\n", emptyString); // emptyString


	// Implicit Typing

	var dyString = "string";
	fmt.Println(dyString);
	fmt.Printf("The type of variable is: %T \n", dyString); // dyString

	// No var style
	// This walrus operator instead of var is allowed only inside functions. They are not allowed outside of the method. 
	numberOfUsers := 5000.0;
	fmt.Println(numberOfUsers);

	fmt.Println(PublicVariable); //
}