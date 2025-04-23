package main

import "fmt"

const N int = 100000

type Trie struct {
	t     [N][26]int
	cnt   [N]int
	nodes int // 标识当前树中节点数目
}

func Constructor() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(word string) {
	p := 0
	for _, c := range word {
		c -= 'a'
		if t.t[p][c] == 0 {
			t.nodes += 1
			t.t[p][c] = t.nodes
		}
		p = t.t[p][c]
	}
	t.cnt[p]++
}

func (t *Trie) Query(word string) int {
	p := 0
	for _, c := range word {
		c -= 'a'
		if p = t.t[p][c]; p == 0 {
			return 0
		}
	}
	return t.cnt[p]
}

func main() {
	t := Constructor()
	t.Insert("hello")
	t.Insert("hello")
	t.Insert("world")
	t.Insert("helloworld")
	fmt.Println(t.Query("hell"))
	fmt.Println(t.Query("hello"))
	fmt.Println(t.Query("helloworl"))
	fmt.Println(t.Query("helloworld"))

}
