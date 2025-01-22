/*
Modify an existing function using a "wrapper" function.
*/

package main

import (
	"fmt"
	"strconv"
)

const number = 7

func main() {
	fmt.Println("number:", strconv.Itoa(number))
	n1 := multiply(number)
	fmt.Println("multiply:", strconv.Itoa(n1))

	wf1 := wrap1(multiply)
	n2 := wf1(number)
	fmt.Println("wf1-multiply:", strconv.Itoa(n2))

	wf2 := wrap2(multiply)
	txt := wf2(number)
	fmt.Println(txt)

	wf3 := wrap3(multiply)
	n3 := wf3(number,number)
	fmt.Println("wf3-multiply:", strconv.Itoa(n3))
}

// Original function
func multiply(n int) int {
	return n * 2
}

// Same signature, modify argument
func wrap1(f func(i int) int) func(i int) int {
    return func(i int) (ret int) {
	i = i + 3
	fmt.Println("wf1-number:", strconv.Itoa(i))
        ret = f(i)
        return
    }
}

// Different return type
func wrap2(f func(i int) int) func(i int) string {
    return func(i int) (ret string) {
	i = i * 3
	fmt.Println("wf2-number:", strconv.Itoa(i))
	result := f(i)
	ret = "wf2-multiply: " + strconv.Itoa(result)
        return 
    }
}

// Add argument
func wrap3(f func(i int) int) func(i int, j int) int {
    return func(i int, j int) (ret int) {
	i = i * j
	fmt.Println("wf3-number:", strconv.Itoa(i))
	ret = f(i)
        return 
    }
}
