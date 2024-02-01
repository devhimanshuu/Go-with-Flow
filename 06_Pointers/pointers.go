package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers")

	// var ptr *int
	// fmt.Println("value of pointer ",ptr);

	myNumber :=23
	
	var ptr=&myNumber 
	fmt.Println("value of actual pointer is ",ptr) //memeory address
	fmt.Println("value of actual pointer is ",*ptr) //get actual value

	*ptr = *ptr*2
	fmt.Println("New value is ",myNumber)
}