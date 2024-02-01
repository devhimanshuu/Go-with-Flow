package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")

	languages:= make(map[string]string) 

	languages["JS"]= "JavaScript"
	languages["PY"]= "Python"
	languages["RB"]= "Ruby"

	fmt.Println("Lists of all languages: ",languages)
	fmt.Println("JS  of all languages: ",languages["JS"])

	delete(languages,"RB")
	fmt.Println("Lists of all languages: ",languages)

	//loops are interesting in golang

	for _,value:=range languages{
		fmt.Printf("for key v, value is %v\n" , value)
	}
}