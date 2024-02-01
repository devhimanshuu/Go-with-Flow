package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("Switch case in Golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber:=rand.Intn(6)+1
	fmt.Println("Value of dice is ",diceNumber)

	switch diceNumber{
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spots")
	case 3:
		fmt.Println("you can move 3 spots")
		fallthrough
	case 4:
		fmt.Println("you can move 4 spots")
		fallthrough
		// use to execute next case also
	case 5:
		fmt.Println("you can move 5 spots")
	case 6:
		fmt.Println("you can move 6 spots and roll Dice again ")
    default:
	fmt.Println("what was that!")}
}