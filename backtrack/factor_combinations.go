/*
Numbers can be regarded as product of its factors. For example,

8 = 2 x 2 x 2;
  = 2 x 4.
Write a function that takes an integer n and return all possible combinations of its factors.

Note:
You may assume that n is always positive.
Factors should be greater than 1 and less than n.
Examples:
input: 1
output:
[]
input: 37
output:
[]
input: 12
output:
[
  [2, 6],
  [2, 2, 3],
  [3, 4]
]
input: 32
output:
[
  [2, 16],
  [2, 2, 8],
  [2, 2, 2, 4],
  [2, 2, 2, 2, 2],
  [2, 4, 4],
  [4, 8]
]
*/

package backtrack

type todo struct {
	n     int
	i     int
	combi []int
}

func iterativeGetFactors(n int) [][]int {
	todos, combis := []todo{{n, 2, []int{}}}, [][]int{}
	for len(todos) > 0 {
		tempTodo := todos[len(todos)-1]
		todos = todos[:len(todos)-1]
		n, i, combi := tempTodo.n, tempTodo.i, tempTodo.combi
		for i*i <= n {
			if n%i == 0 {
				combis = append(combis, append(combi, i, n/i))
				todos = append(todos, todo{n / i, i, append(combi, i)})
			}
			i++
		}
	}
	return combis
}
