package main

import "fmt"

func main()  {
	fmt.Println("Slices tutorial")
	
	var arr = []string{"a","c","b"}

	fmt.Println("Fruitlist", arr);


	arr = append(arr, "ad","bc");

	fmt.Println("Appended", arr);

	arr = append(arr[:2]);

	fmt.Println("Sliced",arr);
}