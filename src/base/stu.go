package main

import (
	"encoding/json"
	"fmt"
)

type StudentInfo struct {
	PID       int
	PPID      int
	name      string
	ChildNode []*StudentInfo
}
type PIDNode struct {
	PID        int
	StuInfo    *StudentInfo
	LeftChild  *PIDNode
	RightChild *PIDNode
}

func SearchStuTree(root *PIDNode, PPID int) (TEMPNODE *PIDNode) {
	if root == nil {
		return nil
	}
	//fmt.Println("root",root)
	if root.PID > PPID {
		TEMPNODE = SearchStuTree(root.LeftChild, PPID)
	} else if root.PID < PPID {
		TEMPNODE = SearchStuTree(root.RightChild, PPID)
	}
	return TEMPNODE
}

func StuInsert(root *PIDNode, node *PIDNode) {
	//fmt.Printf("stuinsert=%p\n", node)
	//fmt.Printf("stuinsert=%v\n", node.StuInfo)

	//k,_:=json.Marshal(root)
	//fmt.Println("stuinsert",string(k))
	//fmt.Println(node)
	if root.PID > node.PID {
		if root.LeftChild != nil {
			StuInsert(root.LeftChild, node)
		} else {
			root.LeftChild = node
		}

	} else if root.PID < node.PID {
		if root.RightChild != nil {
			StuInsert(root.RightChild, node)
		} else {
			root.RightChild = node
		}
	}
}

func (P *PIDNode) midorder() {
	fmt.Println(P.PID)
	if P.LeftChild != nil {
		P.LeftChild.midorder()
	}
	if P.RightChild != nil {
		P.RightChild.midorder()
	}

}

func main() {
	PIDTree := new(PIDNode)
	Data := []StudentInfo{}
	//stu := StudentInfo{1, 2,"ralap.z",nil}
	//nodeINFO := &StuTree{stu.PID, &stu, nil, nil}
	//node := SearchStuTree(Tree, stu.PPID)
	nodelist := []StudentInfo{
		{1, 2, "ralap.z", nil},
		{3, 1, "test1", nil},
		{2, 1, "test2", nil},
		{4, 3, "ralap1", nil}}

	node := new(PIDNode)
	//fmt.Printf("node1=%p\n",&nodelist[0])
	//fmt.Println(len(Data))
	n := 0
	for k, _ := range nodelist {
		N := &nodelist[k]
		//fmt.Printf("PIDTREE=%v\n", PIDTree)
		//fmt.Printf("N%v=%p\n",n,N)
		//fmt.Printf("node%v=%p\n",n,&nodelist[k])
		nodeINFO := &PIDNode{N.PID, N, nil, nil}
		fmt.Printf("nodeINFO=%p", nodeINFO)
		//fmt.Printf("NewNode%v=%v\n", n, nodeINFO.StuInfo)
		node = SearchStuTree(PIDTree, N.PPID)
		//k,_:=json.Marshal(PIDTree)
		//fmt.Println(string(k))
		if node == nil || Data == nil || len(Data) == 0 {
			Data = append(Data, *N)
		} else {
			fmt.Println(node)
			node.StuInfo.ChildNode = append(node.StuInfo.ChildNode, N)
		}
		//fmt.Println("round",n)
		fmt.Println("=============================================")
		fmt.Printf("round%v\n", n)
		StuInsert(PIDTree, nodeINFO)
		P, _ := json.Marshal(PIDTree)
		fmt.Println("PIDTree=", string(P))
		n++
	}
	PIDTree.midorder()

	//fmt.Println(PIDTree.RightChild)
	fmt.Println(Data)
}
