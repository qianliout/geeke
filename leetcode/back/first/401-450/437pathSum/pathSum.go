package main

import (
	"fmt"
)

/*
给定一个二叉树，它的每个结点都存放着一个整数值。
找出路径和等于给定数值的路径总数。
路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。
示例：
root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

	     10
	    /  \
	   5   -3
	  / \    \
	 3   2   11
	/ \   \

3  -2   1
返回 3。和等于 8 的路径有:
1.  5 -> 3
2.  5 -> 2 -> 1
3.  -3 -> 11
*/
func main() {
	root := TreeNode{Val: 3}
	res := pathSum(&root, 3)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
只需一次递归五行代码，用列表记录下当前结果即可，为什么要双重递归呢？
    sumlist[]记录当前路径上的和，在如下样例中：
          10
         /  \
        5   -3
       / \    \
      3   2   11
     / \   \
    3  -2   1

当DFS刚走到2时，此时sumlist[]从根节点10到2的变化过程为：

    10
    15 5
    17 7 2

当DFS继续走到1时，此时sumlist[]从节点2到1的变化为：

    18 8 3 1

因此，只需计算每一步中，sum在数组sumlist中出现的次数，然后与每一轮递归的结果相加即可

    count = sumlist.count(sum)等价于：

     count = 0
     for num in sumlist:
         if num == sum:
             count += 1

count计算本轮sum在数组sumlist中出现的次数

*/

func pathSum(root *TreeNode, sum int) int {
	// sumList := make([]int, 0)
	// n := dfs(root, &sumList, sum)
	// return n

	prefixSumCount := make(map[int]int)
	// 初值,前缀和为0时有一种方法，就是什么结点都不选
	prefixSumCount[0] = 1
	return recursionPathSum(root, prefixSumCount, sum, 0)
}

func dfs(root *TreeNode, sumList *[]int, sum int) int {
	if root == nil {
		return 0
	}
	// 这里一定要进行赋值操作，不然就不能得到正确的结果
	nl := append([]int{}, *sumList...)

	for i, n := range nl {
		nl[i] = n + root.Val
	}
	nl = append(nl, root.Val)
	count := 0
	for _, n := range nl {
		if n == sum {
			count++
		}
	}
	count += dfs(root.Left, &nl, sum)
	count += dfs(root.Right, &nl, sum)

	return count
}

// 前缀和的思想
func recursionPathSum(node *TreeNode, prefixSumCount map[int]int, target, currSum int) int {
	// 1.递归终止条件
	if node == nil {
		return 0
	}
	// 本层要做的事
	res := 0
	// 当前路径上的和
	currSum += node.Val
	// 看看root到当前节点这条路上是否存在节点前缀和加target为currSum的路径
	// 当前节点->root节点反推，有且仅有一条路径，如果此前有和为currSum-target,而当前的和又为currSum,两者的差就肯定为target了
	// currSum-target相当于找路径的起点，起点的sum+target=currSum，当前点到起点的距离就是target
	res += prefixSumCount[currSum-target]
	prefixSumCount[currSum] = prefixSumCount[currSum] + 1

	// 进入下一层
	res += recursionPathSum(node.Left, prefixSumCount, target, currSum)
	res += recursionPathSum(node.Right, prefixSumCount, target, currSum)

	// 回到本层把进入下一层的影响消掉
	prefixSumCount[currSum] = prefixSumCount[currSum] - 1
	return res
}
