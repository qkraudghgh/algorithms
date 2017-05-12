/*
Given a matrix of words and a list of words to search, return a list of words that exists in the board
This is Word Search II on LeetCode

board = [
         ['o','a','a','n'],
         ['e','t','a','e'],
         ['i','h','k','r'],
         ['i','f','l','v']
         ]

words = ["oath","pea","eat","rain"]
*/

package backtrack

type rec map[byte]rec

func findWords(board [][]byte, words []string) []string {
	trie := make(map[byte]rec)
	for _, word := range words {
		curr_trie := trie
		for _, char := range word {
			if _, ok := curr_trie[byte(char)]; !ok {
				curr_trie[byte(char)] = make(map[byte]rec)
			}
			curr_trie = curr_trie[byte(char)]
		}
		curr_trie['#'] = nil
	}

	result := []string{}
	used := [][]bool{}
	for i := 0; i < len(board); i++ {
		tempUsed := []bool{}
		for j := 0; j < len(board[0]); j++ {
			tempUsed = append(tempUsed, false)
		}
		used = append(used, tempUsed)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			searchBacktrack(board, i, j, trie, "", &used, &result)
		}
	}
	return result
}

func searchBacktrack(board [][]byte, i int, j int, trie map[byte]rec, pre string, used *[][]bool, result *[]string) {
	tempUsed := *used
	if _, ok := trie['#']; ok && checker(*result, pre) {
		*result = append(*result, pre)
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}

	if !tempUsed[i][j] {
		if _, ok := trie[board[i][j]]; ok {
			tempUsed[i][j] = true
			searchBacktrack(board, i+1, j, trie[board[i][j]], pre+string(board[i][j]), used, result)
			searchBacktrack(board, i, j+1, trie[board[i][j]], pre+string(board[i][j]), used, result)
			searchBacktrack(board, i-1, j, trie[board[i][j]], pre+string(board[i][j]), used, result)
			searchBacktrack(board, i, j-1, trie[board[i][j]], pre+string(board[i][j]), used, result)
			tempUsed[i][j] = false
		}
	}
}

func checker(list []string, checkee string) bool {
	for _, str := range list {
		if str == checkee {
			return false
		}
	}
	return true
}
