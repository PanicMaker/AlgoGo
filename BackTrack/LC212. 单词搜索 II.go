package BackTrack

// https://leetcode.cn/problems/word-search-ii/description/

func findWords(board [][]byte, words []string) []string {
	type Trie struct {
		children [26]*Trie
		word     string
	}

	trie := &Trie{}

	for _, word := range words {
		node := trie
		for _, ch := range word {
			if node.children[ch-'a'] == nil {
				node.children[ch-'a'] = &Trie{}
			}
			node = node.children[ch-'a']
		}
		node.word = word
	}

	res := make([]string, 0)

	m, n := len(board), len(board[0])

	var backtrack func(i int, j int, node *Trie)
	backtrack = func(i int, j int, node *Trie) {
		// 处理边界
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}

		ch := board[i][j]
		if ch == '*' || node.children[ch-'a'] == nil {
			return
		}

		node = node.children[ch-'a']
		if node.word != "" {
			res = append(res, node.word)
			// 已经访问过就排除，不可重复选
			node.word = ""
		}

		board[i][j] = '*'
		backtrack(i-1, j, node)
		backtrack(i+1, j, node)
		backtrack(i, j-1, node)
		backtrack(i, j+1, node)
		board[i][j] = ch
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			backtrack(i, j, trie)
		}
	}

	return res
}
