package main

import (
	"fmt"
)

const(
	constA = iota
	_
	constB = iota + 1
)

func main() {
	fmt.Println("Data Structures and Algorithm in Golang")
	fmt.Println()

	// Basic Datatypes
	// https://gist.github.com/thatisuday/c17e05de591c2e2021ab402e4c2d4bdc

	// Variable declaration
	fmt.Println("Variable decalration")
	// Multiple variable decalaration
	// var var1, var2, var3 int
	// var var1, var2, var3 = 1, 2.2, false (can also be done in shorthand)
	a, b, c := 1, 2.2, true
	fmt.Printf("Multi variable declaration: %d, %f, %t\n", a, b, c)

	// Type Conversion
	d := a + int(b)
	e := float32(a) + float32(b) // We still convert b because it's original type is float64 (in x64 arcg)
	fmt.Printf("Type conversion: d = %d, e = %f", d, e)
	s := "hello"
	sByte := []byte(s)
	fmt.Printf("Type conversion string: s = %s, sByte = %v\n", s, sByte)
	
	// Type Declaration & Alias
	// Type declaration: type float float64
	// Typle alias: type aliasName = oldType
	// Since alias still reference it's old type, it is comparable with its old type

	// Constants
	// const const_name [data_type] = fixed_value
	// It's value must be known at complite time, you can't use a function for a constant value
	fmt.Printf("Constants: %d, %d \n", constA, constB)

	// Numeric Expressions
	integerDivision := 11/2
	decimalDivision := float32(11)/2
	fmt.Printf("Numeric expressions depend on the variable type: integerDivision = %d, decimalDivision = %f\n", integerDivision, decimalDivision)


	// Arrays
	fmt.Println()
	fmt.Println("Arrays")
	var arr [3]int
	fmt.Printf("Array declaration: %v\n", arr)
	arr = [3]int{1, 2, 3}
	fmt.Printf("Array assignment : %v\n", arr)
	arr[1] = 99
	fmt.Printf("Modified array: %v\n", arr)
	fmt.Printf("Length of array: %d\n", len(arr))


	pallete := [3]string{"red", "yellow", "green"}
	pallete2 := pallete
	pallete2[1] = "pink"
	fmt.Printf("Arrays are copied by value: pallete %v, pallete2 %v\n", pallete, pallete2)
	
}
