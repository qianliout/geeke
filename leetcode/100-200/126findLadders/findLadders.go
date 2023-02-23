package main

import (
	"fmt"
)

func main() {
	beginWord := "cet"
	endWord := "ism"
	wordList := []string{"kid", "tag", "pup", "ail", "tun", "woo", "erg", "luz", "brr", "gay", "sip", "kay", "per", "val", "mes",
		"ohs", "now", "boa", "cet", "pal", "bar", "die", "war", "hay", "eco", "pub", "lob", "rue", "fry", "lit", "rex", "jan",
		"cot", "bid", "ali", "pay", "col", "gum", "ger", "row", "won", "dan", "rum", "fad", "tut", "sag", "yip", "sui", "ark",
		"has", "zip", "fez", "own", "ump", "dis", "ads", "max", "jaw", "out", "btu", "ana", "gap", "cry", "led", "abe",
		"box", "ore", "pig", "fie", "toy", "fat", "cal", "lie", "noh", "sew", "ono", "tam", "flu", "mgm", "ply", "awe",
		"pry", "tit", "tie", "yet", "too", "tax", "jim", "san", "pan", "map", "ski", "ova", "wed", "non", "wac", "nut",
		"why", "bye", "lye", "oct", "old", "fin", "feb", "chi", "sap", "owl", "log", "tod", "dot", "bow", "fob", "for",
		"joe", "ivy", "fan", "age", "fax", "hip", "jib", "mel", "hus", "sob", "ifs", "tab", "ara", "dab", "jag", "jar",
		"arm", "lot", "tom", "sax", "tex", "yum", "pei", "wen", "wry", "ire", "irk", "far", "mew", "wit", "doe", "gas",
		"rte", "ian", "pot", "ask", "wag", "hag", "amy", "nag", "ron", "soy", "gin", "don", "tug", "fay", "vic", "boo",
		"nam", "ave", "buy", "sop", "but", "orb", "fen", "paw", "his", "sub", "bob", "yea", "oft", "inn", "rod", "yam",
		"pew", "web", "hod", "hun", "gyp", "wei", "wis", "rob", "gad", "pie", "mon", "dog", "bib", "rub", "ere", "dig",
		"era", "cat", "fox", "bee", "mod", "day", "apr", "vie", "nev", "jam", "pam", "new", "aye", "ani", "and", "ibm",
		"yap", "can", "pyx", "tar", "kin", "fog", "hum", "pip", "cup", "dye", "lyx", "jog", "nun", "par", "wan", "fey",
		"bus", "oak", "bad", "ats", "set", "qom", "vat", "eat", "pus", "rev", "axe", "ion", "six", "ila", "lao", "mom",
		"mas", "pro", "few", "opt", "poe", "art", "ash", "oar", "cap", "lop", "may", "shy", "rid", "bat", "sum", "rim",
		"fee", "bmw", "sky", "maj", "hue", "thy", "ava", "rap", "den", "fla", "auk", "cox", "ibo", "hey", "saw", "vim",
		"sec", "ltd", "you", "its", "tat", "dew", "eva", "tog", "ram", "let", "see", "zit", "maw", "nix", "ate", "gig",
		"rep", "owe", "ind", "hog", "eve", "sam", "zoo", "any", "dow", "cod", "bed", "vet", "ham", "sis", "hex", "via",
		"fir", "nod", "mao", "aug", "mum", "hoe", "bah", "hal", "keg", "hew", "zed", "tow", "gog", "ass", "dem", "who",
		"bet", "gos", "son", "ear", "spy", "kit", "boy", "due", "sen", "oaf", "mix", "hep", "fur", "ada", "bin", "nil",
		"mia", "ewe", "hit", "fix", "sad", "rib", "eye", "hop", "haw", "wax", "mid", "tad", "ken", "wad", "rye", "pap",
		"bog", "gut", "ito", "woe", "our", "ado", "sin", "mad", "ray", "hon", "roy", "dip", "hen", "iva", "lug", "asp",
		"hui", "yak", "bay", "poi", "yep", "bun", "try", "lad", "elm", "nat", "wyo", "gym", "dug", "toe", "dee", "wig",
		"sly", "rip", "geo", "cog", "pas", "zen", "odd", "nan", "lay", "pod", "fit", "hem", "joy", "bum", "rio", "yon",
		"dec", "leg", "put", "sue", "dim", "pet", "yaw", "nub", "bit", "bur", "sid", "sun", "oil", "red", "doc", "moe",
		"caw", "eel", "dix", "cub", "end", "gem", "off", "yew", "hug", "pop", "tub", "sgt", "lid", "pun", "ton", "sol",
		"din", "yup", "jab", "pea", "bug", "gag", "mil", "jig", "hub", "low", "did", "tin", "get", "gte", "sox", "lei",
		"mig", "fig", "lon", "use", "ban", "flo", "nov", "jut", "bag", "mir", "sty", "lap", "two", "ins", "con", "ant",
		"net", "tux", "ode", "stu", "mug", "cad", "nap", "gun", "fop", "tot", "sow", "sal", "sic", "ted", "wot", "del",
		"imp", "cob", "way", "ann", "tan", "mci", "job", "wet", "ism", "err", "him", "all", "pad", "hah", "hie", "aim"}

	// beginWord := "hit"
	// endWord := "cog"
	// wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	res := findLadders(beginWord, endWord, wordList)
	fmt.Println(res)
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordDict := make(map[string]bool)
	for i := range wordList {
		wordDict[wordList[i]] = true
	}
	wordDict[beginWord] = true
	res := make([][]string, 0)
	if !wordDict[endWord] {
		return res
	}
	graph := make(map[string][]string)
	distance := make(map[string]int)

	bfs(endWord, wordDict, graph, distance)

	dfs(beginWord, endWord, &res, []string{}, graph, distance)

	return res
}

