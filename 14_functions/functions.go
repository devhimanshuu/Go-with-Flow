package main

import "fmt"

func main() {
	fmt.Println("Function in Golang")
	greeter()
	greeter2()
	//it is not allwed to meke a function inside a function itself

	result:=add(4,5)
	fmt.Println("Result is ",result)

	
}
func greeter2(){
	fmt.Println("Another function ")
}
func greeter(){
	fmt.Println("it's Golang bro")
}
func add(a int , b int ) int{
	return a+b
}
