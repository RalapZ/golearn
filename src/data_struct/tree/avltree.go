package main

import "fmt"

type TreeNode struct {
	value     int
	height    int
	RightNode *TreeNode
	LeftNode  *TreeNode
}

func UpdateHeight(node *TreeNode) int {
	LeftHith, RightHigh := 0, 0
	if node.LeftNode != nil {
		LeftHith = node.LeftNode.height
	}
	if node.RightNode != nil {
		RightHigh = node.RightNode.height
	}
	if LeftHith >= RightHigh {
		node.height = LeftHith + 1
	} else {
		node.height = RightHigh + 1
	}
	return node.height
}

func BalanceFactor(node *TreeNode) int {
	return node.LeftNode.height - node.RightNode.height
}

func RightRotate(Root *TreeNode) *TreeNode {
	var TempNode *TreeNode
	TempNode = Root.LeftNode
	Root.LeftNode = TempNode.RightNode
	TempNode.RightNode = Root
	Root.height = UpdateHeight(Root)
	TempNode.height = UpdateHeight(TempNode)
	return TempNode
}
func LeftRotate(Root *TreeNode) *TreeNode {
	var TempNode *TreeNode
	TempNode = Root.RightNode
	Root.RightNode = TempNode.LeftNode
	TempNode.LeftNode = Root
	Root.height = UpdateHeight(Root)
	TempNode.height = UpdateHeight(TempNode)
	return TempNode
}

func (Root *TreeNode) MidOrder() {
	if Root != nil {
		fmt.Println(Root.value)
		Root.LeftNode.MidOrder()
		Root.RightNode.MidOrder()
	}
}

func (Root *TreeNode) Add(value int) {
	if Root == nil {
		Root = &TreeNode{value: value}
		Root.height = UpdateHeight(Root)
		return
	}

	if value == Root.value {
		return
	} else if value > Root.value {
		Root.RightNode.Add(value)
		BalanceF := BalanceFactor(Root)
		if BalanceF == -2 {
			LeftNode := Root.RightNode
			BF := BalanceFactor(LeftNode)
			if BF == 1 {
				Root.RightNode = RightRotate(Root.RightNode)
			}
			Root = LeftRotate(Root)
		}
	} else {
		Root.LeftNode.Add(value)
		BalanceF := BalanceFactor(Root)
		if BalanceF == 2 {
			RightNode := Root.RightNode
			BF := BalanceFactor(RightNode)
			if BF == -1 {
				Root.LeftNode = RightRotate(Root.LeftNode)
			}
			Root = RightRotate(Root)
		}
	}
}

func main() {
	var Root *TreeNode
	Root = &TreeNode{value: 5}
	Root.RightNode = &TreeNode{value: 6}
	Root.LeftNode = &TreeNode{value: 3}
	Root.LeftNode.LeftNode = &TreeNode{value: 2}
	Root.LeftNode.LeftNode.RightNode = &TreeNode{value: 1}

	Root.MidOrder()
	fmt.Println("=====")
	_ = RightRotate(Root)
	Root.MidOrder()
}
