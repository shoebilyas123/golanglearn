package main

import "fmt"




func main(){
	fmt.Println("Structs in golang")
	user := User{"User","user@gmail.com",true,14}


	fmt.Println(user.Name);
	fmt.Printf("The user details are: %+v\n",user);
	user.getStatus()
	user.changeEmail("newemail@gmail.com")
	fmt.Println(user.Email);
	//No inheritance in golang; no super or parent
}


type User struct {
	Name string
	Email string
	Status bool
	Age int

}

func (u User) getStatus() {
	fmt.Println("The user is active: ", u.Status)
}


func (u *User) changeEmail(_Email string) {
	u.Email = _Email
	fmt.Println("The new user email is: ", u.Email)
}