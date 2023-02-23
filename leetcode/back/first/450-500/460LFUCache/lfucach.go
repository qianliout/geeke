package main

/*
以双向链表的方式
*/
type DoubleLinkedNode struct {
	Key       int
	Val       int
	Frequency int
	Pre       *DoubleLinkedNode
	Post      *DoubleLinkedNode
}

type LFUCache struct {
	FrequencyMap map[int]*DoubleLinkedNode
	ValueMap     map[int]*DoubleLinkedNode
	Cap          int
	Current      int
}

func Constructor(capacity int) LFUCache {
	lf := LFUCache{
		FrequencyMap: make(map[int]*DoubleLinkedNode),
		ValueMap:     make(map[int]*DoubleLinkedNode),
		Cap:          capacity,
		Current:      0,
	}
	return lf
}

func (this *LFUCache) Get(key int) int {
	v, ok := this.ValueMap[key]
	if !ok {
		return -1
	}
	// 更新频率,也就是从原来的链表调整到最新的链表表始的地方
	this.IncreaseFrequency(v, v.Frequency)
	return v.Val
}

func (this *LFUCache) Put(key int, value int) {
	if this.Cap <= 0 {
		return
	}
	n, ok := this.ValueMap[key]
	if !ok || n == nil {
		this.AddNewNode(&DoubleLinkedNode{
			Key:       key,
			Val:       value,
			Frequency: 1,
		})
	} else {
		this.AddOldNode(&DoubleLinkedNode{
			Key:       key,
			Val:       value,
			Frequency: n.Frequency + 1,
		})
	}

}

func (this *LFUCache) IncreaseFrequency(now *DoubleLinkedNode, prefre int) {
	// 删除原链表的结点
	freNode, ok := this.FrequencyMap[prefre]

	if ok && freNode != nil {
		cur := freNode.Post
		if cur == nil {
			delete(this.FrequencyMap, prefre)
		} else {
			for cur != nil {
				if cur.Key == now.Key {
					cur.Pre.Post = cur.Post
					break
				}
				cur = cur.Post
			}
		}
	}

	// 增加到新的联表的表头
	n, ok := this.FrequencyMap[prefre+1]
	if !ok || n == nil {
		this.FrequencyMap[prefre+1] = now
	} else {
		now.Post = n
		this.FrequencyMap[prefre+1] = now
	}
}

func (this *LFUCache) DeleteTail(node *DoubleLinkedNode) {

	if node.Post == nil {
		delete(this.FrequencyMap, node.Frequency)
	}
	cur := node
	for cur.Post.Post != nil {
		cur = cur.Post
	}
	cur.Post = nil
}

func (this *LFUCache) AddNewNode(node *DoubleLinkedNode) {
	// 如果容量不够,先腾出容量
	if this.Current >= this.Cap {
		// 找到最小的那个,只能是1,那这个就加不进去
		minNode, ok := this.FrequencyMap[1]
		if !ok || minNode == nil {
			3
			return
		}
		// 删除链表最后一个
		this.DeleteTail(minNode)
		this.Current--
	}

	n, ok2 := this.FrequencyMap[1]
	if !ok2 || n == nil {
		this.FrequencyMap[1] = node
	} else {
		this.IncreaseFrequency(n, 0)
	}
	this.ValueMap[node.Key] = node
	this.Current++
}

func (this *LFUCache) AddOldNode(node *DoubleLinkedNode) {
	this.ValueMap[node.Key] = node
	this.IncreaseFrequency(node, node.Frequency-1)
}
