package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name string `json:"coursename"`
	Price int 
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Create Json Data in Golang")

}

func EncodeJson(){
	lcoCouse:=[]course{
		{"ReactJS Bootcamp",199,"learnCodeOnline.in","abc123",[]string{"webdev","js"}},
		{"Mern Bootcamp",1679,"learnCodeOnline.in","bcd123",[]string{"webdev","js"}},
		{"Frontend Bootcamp",149,"learnCodeOnline.in","hit123",nil},
	}

	//Package this data as json data 

	finalJson,err:=json.MarshalIndent(lcoCouse,"","\t")
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n",finalJson)
}
