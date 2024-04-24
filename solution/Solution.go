package Solution

import (
	"fmt"
)

var dx [4]int = [4]int{-1, 0, 0, 1} // desribes moves by change of coordinates
var dy [4]int = [4]int{0, -1, 1, 0}
var N int = 0

func arrayContains(arr [][3]int, val [3]int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i][0] == val[0] && arr[i][1] == val[1] && arr[i][2] == val[2] {
			return true
		}
	}
	return false
}

func isValid(arr [3]int) bool {
	return (arr[0] >= 0 && arr[0] <= N && arr[1] >= 0 && arr[1] <= N)
}

func Solve(inputData [][3]int) {
	// find N of NxN board
	for _, dot := range inputData {
		N = max(dot[0], dot[1], N)
	}

	// init slice board
	var board [][]int = make([][]int, N+1)
	for i := 0; i <= N; i++ {
		board[i] = make([]int, N+1)
	}

	// fill board to set up references between cells
	for _, dot := range inputData {
		board[dot[1]][dot[0]] = dot[2]
	}

	// iterative BFS algorithm, var cur_path is slice of connected cells for current cell
	var visited = [][3]int{}
	for i := 0; i < len(inputData); i++ {
		if arrayContains(visited, inputData[i]) {
			continue
		}

		var queue = [][3]int{}
		var cur_path = [][3]int{}
		var node = [3]int{}

		queue = append(queue, inputData[i])
		cur_path = append(cur_path, inputData[i])
		if !arrayContains(visited, inputData[i]) {
			visited = append(visited, inputData[i])
		}

		for len(queue) > 0 {
			queue, node = queue[1:], queue[0]

			for j := 0; j < 4; j++ {
				var next_node = node
				next_node[0] += dx[j]
				next_node[1] += dy[j]

				if isValid(next_node) && board[next_node[1]][next_node[0]] == next_node[2] && !arrayContains(visited, next_node) {
					visited = append(visited, next_node)
					cur_path = append(cur_path, next_node)
					queue = append(queue, next_node)
				}
			}
		}
		fmt.Println(cur_path)
	}
	fmt.Println("all done!")
}