func dfs(beginWord, endWord string, res *[][]string, path []string, graph map[string][]string, distance map[string]int) {
	if beginWord == endWord {
		*res = append(*res, append(append([]string{}, path...), endWord))
		return
	}
	for _, v := range graph[beginWord] {
		if distance[beginWord] == distance[v]+1 {
			path = append(path, beginWord)
			dfs(v, endWord, res, path, graph, distance)
			// 回溯（执行完上述的所有时，将其回溯回去）
			path = path[:len(path)-1]
		}
	}
}

func bfs(endWord string, wordDict map[string]bool, graph map[string][]string, distance map[string]int) {
	distance[endWord] = 0
	queue := make([]string, 0)
	queue = append(queue, endWord)

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		expands := expand(first, wordDict)
		// 构造图
		for _, v := range expands {
			graph[v] = append(graph[v], first)

			if _, ok := distance[v]; !ok {
				distance[v] = distance[first] + 1
				queue = append(queue, v)
			}
		}
	}
}

func expand(word string, dict map[string]bool) []string {
	wordByte := []byte(word)
	res := make([]string, 0)
	for i := range wordByte {
		b := wordByte[i]
		for j := 'a'; j <= 'z'; j++ {
			wordByte[i] = byte(j)
			if dict[string(wordByte)] && byte(j) != b {
				res = append(res, string(wordByte))
			}
			wordByte[i] = b
		}
	}
	return res
}

