package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to videop on slices ")

	var fruitList = []string{"Apple","Tomato","Peach"}
	fmt.Printf("Type of fruitList %T\n" ,fruitList)
	// to know about type of data structure 

	fruitList = append(fruitList,"Mango" , "banana")
	//to add the data in slices 

	fmt.Println(fruitList)

	
	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScore :=make([]int,4)

	highScore[0]=243;
	highScore[1]=2748;
	highScore[2]=657;
	highScore[3]=937;
	// highScore[4]=90;

	highScore = append(highScore, 555,2723,8474)
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//How to remove a value from sluces based on index
	
	var courses =[]string{"reactjs","javascript","swift","python","ruby","rust"}
	fmt.Println(courses)
	var index int = 2
	courses= append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)

}