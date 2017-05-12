package utils

func Pop(slice []int) []int {
	slice = slice[:len(slice)-1]
	return slice
}