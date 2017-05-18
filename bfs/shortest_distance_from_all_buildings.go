/*
do BFS from each building, and decrement all empty place for every building visit
when grid[i][j] == -b_nums, it means that grid[i][j] are already visited from all b_nums
and use dist to record distances from b_nums
*/

package bfs

import (
	"math"
)

type flow struct {
	i    int
	j    int
	step int
}

func shortestDistance(grid [][]int) float64 {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	matrix := generateMatrix(len(grid), len(grid[0]), 2)

	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				bfs(grid, &matrix, i, j, count)
				count += 1
			}
		}
	}

	res := math.Inf(0) // +Inf
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j][1] == count {
				res = math.Min(res, float64(matrix[i][j][0]))
			}
		}
	}

	if res != math.Inf(0) {
		return res
	}

	return -1
}

func bfs(grid [][]int, matrix *[][][]int, i int, j int, count int) {
	q := []flow{{i, j, 0}}
	valueMatrix := *matrix
	for len(q) > 0 {
		popValueByQ := q[0]
		q = q[1:]
		i, j, step := popValueByQ.i, popValueByQ.j, popValueByQ.step
		tempDistance := [][]int{
			{i - 1, j},
			{i + 1, j},
			{i, j - 1},
			{i, j + 1},
		}
		for _, list := range tempDistance {
			var k, l int
			for i, value := range list {
				if i == 0 {
					k = value
				} else {
					l = value
				}
			}
			if 0 <= k && k < len(grid) && 0 <= l && l < len(grid[0]) && valueMatrix[k][l][1] == count && grid[k][l] == 0 {
				valueMatrix[k][l][0] += step + 1
				valueMatrix[k][l][1] = count + 1
				tempQValue := new(flow)
				tempQValue.i = k
				tempQValue.j = l
				tempQValue.step = step + 1
				q = append(q, *tempQValue)
			}
		}
	}
}

func generateMatrix(x int, y int, z int) [][][]int {
	var matrix [][][]int

	matrix = make([][][]int, x)

	for i := range matrix {
		matrix[i] = make([][]int, y)
		for j := range matrix[i] {
			matrix[i][j] = make([]int, z)
		}
	}

	return matrix
}
