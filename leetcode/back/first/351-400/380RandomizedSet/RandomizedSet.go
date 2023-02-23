package main

import (
	"fmt"
	"math/rand"
)

func main() {
	r := Constructor()
	fmt.Println(r.Insert(0))
	fmt.Println(r.data)
	fmt.Println(r.Insert(1))
	fmt.Println(r.data)
	fmt.Println(r.Remove(0))
	fmt.Println(r.data)
	fmt.Println(r.Insert(2))
	fmt.Println(r.data)
	fmt.Println(r.Remove(1))
	fmt.Println(r.data)
	fmt.Println(r.data)
	fmt.Println(r.GetRandom())

}

// 方法一，数据实现
type RandomizedSet struct {
	intToIndex map[int]int
	data       []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	r := new(RandomizedSet)
	r.intToIndex = make(map[int]int)
	r.data = make([]int, 0)
	return *r
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.intToIndex[val]; ok {
		return false
	}
	index := len(this.data)
	this.data = append(this.data, val)
	this.intToIndex[val] = index
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.intToIndex[val]; !ok {
		return false
	}
	index := this.intToIndex[val]
	// 这里有个技巧，就是把要删除的元素和最后一个元素互换
	l := len(this.data)
	this.data[index] = this.data[l-1]
	this.intToIndex[this.data[index]] = index
	this.data = this.data[:l-1]
	delete(this.intToIndex, val)

	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	l := len(this.data)
	i := rand.Intn(l)
	return this.data[i]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
