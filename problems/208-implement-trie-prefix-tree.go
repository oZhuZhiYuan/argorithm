package problems

/*
word 和 prefix 仅由小写英文字母组成
*/

type Trie struct {
	Children    [26]*Trie
	IsEndOfWord bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (trie *Trie) Insert(word string) {
	if word == "" {
		return
	}
	// for _, c := range []byte(word)
	for _, c := range word {
		if trie.Children[c-97] == nil {
			trie.Children[c-97] = &Trie{}
		}
		trie = trie.Children[c-97]
	}
	trie.IsEndOfWord = true
}

/** Returns if the word is in the trie. */
func (trie *Trie) Search(word string) bool {
	if word == "" {
		return false
	}
	for _, c := range word {
		if trie.Children[c-97] == nil {
			return false
		}
		trie = trie.Children[c-97]
	}
	return trie.IsEndOfWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (trie *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return false
	}
	for _, c := range prefix {
		if trie.Children[c-97] == nil {
			return false
		}
		trie = trie.Children[c-97]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
