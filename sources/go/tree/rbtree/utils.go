package rbtree

import (
	"container/list"
	"fmt"
	"math"
)

// 获取兄弟节点
func getSibling(n *node) *node {

	// 如果是左子节点
	if n.parent != nil && n == n.parent.left {
		return n.parent.right
	}

	// 如果是右子节点
	if n.parent != nil && n == n.parent.right {
		return n.parent.left
	}

	return nil
}

/**
	本文件内的函数都是二叉树常用的一些工具
 */

// 对于所有的树来说，遍历都是通用的

// 前序遍历
func preOrderTraverse(n *node) {
	if n == nil {
		return
	}
	fmt.Printf("%v ", n.data)			// 前序遍历就是从node开始遍历，所以要先打印
	preOrderTraverse(n.left)
	preOrderTraverse(n.right)
}

// 中序遍历
func inOrderTraverse(n *node) {
	if n == nil {
		return
	}
	// 会产生式升序结果
	inOrderTraverse(n.left)
	fmt.Printf("%v ", n.data)
	inOrderTraverse(n.right)

	// 会产生降序结果
	//InOrderTraverse(node.right)
	//fmt.Printf("%v ", node.data)
	//InOrderTraverse(node.left)
}

// 后序遍历
func postOrderTraverse(n *node) {
	if n == nil {
		return
	}
	postOrderTraverse(n.left)
	postOrderTraverse(n.right)
	fmt.Printf("%v ", n.data)
}

// 层序遍历：附带换行打印
/**
实现思路：无法用递归实现
1 将各节点入队
2 循环执行以下操作，直到队列为空
	取出队头节点出队，进行访问
	将队头节点的左子节点入队
	将队头节点的右子节点入队
*/
func levelOrderTraverse(n *node) {

	if n == nil {
		fmt.Println("传入节点为空")
		return
	}

	// 制作一个队列
	queue := list.New()
	queue.PushBack(n)

	// 记录当前行的元素个数，以便打印时按层换行
	levelLength := 1		// 每层存储的元素个数

	for queue.Len() != 0 {

		// 取出队头
		queueHead := queue.Remove(queue.Front())
		tempNode := queueHead.(*node)				// 类型断言
		fmt.Printf("%v ", tempNode.data)

		levelLength--

		// 将队头的左右子节点插入队列
		if tempNode.left != nil {
			queue.PushBack(tempNode.left)
		}
		if tempNode.right != nil {
			queue.PushBack(tempNode.right)
		}

		if levelLength == 0 {
			levelLength = queue.Len()
			fmt.Println()					// 遇到当前行数已经为0，则打印一次换行
		}

	}

}

// 插入时的递归函数
func insertRC(currentNode *node, insertNode *node){

	// 如果插入数据和当前节点数据相同
	if currentNode.data == insertNode.data {
		currentNode = insertNode
		return
	}

	parentNode := currentNode

	if currentNode.data > insertNode.data {		// 向左查找
		if currentNode.left == nil {
			insertNode.parent = parentNode
			currentNode.left = insertNode
		} else {
			parentNode = currentNode.left
			insertRC(currentNode.left, insertNode)
		}
	}

	if currentNode.data < insertNode.data { // 向右查找
		if currentNode.right == nil {
			insertNode.parent = parentNode
			currentNode.right = insertNode
		} else {
			parentNode = currentNode.right
			insertRC(currentNode.right, insertNode)
		}
	}
}

// 计算二叉树的高度：递归方式
func heightByGe(n *node) int{
	if n == nil {
		return 0
	}
	r := 1 + math.Max(float64(heightByGe(n.left)), float64(heightByGe(n.right)))
	return int(r)
}

// 计算二叉树的高度：迭代方式
func heightByRC(n *node) int {

	if n == nil {
		return 0
	}

	// 层序遍历
	queue := list.New()		// 制作一个队列
	queue.PushBack(n)

	height := 0				// 树的高度
	levelLength := 1		// 每层存储的元素个数

	for queue.Len() != 0 {

		queueHead := queue.Remove(queue.Front())	// 队首出队
		tempNode := queueHead.(*node)				// 类型断言

		levelLength--

		if tempNode.left != nil {
			queue.PushBack(tempNode.left)
		}
		if tempNode.right != nil {
			queue.PushBack(tempNode.right)
		}

		if levelLength == 0 {				// 准备访问下一层
			levelLength = queue.Len()
			height++
		}
	}
	return height
}

// 判断一棵树是否是完全二叉树：即最后一层的叶子节点是向左对其
func isCompleteTree(n *node) bool{
	if n == nil {
		return false
	}

	// 层序遍历
	queue := list.New()		// 制作一个队列
	queue.PushBack(n)

	isLeaf := false			// 当前节点是否是叶子节点
	for queue.Len() != 0 {

		queueHead := queue.Remove(queue.Front())	// 队首出队
		tempNode := queueHead.(*node)				// 类型断言

		if tempNode.left == nil && tempNode.right != nil {
			return false
		}

		// 如果已经要求是叶子节点，但是又不是叶子节点
		if isLeaf && (tempNode.left == nil && tempNode.right == nil)  {
			return false
		}

		// 要求下次遍历必须是叶子节点的情况
		if tempNode.right == nil || (tempNode.left != nil && tempNode.right == nil) {
			isLeaf = true
		}

		// 原层序遍历代码
		if tempNode.left != nil {
			queue.PushBack(tempNode.left)
		}
		if tempNode.right != nil {
			queue.PushBack(tempNode.right)
		}

	}

	return true
}

// 反转二叉树：所有节点的左右节点交换。其本质是遍历二叉树，使用任何遍历方式都行
func invertBinaryTree(n *node) *node{

	if n ==nil {
		return n
	}

	// 交换
	tempNode := n.left
	n.left = n.right
	n.right = tempNode

	preOrderTraverse(n.left)
	preOrderTraverse(n.right)

	return n
}