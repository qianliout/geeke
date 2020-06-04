package main

//type LRUCache struct {
//	ValueDict map[int]int
//	Remain    int
//	KeyList   []int
//}
//
//func Constructor(capacity int) LRUCache {
//	return LRUCache{
//		ValueDict: make(map[int]int),
//		Remain:    capacity,
//		KeyList:   make([]int, 0),
//	}
//}
//
//func (this *LRUCache) Get(key int) int {
//	val, ok := this.ValueDict[key]
//	if !ok {
//		return -1
//	}
//	this.KeyList = UpdateKeyList(this.KeyList, key)
//	// fmt.Printf("%+v\n", this.ValueDict)
//	// fmt.Printf("%+v\n", this.KeyList)
//	return val
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	_, ok := this.ValueDict[key]
//	if ok {
//		this.ValueDict[key] = value
//		// 更新KeyList
//		this.KeyList = UpdateKeyList(this.KeyList, key)
//	}
//	if !ok {
//		if this.Remain == 0 {
//			n := len(this.KeyList)
//			lastKey := this.KeyList[n-1]
//			delete(this.ValueDict, lastKey)
//			keyList := []int{key}
//			keyList = append(keyList, this.KeyList[0:n-1]...)
//			this.KeyList = keyList
//			this.ValueDict[key] = value
//		} else {
//			keyList := []int{key}
//			keyList = append(keyList, this.KeyList...)
//			this.KeyList = keyList
//			this.ValueDict[key] = value
//			this.Remain--
//		}
//	}
//	//fmt.Printf("%+v\n", this.ValueDict)
//}
//
//func UpdateKeyList(list []int, key int) []int {
//	var index int
//	var flag bool = false
//	for i, k := range list {
//		if k == key {
//			index = i
//			flag = true
//			break
//		}
//	}
//	if !flag {
//		return list
//	}
//	n := len(list)
//	keyList := []int{key}
//	keyList = append(keyList, list[0:index]...)
//	keyList = append(keyList, list[index+1:n]...)
//	return keyList
//}
//
