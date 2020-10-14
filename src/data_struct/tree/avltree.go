package main

import "fmt"

type TreeNode struct {
	value     int
	height    int
	RightNode *TreeNode
	LeftNode  *TreeNode
}

func UpdateHeight(node *TreeNode) int {
	LeftHigh, RightHigh := 0, 0
	if node.LeftNode != nil {
		LeftHigh = node.LeftNode.height
	}
	if node.RightNode != nil {
		RightHigh = node.RightNode.height
	}
	if LeftHigh >= RightHigh {
		node.height = LeftHigh + 1
	} else {
		node.height = RightHigh + 1
	}
	//fmt.Println("====")
	//fmt.Println(node.height)
	//fmt.Println(node.RightNode,node.LeftNode)
	return node.height
}

func BalanceFactor(node *TreeNode) int {
	LeftHigh, RightHigh := 0, 0
	if node.LeftNode != nil {
		LeftHigh = node.LeftNode.height
	}
	if node.RightNode != nil {
		RightHigh = node.RightNode.height
	}
	return LeftHigh - RightHigh
}

func RightRotate(Root *TreeNode) *TreeNode {
	var TempNode *TreeNode
	TempNode = Root.LeftNode
	Root.LeftNode = TempNode.RightNode
	TempNode.RightNode = Root
	//TempNode.RightNode.height=UpdateHeight(TempNode.RightNode)
	//TempNode.height=UpdateHeight(TempNode)
	return TempNode
}
func LeftRotate(Root *TreeNode) *TreeNode {
	var TempNode *TreeNode
	TempNode = Root.RightNode
	Root.RightNode = TempNode.LeftNode
	TempNode.LeftNode = Root
	//TempNode.LeftNode.height=UpdateHeight(TempNode.LeftNode)
	//TempNode.height=UpdateHeight(TempNode)
	return TempNode
}

func (Root *TreeNode) MidOrder() {
	if Root != nil {
		fmt.Println(Root.value, Root.height, Root.LeftNode, Root.RightNode)
		Root.LeftNode.MidOrder()
		Root.RightNode.MidOrder()
	}
}

func (Root *TreeNode) Add(value int) *TreeNode {
	if Root == nil {
		return &TreeNode{value, 1, nil, nil}
	}

	if value > Root.value {
		Root.RightNode = Root.RightNode.Add(value)
		BalanceF := BalanceFactor(Root)
		if BalanceF == -2 {
			LeftNode := Root.RightNode
			BF := BalanceFactor(LeftNode)
			if BF == 1 {
				Root.RightNode = RightRotate(Root.RightNode)
				Root.RightNode.height = UpdateHeight(Root.RightNode)
			}
			Root = LeftRotate(Root)
			//Root.height=UpdateHeight(Root)
			//Root.RightNode.height=UpdateHeight(Root.RightNode)
		}
	} else {
		Root.LeftNode = Root.LeftNode.Add(value)
		BalanceF := BalanceFactor(Root)
		if BalanceF == 2 {
			RightNode := Root.RightNode
			BF := BalanceFactor(RightNode)
			if BF == -1 {
				Root.LeftNode = RightRotate(Root.LeftNode)
				Root.LeftNode.height = UpdateHeight(Root.LeftNode)
			}
			Root = RightRotate(Root)
			//Root.height = UpdateHeight(Root)
			//Root.LeftNode.height=UpdateHeight(Root.LeftNode)
		}
	}
	//fmt.Println(Root)
	Root.height = UpdateHeight(Root)
	return Root
}

func main() {
	var Root *TreeNode
	//Root = &TreeNode{value: 5}
	//Root.RightNode = &TreeNode{value: 6}
	//Root.LeftNode = &TreeNode{value: 3}
	//Root.LeftNode.LeftNode = &TreeNode{value: 2}
	//Root.LeftNode.LeftNode.RightNode = &TreeNode{value: 1}
	fmt.Println(Root)
	Root = Root.Add(2)
	Root = Root.Add(3)
	Root = Root.Add(1)
	Root = Root.Add(5)
	Root = Root.Add(0)
	Root = Root.Add(4)
	Root = Root.Add(10)
	Root = Root.Add(11)
	Root.MidOrder()
	fmt.Println(Root)
	//fmt.Println("=====")
	//_ = RightRotate(Root)
	//Root.MidOrder()
}
