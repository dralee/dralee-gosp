/*
Trie字符替换器
2025.4.21 by dralee
*/
package utils

import (
	"strings"
)

type TrieNode struct {
	children map[rune]*TrieNode
	word     string // 记录完整的单词
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

type Trie struct {
	root           *TrieNode
	minReplaceSize int // 最短替换长度，小于该长度，则不执行替换。默认为0（全部替换）
}

func NewTrie(minReplaceSize int) *Trie {
	return &Trie{root: NewTrieNode(), minReplaceSize: minReplaceSize}
}

// 添加关键词
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch]
	}
	node.word = word
}

// 查找文本中所有匹配到的词（可以替换用）
func (t *Trie) Replace(text string, replaceMap map[string]string) string {
	var builder strings.Builder
	node := t.root
	runes := []rune(text)

	cache := map[string]string{}
	for i := 0; i < len(runes); {
		cur := node
		j := i
		var (
			longestMatch string
			matchEnd     int
		)

		// 往下走，找到最长匹配
		for j < len(runes) && cur.children[runes[j]] != nil {
			cur = cur.children[runes[j]]
			if cur.word != "" {
				longestMatch = cur.word
				matchEnd = j + 1
			}
			j++
		}

		if longestMatch != "" && (t.minReplaceSize == 0 || j-i >= t.minReplaceSize) {
			cache[longestMatch] = replaceMap[longestMatch]
			builder.WriteString(replaceMap[longestMatch])
			i = matchEnd
		} else {
			builder.WriteRune(runes[i])
			i++
		}
	}

	WriteMapString("./data/logtemp.log", cache)

	return builder.String()
}
