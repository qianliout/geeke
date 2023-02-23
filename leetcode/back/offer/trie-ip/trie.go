package main

import (
	"sync"
)

func main() {
	tr := Constructor()
	tr.Insert("215.32.60.198/1")
}

type Trie struct {
	Root *Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	root := NewNode()
	// 标准的trie树是在需要时再生成子结点，生成结点总的耗时较大(pprof中表现就是NewNode函数耗时较长)，
	// 我们可以预计这棵trie树只会有32层,所以可以在初始化时提前把结点准备好，又因为层数越高，节点越多，使用率越低，
	// 所以做个取舍,提前生成前20层，
	constructorNode(root, 20)
	return Trie{Root: root}
}

func (tr *Trie) Insert(ip string) {
	addWord(tr.Root, ipToUint8Slice(ip), 0)
}

// 执行查询
func (tr *Trie) Query(ip string) (bool, error) {
	exit := findWord(tr.Root, ipToUint8Slice(ip), 0)
	return exit, nil
}

// 更新，我理解是使用提供的newCIDRs重新构造树，这种情况下，就重新构造并替换，
// 如果是插入新的值的话，就直接使用Insert方法就行
// 如果是把一个旧值改成新值，就先删除(Delete方法)，再插入(Insert方法)
// 这里只实现插入新值的情况
func (tr *Trie) Update(newCIDRs []string) error {
	for _, ip := range newCIDRs {
		tr.Insert(ip)
	}
	return nil
}

// 删除单个节点
func (tr *Trie) Delete(ip string) {
	deleteWord(tr.Root, ipToUint8Slice(ip), 0)
}

type Node struct {
	IsEnd bool

	// 在高并发下为了使锁的细腻度更低，在每个节点加读写锁，高并发下性能更好
	// 如果不是高并发下，这种在每个节点加锁比直接锁整个trie树性能更差，因为每个节点加锁，在查询，插入等操作时会递归加锁和释放锁，性能更差
	Lock sync.RWMutex

	// 对于标准的trie树，一般都是用map进行存储，但是本例中只会两种情况的值，如果使用map的话取值消耗时间较多(pprof中体现在mapaccess耗时太多)
	// 因为这里只会有两种值，所以可使用数组的方式进行存储，节约内存，取值更快，
	Next *[2]*Node
}

func NewNode() *Node {
	return &Node{Next: new([2]*Node)}
}

func addWord(root *Node, word []uint8, index int) {
	if len(word) == index {
		root.Lock.Lock()
		defer root.Lock.Unlock()
		root.IsEnd = true
		return
	}
	c := word[index]
	root.Lock.RLock()
	defer root.Lock.RUnlock()

	nex := root.Next[c]
	if nex == nil || nex.Next == nil {
		root.Lock.RUnlock() // go语言没有锁升级的慨念，所以这里要先释放读锁
		root.Lock.Lock()    // 再加写锁
		nex = NewNode()
		root.Next[c] = nex
		root.Lock.Unlock() // 释放写锁，这里不能用defer,因为defer是在程序退出前执行，如果用defer就在会下面读锁之后，形成死锁
		root.Lock.RLock()  // 再加读锁
	}
	addWord(nex, word, index+1)
}

func findWord(root *Node, word []uint8, index int) bool {
	root.Lock.RLock()
	defer root.Lock.RUnlock()
	if len(word) == index {
		return root.IsEnd
	}
	c := word[index]
	if root.Next == nil || root.Next[c] == nil {
		return false
	}
	nex := root.Next[c]
	return findWord(nex, word, index+1)
}

func deleteWord(root *Node, word []uint8, index int) {
	if len(word) == index {
		root.Lock.Lock()
		defer root.Lock.Unlock()
		// 对于ip掩码来说，并发量大，且更新频繁，不直接删除节点，只是把节点的值改了,这样避免了节点频繁删除后的GC负担
		root.IsEnd = false
		return
	}
	c := word[index]
	root.Lock.RLock()
	defer root.Lock.RUnlock()
	if root.Next[c] == nil {
		return
	}
	nex := root.Next[c]
	deleteWord(nex, word, index+1)
}

func constructorNode(root *Node, index int) {
	if index <= 0 {
		return
	}
	if root.Next == nil {
		root.Next = new([2]*Node)
	}
	root.Next[0] = new(Node)
	root.Next[1] = new(Node)
	constructorNode(root.Next[0], index-1)
	constructorNode(root.Next[1], index-1)
}
