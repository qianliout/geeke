package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	random := Constructor()
	fmt.Println(random.Insert(0))
	random.Insert(1)
	random.Remove(0)
	random.Insert(2)
	random.Remove(1)
	fmt.Println(random.GetRandom())
}

/*
设计一个支持在平均 时间复杂度 O(1) 下， 执行以下操作的数据结构。
注意: 允许出现重复元素。
    insert(val)：向集合中插入元素 val。
    remove(val)：当 val 存在时，从集合中移除一个 val。
    getRandom：从现有集合中随机获取一个元素。每个元素被返回的概率应该与其在集合中的数量呈线性相关。
示例:
// 初始化一个空的集合。
RandomizedCollection collection = new RandomizedCollection();
// 向集合中插入 1 。返回 true 表示集合不包含 1 。
collection.insert(1);
// 向集合中插入另一个 1 。返回 false 表示集合包含 1 。集合现在包含 [1,1] 。
collection.insert(1);
// 向集合中插入 2 ，返回 true 。集合现在包含 [1,1,2] 。
collection.insert(2);
// getRandom 应当有 2/3 的概率返回 1 ，1/3 的概率返回 2 。
collection.getRandom();
// 从集合中删除 1 ，返回 true 。集合现在包含 [1,2] 。
collection.remove(1);
// getRandom 应有相同概率返回 1 和 2 。
collection.getRandom();
*/

type RandomizedCollection struct {
	list   []int
	index  map[int]map[int]bool
	rander *rand.Rand
	length int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	list := make([]int, 0)
	index := make(map[int]map[int]bool)
	rander := rand.New(rand.NewSource(time.Now().Unix()))
	return RandomizedCollection{
		list:   list,
		index:  index,
		rander: rander,
		length: 0,
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	indexSet, ok := this.index[val]
	res := false
	if !ok || indexSet == nil {
		this.index[val] = make(map[int]bool)
		res = true
	}
	
	this.list = append(this.list[:this.length], val)
	this.index[val][this.length] = true
	this.length++
	return res
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	indexset, ok := this.index[val]
	if !ok || indexset == nil {
		return false
	}
	delIndex := make([]int, 0)
	for k, v := range indexset {
		if v {
			delIndex = append(delIndex, k)
		}
	}
	if len(delIndex) == 0 {
		return false
	}
	
	sort.Ints(delIndex)
	del := delIndex[len(delIndex)-1]
	this.index[val][del] = false
	last := this.list[this.length-1]
	// 这里一定要是先置为true再置为false，因为可能删除原素就是最后一下
	this.index[last][del] = true
	this.index[last][this.length-1] = false
	this.list[del] = this.list[this.length-1]
	this.length--
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	in := this.rander.Intn(this.length)
	return this.list[in]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
