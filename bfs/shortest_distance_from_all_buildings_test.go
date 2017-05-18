package bfs

import (
	"fmt"
	"testing"
)

func TestShortestDistanceFromAllBuildingsFunc(t *testing.T) {
	t.Log("Test Shortes Distance From All Buildings algorithm")
	grid := [][]int{
		{1, 0, 2, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
	}
	fmt.Println("INPUT:", grid)
	fmt.Println("OUTPUT: ", shortestDistance(grid))
}
