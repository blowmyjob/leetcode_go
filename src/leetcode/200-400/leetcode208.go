package leetcode

type Trie struct {
	isEnd    bool
	children [26]*Trie
}

/** Initialize your data structure here. */
func Constructor7() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		if node.children[word[i]-'a'] == nil {
			node.children[word[i]-'a'] = &Trie{}
		}
		node = node.children[word[i]-'a']
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for i := 0; i < len(word); i++ {
		if node.children[word[i]-'a'] == nil {
			return false
		}
		node = node.children[word[i]-'a']
	}
	return node.isEnd == true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for i := 0; i < len(prefix); i++ {
		if node.children[prefix[i]-'a'] == nil {
			return false
		}
		node = node.children[prefix[i]-'a']
	}
	return true
}
