package main

import "fmt"

const LoginToken string = "randomValue"//public variable

// jwtToken :=30000   outside they are not allowed in no var style
func main() {
	var username string = "Himanshu Gupta"
	fmt.Println(username)
	fmt.Printf("variable is of type:%T \n",username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type:%T \n",isLoggedIn)


	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type:%T \n",smallVal)

	var smallFloat float64 = 255.56278467426468764
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type:%T \n",smallFloat)



	//default values and some aliases
	var anotherVarible int 
	fmt.Println(anotherVarible)
	fmt.Printf("Variable is of type : %T \n",anotherVarible)

	//implicit type 
	var website = "Dev.dev"
	fmt.Println(website)

	//no var style4	

	numberOfuser :=300000
	fmt.Println(numberOfuser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n",LoginToken)
}

// uint8       the set of all unsigned  8-bit integers (0 to 255)
// uint16      the set of all unsigned 16-bit integers (0 to 65535)
// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

// int8        the set of all signed  8-bit integers (-128 to 127)
// int16       the set of all signed 16-bit integers (-32768 to 32767)
// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

// float32     the set of all IEEE-754 32-bit floating-point numbers
// float64     the set of all IEEE-754 64-bit floating-point numbers

// complex64   the set of all complex numbers with float32 real and imaginary parts
// complex128  the set of all complex numbers with float64 real and imaginary parts

// byte        alias for uint8
// rune        alias for int32