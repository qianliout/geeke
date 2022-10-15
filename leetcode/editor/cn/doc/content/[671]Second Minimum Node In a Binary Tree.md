<p>Given a non-empty special binary tree consisting of nodes with the non-negative value, where each node in this tree has exactly <code>two</code> or <code>zero</code> sub-node. If the node has two sub-nodes, then this node's value is the smaller value among its two sub-nodes. More formally, the property&nbsp;<code>root.val = min(root.left.val, root.right.val)</code>&nbsp;always holds.</p>

<p>Given such a binary tree, you need to output the <b>second minimum</b> value in the set made of all the nodes' value in the whole tree.</p>

<p>If no such second minimum value exists, output -1 instead.</p>

<p>&nbsp;</p>

<p>&nbsp;</p> 
<p><strong class="example">Example 1:</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/15/smbt1.jpg" style="width: 431px; height: 302px;" /> 
<pre>
<strong>Input:</strong> root = [2,2,5,null,null,5,7]
<strong>Output:</strong> 5
<strong>Explanation:</strong> The smallest value is 2, the second smallest value is 5.
</pre>

<p><strong class="example">Example 2:</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/15/smbt2.jpg" style="width: 321px; height: 182px;" /> 
<pre>
<strong>Input:</strong> root = [2,2,2]
<strong>Output:</strong> -1
<strong>Explanation:</strong> The smallest value is 2, but there isn't any second smallest value.
</pre>

<p>&nbsp;</p> 
<p><strong>Constraints:</strong></p>

<ul> 
 <li>The number of nodes in the tree is in the range <code>[1, 25]</code>.</li> 
 <li><code>1 &lt;= Node.val &lt;= 2<sup>31</sup> - 1</code></li> 
 <li><code>root.val == min(root.left.val, root.right.val)</code>&nbsp;for each internal node of the tree.</li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>树 | 深度优先搜索 | 二叉树</details><br>

<div>👍 276, 👎 0</div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.0，第 12 期刷题打卡 [开始报名](https://mp.weixin.qq.com/s/eUG2OOzY3k_ZTz-CFvtv5Q)，点击这里体验 [刷题全家桶](https://labuladong.gitee.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg)。**

<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

前文 [手把手刷二叉树总结篇](https://labuladong.github.io/article/fname.html?fname=二叉树总结) 说过二叉树的递归分为「遍历」和「分解问题」两种思维模式，这道题需要用到「分解问题」的思维。

这题很有意思，按照这棵二叉树的特性，根节点无疑是最小的那个元素，但问题是第二小的那个元素在哪里呢？

如果根节点的左右子节点的值不同，根节点的值就是较小的那个节点（假设是左子节点）的值，那么较大的那个节点（右子节点）的值是不是就一定是第二小的那个元素？

不一定，第二小的值也可能在左子树中，这个节点是左子树中第二小的节点：

![](https://labuladong.github.io/algo/images/短题解/671_1.png)

类似的道理，如果根节点的左右子节点相同，那么需要把左右子树的第二小元素计算出来，比较其中较小的元素，作为整棵树的第二小元素：

![](https://labuladong.github.io/algo/images/短题解/671_2.png)

如何计算子树中的第二小元素？函数 `findSecondMinimumValue` 就是干这个的，递归调用子树即可算出来。

**标签：[二叉树](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2121994699837177859)**

## 解法代码

```java
class Solution {
    // 定义：输入一棵二叉树，返回这棵二叉树中第二小的节点值
    public int findSecondMinimumValue(TreeNode root) {
        if (root.left == null && root.right == null) {
            return -1;
        }
        // 左右子节点中不同于 root.val 的那个值可能是第二小的值
        int left = root.left.val, right = root.right.val;
        // 如果左右子节点都等于 root.val，则去左右子树递归寻找第二小的值
        if (root.val == root.left.val) {
            left = findSecondMinimumValue(root.left);
        }
        if (root.val == root.right.val) {
            right = findSecondMinimumValue(root.right);
        }
        if (left == -1) {
            return right;
        }
        if (right == -1) {
            return left;
        }
        // 如果左右子树都找到一个第二小的值，更小的那个是整棵树的第二小元素
        return Math.min(left, right);
    }
}
```

</details>
</div>





