package day55

const trie_widths = 256

type trie_node struct {
	count int
	pass int
	next [trie_widths]*trie_node
}

type trie struct {
	root *trie_node
}

func NewTrie() *trie {
	return &trie{
		root: &trie_node{
			count: 0,
			pass: 0,
			next: [trie_widths]*trie_node{},
		},
	}
}

// 字典树插入
func (t *trie) Insert(str string) {
	rune_str := []rune(str)

	cur := t.root
	for _, v := range rune_str {
		if cur.next[v] == nil {
			newNode := &trie_node{
				count: 0,
				pass: 0,
				next: [trie_widths]*trie_node{},	
			}

			cur.next[v] = newNode
		}
		cur = cur.next[v]
		cur.pass++
	}
	cur.count++
}

func (t *trie) Search(str string) bool {
	rune_str := []rune(str)

	cur := t.root
	for _, v := range rune_str {
		cur.pass++
		if cur.next[v] == nil {
			return false
		}
		cur = cur.next[v]
	}

	if cur.count == 0 {
		return false
	}

	return true
}

// 查找字符串出现次数
func (t *trie) Lookup(str string) int {
	rune_str := []rune(str)

	cur := t.root
	for _, v := range rune_str {
		cur.pass++
		if cur.next[v] == nil {
			return 0
		}
		cur = cur.next[v]
	}

	return cur.count
}

// 删除字符串
func (t *trie) Delete(str string) {
	if !t.Search(str) {
		return 
	}
	rune_str := []rune(str)

	cur := t.root
	for _, v := range rune_str {
		cur.pass--
		cur = cur.next[v]
		if cur.count == 0 && cur.pass == 0 {
			cur = nil
		}
	}
	cur.count--

	if cur.count == 0 && cur.pass == 0 {
		cur = nil
	}
}