package main

import Solution "github.com/Razzle131/iterativeGoBFS/solution"

func main() {
	var inputData [][3]int

	inputData = append(inputData, [3]int{0, 0, -1})
	inputData = append(inputData, [3]int{1, 0, -1})
	inputData = append(inputData, [3]int{0, 1, -1})
	inputData = append(inputData, [3]int{2, 1, -1})
	inputData = append(inputData, [3]int{1, 1, 1})

	Solution.Solve(inputData)
}
