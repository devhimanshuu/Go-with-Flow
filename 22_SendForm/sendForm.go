package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Send Form Data in Golang ")
	PerformPostFormRequest()
}

func PerformPostFormRequest(){
	const myurl ="https://localhost:8000/postform"

	//form data

	data:=url.Values{}
	data.Add("firstname","Himanshu")
	data.Add("Lastname","Gupta")
	data.Add("Email","Himanshu@gmail.com")

	response,err:=http.PostForm(myurl,data)
	if err!=nil{
		panic(err)
	}

	defer response.Body.Close()
	content,_:=ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}