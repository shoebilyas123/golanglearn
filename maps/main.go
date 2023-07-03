package main

import "fmt"


func main()  {
	courses := make(map[string]string);

	courses["JS"] = "Javascropt";
	courses["RJS"] = "ReactJS";
	courses["PY"] = "Python";


	fmt.Println("Courses - ", courses);
	delete(courses, "PY")
	fmt.Println("Courses - ", courses);
	
	fmt.Println("Courses JS - ", courses["JS"]);


	for key,val := range courses {
		fmt.Printf("Value for key %v is %v\n",key, val);
	}
}