package main

import (
	"fmt"
	json "github.com/json-iterator/go"
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
	//fmt.Printf("%v,%v\n",root,PPID)
	if root == nil {
		return nil
	}
	if root.PID > PPID {
		TEMPNODE = SearchStuTree(root.LeftChild, PPID)
	} else if root.PID < PPID {
		TEMPNODE = SearchStuTree(root.RightChild, PPID)
	} else {
		TEMPNODE = root
	}
	return TEMPNODE
}

func StuInsert(root *PIDNode, node *PIDNode) {
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

//func SetHandler(DATA []*StudentInfo){
//
//}

func main() {
	PIDTree := new(PIDNode)
	Data := []*StudentInfo{}

	nodelist := []StudentInfo{
		{13, 1, "myzone", nil},
		{1, 10, "ralap.z", nil},
		{3, 1, "test1", nil},
		{2, 1, "test2", nil},
		{4, 3, "ralap1", nil},
		{11, 12, "dd", nil},
	}
	node := new(PIDNode)
	for i, _ := range nodelist {
		fmt.Printf("nodelist %v   %p\n", i, &nodelist[i])
	}
	//n:=0
	for k, _ := range nodelist {
		fmt.Println("=============================================")
		N := &nodelist[k]
		nodeINFO := &PIDNode{N.PID, N, nil, nil}
		//fmt.Printf("nodeINFO=%v\n", nodeINFO)
		//fmt.Printf("%#v\n",Data)
		node = SearchStuTree(PIDTree, N.PPID)
		fmt.Println(node)
		if node == nil || Data == nil || len(Data) == 0 {
			//fmt.Println("first")
			Data = append(Data, N)
		} else {
			node.StuInfo.ChildNode = append(node.StuInfo.ChildNode, N)
			//fmt.Printf("nodeappend%v   %p\n",node.StuInfo.ChildNode,node)
		}
		StuInsert(PIDTree, nodeINFO)
		D, _ := json.Marshal(Data)
		fmt.Println("DATA", string(D))
		//n++
	}

	P, _ := json.Marshal(Data)
	fmt.Println(string(P))
	fmt.Println(len(Data))
	k := 0
	for k < len(Data) {
		//node=SearchStuTree(Data)
		node = SearchStuTree(PIDTree, Data[k].PPID)
		fmt.Println(node, Data[k])
		if node == nil || Data == nil || len(Data) == 0 {
			k++
			continue
		} else {
			node.StuInfo.ChildNode = append(node.StuInfo.ChildNode, Data[k])
			Data[k], Data[len(Data)-1] = Data[len(Data)-1], Data[k]
			Data = Data[:len(Data)-1]
			//fmt.Printf("Data=%v\n",len(Data))
		}
	}
	P, _ = json.Marshal(Data)
	fmt.Println(string(P))

}
