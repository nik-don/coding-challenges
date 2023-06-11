/*
Happy Numbers

This puzzle was confusing as I didn't understand the problem statement as it was
written in a way to make it challenging.
The trick is to understand that sum of squares of digits needs to be calculated
each time, and the result of that calculation is used as the input for the next
calculation. The program ends when the result of the calculation is either 1 or
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
	"math"
	"strconv"
)

func main() {
	n := 13
	calcHappy(n)

	n = 4
	calcHappy(n)
}

func calcHappy(n int) {
	original := n
	var arr []int
	for {
		n = sumOfSquares(n)
		arr = append(arr, n)
		if n == 1 || n == original {
			break
		}
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
	var sum float64

	for _, n := range strconv.Itoa(n) {
		num, _ := strconv.Atoi(string(n))
		sum += math.Pow(float64(num), 2)
	}
	return int(sum)
}
