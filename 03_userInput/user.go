package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	reader:=bufio.NewReader(os.Stdin) // library to get input from user
	fmt.Println("Enter the rating for our Pizza")

	//comma ok || err err
	input,_:=reader.ReadString('\n')
	fmt.Println("Thanks for Rating,  ",input)
	fmt.Printf("Type of Rating is %T,  ",input)
	

}