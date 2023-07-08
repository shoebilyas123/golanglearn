package main

import "fmt"




func main(){
	fmt.Println("Structs in golang")
	user := User{"User","user@gmail.com",true,14}


	fmt.Println(user.Name);
	fmt.Printf("The user details are: %+v\n",user);
	//No inheritance in golang; no super or parent
}


type User struct {
	Name string
	Email string
	Status bool
	Age int
}