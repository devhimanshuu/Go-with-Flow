package main

import "fmt"

func main(){
	fmt.Println("Welcome to array in Golang")

	var fruitList [4]string

	fruitList[0]="Apple"
	fruitList[1]="PineApple"
	fruitList[3]="Peach"
	
	fmt.Println("Fruit list are: ",fruitList)
	fmt.Println("Fruit list are: ",len(fruitList))
	var vegList = [30]string{"potato","beans" ,"mushroom"}
	fmt.Println(vegList)
	fmt.Println(len(vegList))
}