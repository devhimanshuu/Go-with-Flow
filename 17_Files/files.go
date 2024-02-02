package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working With Files in Golang")
	content :="this needs to go in a file"

	file,err:=os.Create("./myfile.txt")
	if err!=nil{
		panic(err)
	}

	length , err := io.WriteString(file,content)
	if err!=nil{
		panic(err)
	}
	fmt.Println("Length is :", length)
	 defer file.Close()
	 readfile("./myfile.txt")
}
func readfile(fileName string){
	 dataByte,err:=ioutil.ReadFile(fileName)
	 if err!=nil{
		panic(err)
	}

	fmt.Println("Text Data inside the file is \n",string(dataByte))

}