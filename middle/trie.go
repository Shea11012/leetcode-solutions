package middle

// 前缀树

type trienode struct {
	level   int // 层级
	data    string
	wordEnd bool // 判断一个字符串的边界
	t       *trie
}

type trie struct {
	children map[uint8]*trienode
}

func newTrier() trie {
	t := trie{children: make(map[uint8]*trienode)}
	return t
}

func (this *trie) addChild(word string, height int) {
	if height == len(word) {
		return
	}

	value := word[height]
	end := false
	if height == len(word)-1 {
		end = true
	}
	t := newTrier()
	node := &trienode{data: string(value), wordEnd: end, level: height, t: &t}

	v, ok := this.children[value] // 在孩子节点中查找是否包含这个 key
	if !ok {
		this.children[value] = node
		v = node
	} else if ok && v.level < height { // 判断 level 避免键冲突
		v.t.children[value] = node
	} else if ok && v.data == string(value) && !v.wordEnd && len(word)-1 == height {
		v.wordEnd = true
	}

	v.t.addChild(word, height+1)
}

func (this *trie) findChild(word string, height int, starwith bool) bool {
	if height == len(word) {
		return false
	}
	key := word[height]
	v, ok := this.children[key]
	if !ok {
		return false
	}

	if v.data == string(key) && len(word)-1 == height && v.wordEnd == true {
		return true
	}

	if starwith && v.data == string(key) && len(word)-1 == height {
		return true
	}

	return v.t.findChild(word, height+1, starwith)
}

func (this *trie) insert(word string) {
	this.addChild(word, 0)
}

func (this *trie) serach(word string) bool {
	return this.findChild(word, 0, false)
}

func (this *trie) startswith(word string) bool {
	return this.findChild(word, 0, true)
}
