package main

import "fmt"


func main(){
	fmt.Println("Welcome to structs in golang")

	//no inheritance in golang; no super or parent
	himanshu:=User{"Himanshu","devhimanshuu@gmail.com",true,22}
	fmt.Println(himanshu)
	fmt.Printf("Himanshu details are: %+v\n",himanshu)
	fmt.Printf("Name is : %+v and Email is %v",himanshu.Name,himanshu.Email)
}
type User struct{
	Name string
	Email string
	Status bool
	Age int  
}