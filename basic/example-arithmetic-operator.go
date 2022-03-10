package main

import "fmt"

func basic(){
	var a int = 10
	var b int = 20
	fmt.Printf("Addition of numbers: %d\n", a+b)
	fmt.Printf("Subraction of numbers: %d\n", a-b)
	// Only float64 can multiply by float64, int * float or float32 * float64 is not allowed
	fmt.Printf("Multiplication of numbers: %f\n", float64(a)*float64(b))
}

func multiplication_tables(start_index int, end_index int, input_number int){
	var a int = input_number
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d X %d = %d\n",a,i, a*i)
	}
}

func main(){
	basic()
	multiplication_tables(1,10,10)
}