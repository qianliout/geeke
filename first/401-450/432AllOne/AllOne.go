package main

func main() {

}

// 使用map加双端队列的方法
type AllOne struct {
	StringMap map[string]*DoubleListNode
	NodeMap   map[int]*DoubleListNode
	min       *DoubleListNode
	max       *DoubleListNode
}

type DoubleListNode struct {
	key  string
	val  int
	prev *DoubleListNode
	next *DoubleListNode
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	// head、tail节点不存具体的数据
	return AllOne{
		StringMap: make(map[string]*DoubleListNode),
		NodeMap:   make(map[int]*DoubleListNode),
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	var node *DoubleListNode
	if node, ok := this.StringMap[key]; ok { // key存在直接加1 需要维护链表有序
		node.val++

		cur := node.next
		for cur != this.DList.tail {
			if cur.val >= node.val { // 寻找下一个大于等于node的节点 将node移动到该节点前边
				break
			}
			cur = cur.next
		}
		if node.next != cur {
			this.RemoveNode(node)
			this.InsertNode(node, cur)
		}
	} else { // key不存在 生成节点插在head节点后
		node := &DoubleListNode{
			key: key,
			val: 1,
		}
		node.next = this.DList.head.next
		node.prev = this.DList.head
		this.DList.head.next.prev = node
		this.DList.head.next = node

		this.H[key] = node
		this.DList.length++
	}

	if node.val > this.max.val {
		this.max = node
	}
	if this.min == nil {
		this.min = node
	}

}

// 删除节点
func (this *AllOne) RemoveNode(node *DoubleListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 在cur节点前插入node节点
func (this *AllOne) InsertNode(node, cur *DoubleListNode) {
	node.next = cur
	node.prev = cur.prev
	cur.prev.next = node
	cur.prev = node
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	if node, ok := this.H[key]; ok { // key存在 key不存在不需处理
		if node.val <= 1 { // 删除节点
			this.RemoveNode(node)
			this.DList.length--
			delete(this.H, key)
		} else { // 节点减1 需要维护链表有序
			node.val--
			cur := node.prev
			for cur != this.DList.tail {
				if cur.val <= node.val { // 寻找上一个小于等于node的节点 将node移动到该节点前边
					break
				}
				cur = cur.prev
			}
			if node.prev != cur {
				this.RemoveNode(node)
				this.InsertNode(node, cur)
			}
		}
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.DList.length == 0 {
		return ""
	}
	return this.DList.tail.prev.key
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.DList.length == 0 {
		return ""
	}
	return this.DList.head.next.key
}
