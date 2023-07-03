package main

import "fmt"

func main()  {
	var fruitlist [4]string
	fruitlist[0] = "apple"
	fruitlist[1] = "banana"
	// fruitlist[2] = "orange"
	fruitlist[3] = "tomato"
	// fruitlist[4] = "grapes"

	fmt.Println("Fruitlist is: ", fruitlist);
	fmt.Println("Fruitlist length is: ", len(fruitlist));


	var veglist = [5]string{"a","b","c"};
	fmt.Println("Veg list", veglist);
}