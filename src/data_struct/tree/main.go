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

import "fmt"

type Avltree struct {
	value     int
	hight     int
	LeftNode  *Avltree
	RightNode *Avltree
}

func (node *Avltree) FindPar(RootPar *Avltree, Root *Avltree) *Avltree {
	if RootPar == nil {
		return nil
	}
	if node.value > Root.value {
		return node.FindPar(Root, Root.RightNode)
	} else if node.value < Root.value {
		return node.FindPar(Root, Root.LeftNode)
	} else {
		return RootPar
	}
}

//右旋转
func (node *Avltree) RightRotate(NodeParent *Avltree) {
	//NodeGraP := NodeParent.FindPar(nil, Root)
	var TempNode *Avltree
	TempNode = node.LeftNode
	node.LeftNode = TempNode.RightNode
	TempNode.RightNode = node
	//node.LeftNode=nil
	NodeParent.LeftNode = TempNode
}

func (node *Avltree) RRRotate(NodeParent *Avltree) {
	if node.LeftNode.RightNode == nil {
		node.RightRotate(NodeParent)
	} else {
		TempNode := node.LeftNode
		node.LeftNode = TempNode.RightNode
		node.LeftNode.LeftNode = TempNode
		TempNode.RightNode = nil
		node.RightRotate(NodeParent)
	}
}

func (node *Avltree) LeftRotate(NodeParent *Avltree) {
	var TempNode *Avltree
	TempNode = node.RightNode
	node.RightNode = TempNode.LeftNode
	TempNode.RightNode = node
	NodeParent.RightNode = TempNode
}
func (node *Avltree) LeftLeftTor(NodeParent *Avltree) {
	var TempNode *Avltree
	TempNode = node
	node.LeftNode = TempNode.RightNode
	TempNode.RightNode = node
	NodeParent = TempNode
}
func (node *Avltree) LeftSecTor(Nodeparent *Avltree) {
	var TempNode *Avltree
	BalanceP := node.LeftNode.LeftNode.hight - node.LeftNode.RightNode.hight
	if BalanceP == 1 {
		node.LeftLeftTor(Nodeparent)
	} else if BalanceP == -1 {
		//	if BalanceP==-1{
		TempNode = node.LeftNode
		node.LeftNode = TempNode.RightNode
		node.LeftNode.LeftNode = TempNode
		node.LeftLeftTor(Nodeparent)
	} else {
		panic("this exist unused erro")
	}
	//node.LeftLeftTor(Nodeparent)
}

func (Root *Avltree) MidOrder() {
	if Root != nil {
		fmt.Println(Root.value)
		Root.LeftNode.MidOrder()
		Root.RightNode.MidOrder()
	}
}

func main() {
	var Root *Avltree
	Root = &Avltree{value: 5}
	Root.RightNode = &Avltree{value: 6}
	Root.LeftNode = &Avltree{value: 3}
	Root.LeftNode.LeftNode = &Avltree{value: 2}
	Root.LeftNode.LeftNode.RightNode = &Avltree{value: 1}

	Root.MidOrder()
	fmt.Println("=====")
	Root.LeftNode.Tor(Root)
	Root.MidOrder()
	//fmt.Println(Root)

}
