package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(strongPasswordChecker("aaaaa"))
}

/*
一个强密码应满足以下所有条件：
    由至少6个，至多20个字符组成。
    至少包含一个小写字母，一个大写字母，和一个数字。
    同一字符不能连续出现三次 (比如 "...aaa..." 是不允许的, 但是 "...aa...a..." 是可以的)。
编写函数 strongPasswordChecker(s)，s 代表输入字符串，如果 s 已经符合强密码条件，则返回0；否则返回要将 s
修改为满足强密码条件的字符串所需要进行修改的最小步数。
插入、删除、替换任一字符都算作一次修改。
*/

func strongPasswordChecker(s string) int {
	if s == "" {
		return 6
	}

	ss := []byte(s)
	numCount, lowCount, upCount, changStep, delStep, count := 0, 0, 0, 0, 0, len(s)
	var lastLetter byte
	lastLetterCount := 0
	pq := make(PriorityQueue, 0)

	for _, ch := range ss {
		if isLowercase(ch) {
			lowCount++
		}
		if isUppercase(ch) {
			upCount++
		}
		if isNumber(ch) {
			numCount++
		}

		if ch == lastLetter {
			lastLetterCount++
		} else {
			if lastLetterCount >= 3 {
				heap.Push(&pq, ByteItem{
					Value:       lastLetter,
					LetterCount: lastLetterCount,
					Priority:    lastLetterCount % 3,
				})
			}
			lastLetter = ch
			lastLetterCount = 1
		}
	}
	// 判断最后一个
	if lastLetterCount >= 3 {
		heap.Push(&pq, ByteItem{
			Value:       lastLetter,
			LetterCount: lastLetterCount,
			Priority:    lastLetterCount % 3,
		})
	}

	// 出优先队列
	for len(pq) > 0 {
		// 按照优先级，处理连续的子串，规则是优先处理%3==0的子串，其次是%3==1的子串，最后是%3==2的子串
		item := heap.Pop(&pq).(*ByteItem)
		letterCount := item.LetterCount
		fmt.Println("letterCount", letterCount)
		if count < 6 {
			// 如果小于6就增加一个不是item.Letter的进去，所以changStep++,count++,因为加了一个去了,把原来连续的三个隔开了，所以letterCount-=3
			// 比如 aaaaa  变成 aabaaa，把以lettercount就是2,发现还是有连续三个重复的，这里就是技巧了，因为必须大于6个，且有数字，大写，小写，
			// 所以这里变了，即使还有三个重复的，也不能满足其他要求，所以下面会判断
			changStep += 1
			count += 1
			letterCount -= 3
		} else if count > 20 {
			// 就删除一下字母
			letterCount -= 1
			count -= 1
			delStep += 1
		} else {
			changStep += letterCount / 3
			letterCount %= 3
		}
		//
		if letterCount >= 3 {
			heap.Push(&pq, ByteItem{
				Value:       item.Value,
				LetterCount: letterCount,
				Priority:    letterCount % 3,
			})
		}
	}
	noEnoughStep := 0
	if numCount <= 0 {
		noEnoughStep++
	}
	if upCount <= 0 {
		noEnoughStep++
	}
	if lowCount <= 0 {
		noEnoughStep++
	}
	// 判断长度不够或者超过限制所需要增删的步数
	if count < 6 {
		changStep += (6 - count)
	}
	if count > 20 {
		delStep += (count - 20)
	}
	return Max(noEnoughStep, changStep) + delStep
}

func detection(s string) bool {
	if len(s) < 6 || len(s) > 20 {
		return false
	}
	hasDigital, hasUppercase, hasLowercase := false, false, false
	curmax := 1
	var pre byte
	for _, b := range []byte(s) {
		if b >= 'a' && b <= 'z' {
			hasLowercase = true
		}
		if b >= 'A' && b <= 'Z' {
			hasUppercase = true
		}
		if b >= '0' && b <= '9' {
			hasDigital = true
		}

		if b == pre {
			curmax++
			if curmax >= 3 {
				return false
			}
		} else {
			curmax = 1
			pre = b
		}
	}
	return hasDigital && hasUppercase && hasLowercase
}

func isNumber(i byte) bool {
	if i >= '0' && i <= '9' {
		return true
	}
	return false
}

func isUppercase(i byte) bool {
	if i >= 'A' && i <= 'Z' {
		return true
	}
	return false
}
func isLowercase(i byte) bool {
	if i >= 'a' && i <= 'z' {
		return true
	}
	return false
}

type ByteItem struct {
	Value       byte // The Key of the item; arbitrary.
	LetterCount int
	Priority    int // The Priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*ByteItem

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(ByteItem)
	item.index = n
	*pq = append(*pq, &item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// 这种方式就是按升序排序，符合题目，如果是按降序，则要反着来
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func Max(nums ...int) int {
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
