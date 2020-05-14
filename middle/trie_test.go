package middle

import (
	"testing"
)

func Test_trie_insert(t *testing.T) {
	root := newTrier()
	testcases := []struct {
		name string
			word string
	}{
		{name:"one",word:"abc"},
		{name:"three",word:"bpple"},
		{name:"two",word:"apple"},
	}

	for _, testcase := range testcases {
		root.insert(testcase.word)
	}

	got := len(root.children)
	if got != 2 {
		t.Errorf("root 的子孩子数量错误：got %d,期望：%d\n",got,2)
	}
}

func Test_trie_search(t *testing.T) {
	testcase := []struct{
		name string
		word string
		want bool
	}{
		{name:"one",word:"apple",want:true},
		{name:"two",word:"app",want:false},
	}
	root := createTrie(t)
	for _, s := range testcase {
		t.Run(s.name, func(t *testing.T) {
			got := root.serach(s.word)
			if got != s.want {
				t.Errorf("查找字符 %s,期望：%t,实际:%t\n",s.word,s.want,got)
			}
		})
	}
}

func createTrie(t *testing.T) trie {
	t.Helper()
	root := newTrier()
	testcases := []struct {
		name string
		word string
	}{
		{name:"one",word:"abc"},
		{name:"three",word:"bpple"},
		{name:"two",word:"apple"},
	}

	for _, testcase := range testcases {
		root.insert(testcase.word)
	}
	return root
}

func Test_trie_startswith(t *testing.T) {
	testcase := []struct{
		name string
		word string
		want bool
	}{
		{name:"one",word:"app",want:true},
		{name:"two",word:"apc",want:false},
	}
	root := createTrie(t)
	for _, s := range testcase {
		t.Run(s.name, func(t *testing.T) {
			got := root.startswith(s.word)
			if got != s.want {
				t.Errorf("查询开头字符串：%s，期望结果：%t，实际结果：%t\n",s.word,s.want,got)
			}
		})
	}
}

func Test_Trie(t *testing.T) {
	testcase := []struct{
		name string
		word string
		action string
		want bool
	}{
		{name:"one",word:"apple",action:"insert",want:true},
		{name:"two",word:"apple",action:"search",want:true},
		{name:"three",word:"app",action:"search",want:false},
		{name:"four",word:"app",action:"startswith",want:true},
		{name:"five",word:"app",action:"insert",want:true},
		{name:"six",word:"app",action:"search",want:true},
	}

	root := newTrier()
	for _, s := range testcase {
		t.Run(s.name, func(t *testing.T) {
			switch s.action {
			case "insert":
				root.insert(s.word)
			case "search":
				got := root.serach(s.word)
				if got != s.want {
					t.Errorf("查询字符串%s，期望结果：%t，实际结果：%t\n",s.word,s.want,got)
				}
			case "startswith":
				got := root.startswith(s.word)
				if got != s.want {
					t.Errorf("前缀字符串%s，期望结果：%t，实际结果：%t\n",s.word,s.want,got)
				}
			}
		})
	}
}