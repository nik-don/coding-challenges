/*
	Bitmap Holes

The Bitmap Holes Challenge is a type of problem where you are given a 2D array
representing a bitmap, and you have to find "holes" in the bitmap.
A "hole" is defined as a group of 0s surrounded by 1s. The group of 0s can be
in any shape but must be bounded by 1s along the horizontal and vertical directions.
*/
package main

import "fmt"

func main() {

	bitmap := []string{
		"10111",
		"10101",
		"11101",
		"11111",
	}

	fmt.Println(countHoles(bitmap)) //answer: 2 areas

	bitmap = []string{
		"11110",
		"11001",
		"11101",
		"10011",
	}

	fmt.Println(countHoles(bitmap)) //answer: 3 areas

}

func countHoles(strarr []string) int {

	// convert to rune array
	var strarrRune [][]rune
	for _, v := range strarr {
		strarrRune = append(strarrRune, []rune(v))
	}
	count := 0
	for i := 0; i < len(strarrRune); i++ {
		for j := 0; j < len(strarrRune[i]); j++ {
			if strarrRune[i][j] == '0' {
				count++
				ReplaceZeros(strarrRune, i, j)
			}
		}
	}

	PrintGrid(strarrRune)
	return count

}

// ReplaceZeros recursively checks for a 0 in 2d array up, down, left, right until there is no zero
// and replaces it, so that it is not counted again.
// The recursion ends when it reaches a cell that either doesn't contain '0' or is
// outside the boundaries of the 2D slice
func ReplaceZeros(strarr [][]rune, i, j int) {
	// Check if indices are out of bounds or if the cell is not '0'
	if i < 0 || i >= len(strarr) || j < 0 || j >= len(strarr[i]) || strarr[i][j] != '0' {
		return
	}

	// Replace the current cell with a chosen character
	strarr[i][j] = 'x'

	// Recursively call for adjacent cells
	ReplaceZeros(strarr, i-1, j) // up
	ReplaceZeros(strarr, i+1, j) // down
	ReplaceZeros(strarr, i, j-1) // left
	ReplaceZeros(strarr, i, j+1) // right
}

func PrintGrid(grid [][]rune) {
	for i := 0; i < len(grid); i++ {
		fmt.Println(string(grid[i]))
	}
}
