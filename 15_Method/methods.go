//methods are same as function but when function are define in struct they are called methods

package main

import "fmt"


func main(){
	fmt.Println("Welcome to structs in golang")

	//no inheritance in golang; no super or parent
	himanshu:=User{"Himanshu","devhimanshuu@gmail.com",true,22}
	fmt.Println(himanshu)
	fmt.Printf("Himanshu details are: %+v\n",himanshu)
	fmt.Printf("Name is : %+v and Email is %v",himanshu.Name,himanshu.Email+" \n")
	himanshu.GetStatus()
	himanshu.NewMail()

	fmt.Printf("Name is : %+v and Email is %v",himanshu.Name,himanshu.Email+" \n")
}
type User struct{
	Name string
	Email string
	Status bool
	Age int  
}
//Methods declarations

func (u User) GetStatus(){
	fmt.Println("Is User active: ",u.Status)
}
func (u User) NewMail(){
	u.Email="test@go.dev"
	fmt.Println("Email of this User is: ",u.Email)
}