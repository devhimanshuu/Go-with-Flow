package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Post Request in Golang ")
	PerformPostJsonRequest()
}

func PerformPostJsonRequest(){
	const myurl = "https://localhost:8000/post"

	//fake json payload

	requestBody :=strings.NewReader(`
	
	{"coursename":"go with flow",
	"price":0,
	"platform":"learnGolang"
	}`)
	response ,err:=http.Post(myurl,"application/json",requestBody)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	content ,_:=ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}