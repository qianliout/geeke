package main

import (
	"fmt"
)

func main() {
	wordList := []string{"hot", "cog", "dot", "dog", "hit", "lot", "log"}
	res := ladderLength("hit", "cog", wordList)
	//wordList := []string{"peale", "wilts", "place", "fetch", "purer", "pooch", "peace", "poach", "berra", "teach", "rheum", "peach"}
	//res := ladderLength("teach", "place", wordList)
	fmt.Println("res is ", res)
}

/*
   给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
       每次转换只能改变一个字母。
       转换过程中的中间单词必须是字典中的单词。
   说明:
       如果不存在这样的转换序列，返回 0。
       所有单词具有相同的长度。
       所有单词只由小写字母组成。
       字典中不存在重复的单词。
       你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
   示例 1:
   输入:
   beginWord = "hit",
   endWord = "cog",
   wordList = ["hot","dot","dog","lot","log","cog"]
   输出: 5
   解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
        返回它的长度 5。
   示例 2:
   输入:
   beginWord = "hit"
   endWord = "cog"
   wordList = ["hot","dot","dog","lot","log"]
   输出: 0
   解释: endWord "cog" 不在字典中，所以无法进行转换。
*/

// 方法一,单边bfs
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return 0
	}
	count := 1
	visited := make(map[string]bool)
	visited[beginWord] = true

	// bfs start there
	queue := make([]string, 0)
	temQueue := make([]string, 0)
	queue = append(queue, beginWord)
	for len(queue) > 0 {
		for _, first := range queue {

			for i := 0; i < len(first); i++ {
				firstS := []byte(first)
				for j := 97; j < 97+26; j++ {
					if int(firstS[i]) != j {
						firstS[i] = byte(j)
						if string(firstS) == endWord {
							return count + 1
						}
						if wordSet[string(firstS)] && !visited[string(firstS)] {
							//fmt.Println("fisrtS is ", first, string(firstS))
							visited[string(firstS)] = true
							temQueue = append(temQueue, string(firstS))
						}
					}
				}
			}
		}
		//fmt.Println("tem queue", temQueue)
		queue = temQueue
		temQueue = make([]string, 0)
		count++
	}
	return 0
}

// 方法二,双边dfs
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}
	wordSet := make(map[string]bool)
	wordSet[beginWord] = true
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return 0
	}
	count := 1
	visited := make(map[string]bool)
	visited[beginWord] = true

	// bfs start there
	beginSet := make(map[string]bool)
	beginSet[beginWord] = true
	endSet := make(map[string]bool)
	endSet[endWord] = true

	for len(beginSet) > 0 && len(endSet) > 0 {
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}
		temSet := make(map[string]bool)
		for first, _ := range beginSet {
			for i := 0; i < len(first); i++ {
				firstS := []byte(first)
				for j := 97; j < 97+26; j++ {
					if int(firstS[i]) != j {
						firstS[i] = byte(j)
						if endSet[string(firstS)] {
							return count + 1
						}
						if wordSet[string(firstS)] && !visited[string(firstS)] {
							//fmt.Println("fisrtS is ", first, string(firstS))
							visited[string(firstS)] = true
							temSet[string(firstS)] = true
						}
					}
				}
			}
		}
		//fmt.Println("tem queue", temSet)
		beginSet = temSet
		temSet = make(map[string]bool)
		count++
	}
	return 0
}
