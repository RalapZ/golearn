//##################################################################################################//
//                   			         ┌─┐       ┌─┐ + +                                          //
//                   			      ┌──┘ ┴───────┘ ┴──┐++                                         //
//                   			      │       ───       │++ + + +                                   //
//                   			      ███████───███████ │+                                          //
//                   			      │       ─┴─       │+                                          //
//                   			      └───┐         ┌───┘                                           //
//                   			          │         │   + +                                         //
//                   			          │         └──────────────┐                                //
//                   			          │                        ├─┐                              //
//                   			          │                        ┌─┘                              //
//                   			          └─┐  ┐  ┌───────┬──┐  ┌──┘  + + + +                       //
//                   			            │ ─┤ ─┤       │ ─┤ ─┤                                   //
//                   			            └──┴──┘       └──┴──┘  + + + +                          //
//                   			      神兽出没               永无BUG                                 //
//   Author: Ralap                                                                                  //
//   Date  : 2020/10/15                                                                             //
//##################################################################################################//
package main

import (
	"fmt"
)

type TreeNode struct {
	value     int
	hight     int
	RightNode *TreeNode
	LeftNode  *TreeNode
}

func UpdateHeight(node *TreeNode) int {
	LeftHigh, RightHigh, hight := 0, 0, 0
	if node.LeftNode != nil {
		LeftHigh = node.LeftNode.hight
	}
	if node.RightNode != nil {
		RightHigh = node.RightNode.hight
	}
	if LeftHigh >= RightHigh {
		hight = LeftHigh + 1
	} else {
		hight = RightHigh + 1
	}
	//fmt.Println("====")
	//fmt.Println(node.height)
	//fmt.Println(node.RightNode,node.LeftNode)
	return hight
}

func BalanceFactor(node *TreeNode) int {
	LeftHigh, RightHigh := 0, 0
	if node.LeftNode != nil {
		LeftHigh = node.LeftNode.hight
	}
	if node.RightNode != nil {
		RightHigh = node.RightNode.hight
	}
	return LeftHigh - RightHigh
}

func RightRotate(Root *TreeNode) *TreeNode {
	var TempNode *TreeNode
	TempNode = Root.LeftNode
	Root.LeftNode = TempNode.RightNode
	TempNode.RightNode = Root
<<<<<<< Updated upstream
	TempNode.RightNode.hight = UpdateHeight(TempNode.RightNode)
=======
	TempNode.RightNode.height=UpdateHeight(TempNode.RightNode)
	TempNode.height=UpdateHeight(TempNode)
>>>>>>> Stashed changes
	return TempNode
}
func LeftRotate(Root *TreeNode) *TreeNode {
	var TempNode *TreeNode
	TempNode = Root.RightNode
	Root.RightNode = TempNode.LeftNode
	TempNode.LeftNode = Root
<<<<<<< Updated upstream
	TempNode.LeftNode.hight = UpdateHeight(TempNode.LeftNode)
=======
	TempNode.LeftNode.height=UpdateHeight(TempNode.LeftNode)
	TempNode.height=UpdateHeight(TempNode)
>>>>>>> Stashed changes
	return TempNode
}

func (Root *TreeNode) MidOrder() {
	if Root != nil {
		fmt.Println(Root.value, Root.hight, Root.LeftNode, Root.RightNode)
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
			RightNode := Root.RightNode
			BF := BalanceFactor(RightNode)
			if BF == 1 {
				Root.RightNode = RightRotate(Root.RightNode)
<<<<<<< Updated upstream
=======
				//Root.RightNode.height = UpdateHeight(Root.RightNode)
>>>>>>> Stashed changes
			}
			Root = LeftRotate(Root)
		}
	} else {
		Root.LeftNode = Root.LeftNode.Add(value)
		BalanceF := BalanceFactor(Root)
		if BalanceF == 2 {
			LeftNode := Root.LeftNode
			BF := BalanceFactor(LeftNode)
			if BF == -1 {
				Root.LeftNode = RightRotate(Root.LeftNode)
>>>>>>> Stashed changes
			}
			Root = RightRotate(Root)
		}

	}
	//fmt.Println(Root)
<<<<<<< Updated upstream
	Root.hight = UpdateHeight(Root)
	return Root
}

func BalanceOper(Root *TreeNode) *TreeNode {
	BF := BalanceFactor(Root)
	//var TempNode *TreeNode
	if BF == -2 {
		BFR := BalanceFactor(Root.RightNode)
		if BFR == 1 {
			Root.RightNode = RightRotate(Root.RightNode)
		}
		Root = LeftRotate(Root)
	} else if BF == 2 {
		BFL := BalanceFactor(Root.LeftNode)
		if BFL == -1 {
			Root.LeftNode = LeftRotate(Root.LeftNode)
		}
		Root = RightRotate(Root)
	}
	Root.hight = UpdateHeight(Root)
	return Root
}

func RMNode(Root *TreeNode) *TreeNode {
	if Root.LeftNode == nil {
		return Root
	} else {
		TempNode := RMNode(Root.LeftNode)
		if TempNode == Root.LeftNode {
			Root.LeftNode = TempNode.RightNode
			_ = BalanceOper(Root)
		}
		return TempNode
	}
}

func DeleteNode(Root *TreeNode, value int) *TreeNode {
	if Root == nil {
		return nil
	}
	//Root的value等于value
	if Root.value == value {
		if Root.LeftNode == nil {
			return Root.RightNode
		} else if Root.RightNode == nil {
			return Root.LeftNode
		} else {
			TempNode := RMNode(Root.RightNode)
			TempNode.LeftNode = Root.LeftNode
			if TempNode != Root.RightNode {
				TempNode.RightNode = Root.RightNode
			}
			TempNode.hight = UpdateHeight(TempNode)
			return TempNode
		}
	} else if Root.value < value {
		//TempNode:=DeleteNode(Root.RightNode,value)
		Root.RightNode = DeleteNode(Root.RightNode, value)
		Root = BalanceOper(Root)

	} else {
		Root.LeftNode = DeleteNode(Root.LeftNode, value)
		Root = BalanceOper(Root)
	}
=======
	//Root.height = UpdateHeight(Root)
>>>>>>> Stashed changes
	return Root
}

func main() {
	var Root *TreeNode
	//Root = &TreeNode{value: 5}
	//Root.RightNode = &TreeNode{value: 6}
	//Root.LeftNode = &TreeNode{value: 3}
	//Root.LeftNode.LeftNode = &TreeNode{value: 2}
	//Root.LeftNode.LeftNode.RightNode = &TreeNode{value: 1}
	//fmt.Println(Root)
	Root = Root.Add(2)
	Root = Root.Add(3)
	Root = Root.Add(1)
	Root = Root.Add(5)
	Root = Root.Add(0)
	Root = Root.Add(4)
	Root = Root.Add(10)
	Root = Root.Add(11)
	Root.MidOrder()
	//fmt.Println(Root)
	//Root=DeleteNode(Root,11)
	fmt.Println("======")
	Root = DeleteNode(Root, 10)
	Root = DeleteNode(Root, 3)
	//fmt.Println(Root)
	Root.MidOrder()

	//fmt.Println("=====")
	//_ = RightRotate(Root)
	//Root.MidOrder()
}
