package main

import "fmt"

func main()  {
	// for d := 0; d<10;d++{
	// 	fmt.Println(d);
	// }
	days := []string{"Sunday","Monday","Tuesday"};

	for i := range days {
		if i ==0 {
			goto rest
		} else {
			break;
		}

		fmt.Println(days[i]);	}

		rest:
			fmt.Printf("Today is Sunday and a rest day!\n")
}