package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("This is a Decimal to Binary Converter")
	fmt.Println("Enter a number (must be an integer:)")
	var s string
	var i int

	_, err := fmt.Scan(&s)
	i, err = strconv.Atoi(s)
	
	if err != nil {
		fmt.Println("Enter a valid number")
	} else {
		fmt.Printf("%b\n",i)
	}
	
}