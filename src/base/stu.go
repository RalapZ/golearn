package main

import "fmt"

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

func SearchStuTree(root *PIDNode, PPID int) *PIDNode {
	if root == nil {
		return nil
	}
	//fmt.Println("root",root)
	if root.PID > PPID {
		SearchStuTree(root.LeftChild, PPID)
	} else if root.PID < PPID {
		SearchStuTree(root.RightChild, PPID)
	}
	return root
}

func StuInsert(root *PIDNode, node *PIDNode) {
	fmt.Printf("stuinserver=%p\n", node)
	fmt.Printf("stuinserver=%v\n", node)
	if root.PID > node.PID {
		if root.LeftChild != nil {
			StuInsert(root.LeftChild, node)
		}
		root.LeftChild = node
	} else if root.PID < node.PID {
		if root.RightChild != nil {
			StuInsert(root.RightChild, node)
		}
		root.RightChild = node
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
		{2, 1, "test1", nil},
		{3, 1, "test2", nil},
		{3, 1, "ralap1", nil}}

	node := new(PIDNode)
	//fmt.Printf("node1=%p\n",&nodelist[0])
	//fmt.Println(len(Data))
	n := 0
	for k, _ := range nodelist {
		N := &nodelist[k]
		//fmt.Printf("N%v=%p\n",n,N)
		//fmt.Printf("node%v=%p\n",n,&nodelist[k])
		nodeINFO := &PIDNode{N.PID, N, nil, nil}
		fmt.Printf("NewNode%v=%v\n", n, nodeINFO)
		node = SearchStuTree(PIDTree, N.PPID)
		fmt.Printf("PIDTREE=%v\n", PIDTree)
		//fmt.Printf("search node=%v\n",node)
		//if PIDTree==nil{
		//	fmt.Println("PIDTree is null")
		//}
		if node == nil || Data == nil || len(Data) == 0 {
			Data = append(Data, *N)
		} else {
			fmt.Println(node)
			node.StuInfo.ChildNode = append(node.StuInfo.ChildNode, N)
		}
		StuInsert(PIDTree, nodeINFO)
		n++
	}
	fmt.Println(Data)
}
