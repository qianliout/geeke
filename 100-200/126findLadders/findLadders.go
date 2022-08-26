package main

import (
	"math"
)

func main() {

}

/*
按字典 wordList 完成从单词 beginWord 到单词 endWord 转化，一个表示此过程的 转换序列 是形式上像 beginWord -> s1 -> s2 -> ... -> sk 这样的单词序列，并满足：
每对相邻的单词之间仅有单个字母不同。
转换过程中的每个单词 si（1 <= i <= k）必须是字典 wordList 中的单词。注意，beginWord 不必是字典 wordList 中的单词。
sk == endWord
给你两个单词 beginWord 和 endWord ，以及一个字典 wordList 。请你找出并返回所有从 beginWord 到 endWord 的 最短转换序列 ，
如果不存在这样的转换序列，返回一个空列表。每个序列都应该以单词列表 [beginWord, s1, s2, ..., sk] 的形式返回。
*/
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordSet := make(map[string]bool)
	wordSet[beginWord] = true

	for _, word := range wordList {
		wordSet[word] = true
	}
	res := make([][]string, 0)
	path := make([]string, 0)
	path = append(path, beginWord)
	queue := make([]string, 0)
	queue = append(queue, beginWord)
	visited := make(map[string]bool)
	visited[beginWord] = true
	var mins int = math.MaxInt64

	Dfs(beginWord, endWord, wordSet, &visited, queue, path, &res, &mins)
	// 上面得出了全部,接下来找最少的
	newRes := make([][]string, 0)
	minR := math.MaxInt64
	for _, r := range res {
		if len(r) < minR {
			minR = len(r)
		}
	}
	for _, r := range res {
		if len(r) == minR {
			newRes = append(newRes, r)
		}
	}

	return newRes
}

func Dfs(begin, end string, wordSet map[string]bool, visited *map[string]bool, queue, path []string, res *[][]string, minR *int) {
	if begin == end {
		if len(path) < *minR {
			*minR = len(path)
		}
		*res = append(*res, append([]string{}, path...))
		return
	}
	// 改词
	temQueue := make([]string, 0)
	for _, word := range queue {
		for i := 0; i < len(word); i++ {
			wordS := []byte(word)
			for j := 97; j < 97+26; j++ {
				if int(wordS[i]) != j {
					wordS[i] = byte(j)
					if wordSet[string(wordS)] && !(*visited)[string(wordS)] {
						temQueue = append(temQueue, string(wordS))
					}
				}
			}
		}
	}

	for _, word := range temQueue {
		if len(path) < *minR {
			tem := make([]string, 0)
			tem = append(tem, word)
			path := append(path, word)
			(*visited)[word] = true
			Dfs(word, end, wordSet, visited, tem, path, res, minR)
			path = path[:len(path)-1]
			(*visited)[word] = false
		}
	}
}
