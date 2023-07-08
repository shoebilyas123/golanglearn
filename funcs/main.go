package main

import "fmt"

func main(){
	fmt.Println("Functions lessio");

	value := addNums(5,6);
	value1 := addAllNums(0, 5,4,21,3,1,2,213,1,32,3,2,32)	
	greet();

// func() {
// 	fmt.Println("Hello world anonymous");
// }()
  ANONYM();

	fmt.Printf("5+6=%v\n",value);
	fmt.Printf("PRO ADD=%v\n",value1);
}


func addAllNums(total int,values ...int) int {

	for _,val := range values {
		total += val
		
	}

	return total
}


func ANONYM() func()  {
	return func() { fmt.Printf("HELLOWRODL\n");}
}


func greet(){
	fmt.Println("Hello world greeting")
}

func addNums(a int, b int) int{
	return a+b
}