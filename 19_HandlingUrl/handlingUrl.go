package main

import (
	"fmt"
	"net/url"
)

const myUrl string ="https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"


func main() {
	fmt.Println("Handling URL in Golang")

	fmt.Println(myUrl)

	//parsing
	result , err:=url.Parse(myUrl)
	if err!=nil{
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams :=result.Query()
	fmt.Printf("The type of query params are:%T\n",qparams)

	fmt.Println(qparams["coursename"])

	for _,val :=range qparams{
		fmt.Println("Params is  :",val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=himanshu",
	}
	anotherUrl :=partsOfUrl.String()
	fmt.Println(anotherUrl)

	}