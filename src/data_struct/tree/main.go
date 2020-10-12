package main

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
func (node *Avltree) RightTor(NodeParent *Avltree) {
	//NodeGraP := NodeParent.FindPar(nil, Root)
	var TempNode *Avltree
	TempNode = node.LeftNode
	node.LeftNode = TempNode.RightNode
	TempNode.RightNode = node
	//node.LeftNode=nil
	NodeParent.LeftNode = TempNode
}

func (node *Avltree) Tor(NodeParent *Avltree) {
	if node.LeftNode.RightNode == nil {
		node.RightTor(NodeParent)
	} else {
		TempNode := node.LeftNode
		node.LeftNode = TempNode.RightNode
		node.LeftNode.LeftNode = TempNode
		TempNode.RightNode = nil
		node.RightTor(NodeParent)
	}
}

func (node *Avltree) LeftTor(NodeParent *Avltree) {
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

func main() {

}
