package DataStructure

type WordDictionary struct {
	children [26]*WordDictionary
	isWord   bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (wd *WordDictionary) AddWord(word string) {
	node := wd
	for _, char := range word {
		char = char - 'a'
		if node.children[char] == nil {
			node.children[char] = &WordDictionary{}
		}
		node = node.children[char]
	}

	node.isWord = true
}

func (wd *WordDictionary) Search(word string) bool {
	var search func(wd *WordDictionary, idx int) bool
	search = func(wd *WordDictionary, idx int) bool {
		if idx == len(word) {
			return wd.isWord
		}

		char := word[idx]
		if char == '.' {
			for i := range wd.children {
				if wd.children[i] != nil && search(wd.children[i], idx+1) {
					return true
				}
			}
		} else {
			char -= 'a'
			return wd.children[char] != nil && search(wd.children[char], idx+1)
		}
		return false
	}

	return search(wd, 0)
}
