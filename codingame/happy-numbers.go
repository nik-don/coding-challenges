/*
Happy Numbers

This puzzle was confusing as I didn't understand the problem statementn as it was
written in a way to make it challenging.
The trick is to understand that sum of squares of digits needs to be calculated
each time, and the result of that calculation is used as the input for the next
calculation. The program ends when the result of the calculation is either 1 or
a sum has already been seen before (To stop from going in a loop). The sum could be
the original number.
If the result is 1, then the number is happy, otherwise it is sad.

Input
Line 1: integer n

Output
Line 1: the next numbers in the sequence, space-separated.
Line 2: the number of steps to reach 1, or sad.

Constraints
2 <= n <= 1000

Example
Input
13

Output
10 1
2

Example
Input
4

Output
16 37 58 89 145 42 20 4
8
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 13
	calcHappy(n)

	n = 4
	calcHappy(n)
}

func calcHappy(n int) {
	var arr []int
	seen := make(map[int]bool)

	for {
		n = sumOfSquares(n)
		if n == 1 {
			arr = append(arr, 1)
			break
		}
		if seen[n] {
			break
		}
		arr = append(arr, n)
		seen[n] = true
	}

	for i, v := range arr {
		fmt.Printf("%d", v)
		if i < len(arr)-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Println()
	fmt.Println(len(arr))
}

func sumOfSquares(n int) int {
	var sum int

	for _, n := range strconv.Itoa(n) {
		num, _ := strconv.Atoi(string(n))
		sum += num * num
	}
	return sum
}
