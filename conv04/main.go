package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("ENter the rating:");

	reader := bufio.NewReader(os.Stdin);

	input, _ := reader.ReadString('\n');

	ratingNum, err := strconv.ParseFloat(strings.TrimSpace(input), 64);
	
	if err != nil {
		fmt.Println("Invalid input");
	} else {
		fmt.Println("Thanks for rating. We added 1 ;) - ", ratingNum+1);
	}
}