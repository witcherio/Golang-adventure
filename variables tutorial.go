package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Variables declaration - 1 type
	var variableOne int
	variableOne = 24

	fmt.Println("First variable is:", variableOne)
	fmt.Println("Type of 1st variable is:", reflect.TypeOf(variableOne)) // int

	// Variables declaration - 2 type
	variableTwo := "I'm second variable"
	fmt.Println("Second variable is:", variableTwo)
	fmt.Println("Type of 2nd variable is:", reflect.TypeOf(variableTwo)) // string

	// Const variable can't be changed after declaration, it's not dynamic like in a Python / JS and etc
	const variableThree = float64(524.4)
	fmt.Println(variableThree)
	fmt.Println("Type of const variable:", reflect.TypeOf(variableThree)) //float64

	// 2 types of comments: // and
	/*
		as much rows as needed can be commented
	*/
}
