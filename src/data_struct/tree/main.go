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

func (node *Avltree) RightTor(Root *Avltree, NodeParent *Avltree) {
	//NodeGraP := NodeParent.FindPar(nil, Root)
	var TempNode *Avltree
	TempNode = node
	if NodeParent.LeftNode == node {
		NodeParent.LeftNode = node.LeftNode
		NodeParent.LeftNode.RightNode = TempNode
	}
	//fmt.Println(NodeGraP, TempNode)
	//if NodeGraP.RightNode==NodeParent{
	//	NodeGraP.RightNode=node
	//	node.LeftNode
	//
	//}
}

func main() {

}
