package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	days:=[]string{"Sunday","Monday","Tuesday","Wednesday","friday","Saturday"}
	fmt.Println(days)

	//Loop in golang
	for d:=0;d<len(days);d++{
		fmt.Println(days[d])
	}

	for i:=range days{
		fmt.Println(days[i])
	}

	for index,day:=range days{
		fmt.Printf("index is %v and value is %v\n", index , day)
	}

	rougueValue :=1

	for rougueValue<10{
		if rougueValue==2{
			goto lco
		}
		if rougueValue==5{
			// break
			rougueValue++
			continue
		}
		fmt.Println("value is ",rougueValue)
		rougueValue++
	}

	lco:
	fmt.Println("learning Golang")
}