// /*
// 按字典 wordList 完成从单词 beginWord 到单词 endWord 转化，一个表示此过程的 转换序列 是形式上像 beginWord -> s1 -> s2 -> ... -> sk 这样的单词序列，并满足：
// 每对相邻的单词之间仅有单个字母不同。
// 转换过程中的每个单词 si（1 <= i <= k）必须是字典 wordList 中的单词。注意，beginWord 不必是字典 wordList 中的单词。
// sk == endWord
// 给你两个单词 beginWord 和 endWord ，以及一个字典 wordList 。请你找出并返回所有从 beginWord 到 endWord 的 最短转换序列 ，
// 如果不存在这样的转换序列，返回一个空列表。每个序列都应该以单词列表 [beginWord, s1, s2, ..., sk] 的形式返回。
// */
// // bfs+dfs(如果是双向bfs，效果会更好)
// func findLadders(beginWord string, endWord string, wordList []string) [][]string {
// 	// 字典表（将wordList中的单词放入hash表中，方便查找）
// 	dict := make(map[string]bool, 0)
// 	for _, v := range wordList {
// 		dict[v] = true
// 	}
// 	// 如果endWord不在hash表中，表示不存在转换列表，返回空集合
// 	if !dict[endWord] {
// 		return [][]string{}
// 	}
// 	// 将第一个单词放入hash表中，方便实现邻接表（构建图）
// 	dict[beginWord] = true
// 	// 构建邻接表
// 	graph := make(map[string][]string, 0)
// 	// 执行bfs搜索，找到每个点到endWord的距离
// 	distance := bfs(endWord, dict, graph)
// 	res := make([][]string, 0) // 保存结果
// 	// 执行dfs操作
// 	dfs(beginWord, endWord, &res, []string{}, distance, graph)
// 	return res
// }
//
// // 回溯实现方式一：（个人偏好这个，更符合模板）
// func dfs(beginWord string, endWord string, res *[][]string, path []string, distance map[string]int, graph map[string][]string) {
// 	// 出递归条件
// 	if beginWord == endWord {
// 		path = append(path, beginWord) // 加入endWord节点
// 		tmp := make([]string, len(path))
// 		copy(tmp, path)
// 		*res = append(*res, tmp)
// 		path = path[:len(path)-1] // 移除endWord节点
// 		return
// 	}
// 	// 否则遍历图
// 	for _, v := range graph[beginWord] {
// 		// 遍历图时，朝着距离与终点越来越近的方向进行（当前节点的距离肯定要比下一个距离大1）
// 		if distance[beginWord] == distance[v]+1 {
// 			path = append(path, beginWord)
// 			dfs(v, endWord, res, path, distance, graph)
// 			// 回溯（执行完上述的所有时，将其回溯回去）
// 			path = path[:len(path)-1]
// 		}
// 	}
// }
//
// // // 回溯实现方式二：
// // func dfs(beginWord string, endWord string, res *[][]string, path []string, distance map[string]int, graph map[string][]string) {
// // 	path = append(path, beginWord)
// // 	// 出递归条件
// // 	if beginWord == endWord {
// // 		tmp := make([]string, len(path))
// // 		copy(tmp, path)
// // 		(*res) = append((*res), tmp)
// // 		return
// // 	}
// // 	// 否则遍历图
// // 	for _, v := range graph[beginWord] {
// // 		// 遍历图时，朝着距离与终点越来越近的方向进行（当前节点的距离肯定要比下一个距离大1）
// // 		if distance[beginWord] == distance[v]+1 {
// // 			dfs(v, endWord, res, path, distance, graph)
// // 		}
// // 	}
// // 	// 回溯（执行完上述的所有时，将其回溯回去）
// // 	path = path[:len(path)-1]
// // }
//
// // 从终点出发，进行bfs，计算每一个点到达终点的距离
// func bfs(endWord string, dict map[string]bool, graph map[string][]string) map[string]int {
// 	distance := make(map[string]int, 0) // 每个点到终点的距离
// 	queue := make([]string, 0)
// 	queue = append(queue, endWord)
//
// 	distance[endWord] = 0 // 初始值
//
// 	for len(queue) != 0 {
// 		curSize := len(queue)
// 		for i := 0; i < curSize; i++ {
// 			word := queue[0]
// 			queue = queue[1:]
// 			// 找到和word有一位不同的单词列表
// 			expansion := expand(word, dict)
// 			for _, v := range expansion {
// 				// 构造邻接表
// 				// 我们是从beginWord到endWord构造邻接表，而bfs是从endWord开始，所以构造时，反过来构造
// 				// 即graph[v]=append(graph[v],word)而不是graph[word]=append(graph[word],v)
// 				// 两种方式都是正确的
// 				graph[v] = append(graph[v], word)
// 				// graph[word] = append(graph[word], v)
// 				// 表示没访问过
// 				if _, ok := distance[v]; !ok {
// 					distance[v] = distance[word] + 1 // 距离加一
// 					queue = append(queue, v)         // 入队列
// 				}
// 			}
// 		}
// 	}
// 	return distance
// }
//
// // 获得邻接点
// func expand(word string, dict map[string]bool) []string {
// 	expansion := make([]string, 0) // 保存word的邻接点
// 	// 从word的每一位开始
// 	chs := []rune(word)
// 	for i := 0; i < len(word); i++ {
// 		tmp := chs[i] // 保存当前位，方便后序进行复位
// 		for c := 'a'; c <= 'z'; c++ {
// 			// 如果一样则直接跳过，之所以用tmp，是因为chs[i]在变
// 			if tmp == c {
// 				continue
// 			}
// 			chs[i] = c
// 			newStr := string(chs)
// 			// 新单词在dict中不存在，则跳过
// 			if dict[newStr] {
// 				expansion = append(expansion, newStr)
// 			}
// 		}
// 		chs[i] = tmp // 单词位复位
// 	}
// 	return expansion
// }
