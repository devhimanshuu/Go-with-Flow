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
	DecodeJson()
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

func DecodeJson(){
	jsonDataFromWeb:=[]byte(`{
		"coursename":"ReactJS Bootcamp",
		"Price":122,
		"website":"LearnCodeOnline.in",
		"tags":[
			"web-dev",
			"js"
		]
	}`)

	var lcoCouse course

	checkValid:=json.Valid(jsonDataFromWeb)

	if checkValid{
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb,&lcoCouse)
		fmt.Printf("%#v\n",lcoCouse)
	}else{
		fmt.Println("Json was not valid")
	}

	//some cases where you just want to add data to key value

	var mydata map[string]interface{}
	json.Unmarshal(jsonDataFromWeb,&mydata)
	fmt.Printf("%#v\n",lcoCouse)

	for k,v:=range mydata{
		fmt.Printf("key is %v and value is %v and type is %T\n",k,v,v)
	}
}