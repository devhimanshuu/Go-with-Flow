package main

import "fmt"

func main() {
	fmt.Println("Defer in Golang ")
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	//this line gets cut out and execute just before last curly braces 
	fmt.Println("hello")
	myDefer()
	
}

//Output
/*
Defer in Golang
hello
4
3
2
1
0
Two
One
World
*/

func myDefer(){
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}