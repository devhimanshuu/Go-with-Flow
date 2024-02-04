package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func PerformGetRequest(){
	const myurl="http://localhost:8000/get"
	response,err:=http.Get(myurl)
	if err!=nil{
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code:",response.StatusCode)
	fmt.Println("Content length is :",response.ContentLength)

	var responseString strings.Builder
	content,_:=ioutil.ReadAll(response.Body)
	bytecount,_:=responseString.Write(content)

	// fmt.Println(content)
	// fmt.Println(string(content))
	fmt.Println("ByteCount is :",bytecount)
	fmt.Println(responseString.String())
}
func main() {
	fmt.Println("Get Request in Golang")
	PerformGetRequest()
}