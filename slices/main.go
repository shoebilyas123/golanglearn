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


	var courses = []string{"reactjs","javascript","html","golang","deveops", "java"};
	var delElementIndex int = 2;

	fmt.Println("Courses - ", courses);

	courses = append(courses[:delElementIndex], courses[delElementIndex+1:]...);

	fmt.Println("Courses Appended - ",courses)

}