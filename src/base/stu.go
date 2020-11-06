package main

type StudentInfo struct {
	PID       int
	PPID      int
	ChildNode []StudentInfo
}
type StuTree struct {
	PID        int
	StuInfo    *StudentInfo
	LeftChild  *StuTree
	RightChild *StuTree
}

func SearchStuTree(root *StuTree, PPID int) *StuTree {
	if root == nil {
		return root
	}
	if root.PID > PPID {
		SearchStuTree(root.LeftChild, PPID)
	} else if root.PID < PPID {
		SearchStuTree(root.RightChild, PPID)
	}
	return root
}

func main() {
	Tree := &StuTree{}
	Data := []StudentInfo{}
	stu := StudentInfo{1, 2, nil}
	nodeINFO := &StuTree{stu.PID, &stu, nil, nil}
	node := SearchStuTree(Tree, stu.PPID)
	if node == nil {
		Data = append(Data, stu)
		node = nodeINFO
	} else {
		node.StuInfo.ChildNode = append(node.StuInfo.ChildNode, stu)
		temp := SearchStuTree(Tree, stu.PID)
		temp = nodeINFO
	}

	//stu:= StudentInfo{1,1,nil}
	//stu.ChildNode=append(stu.ChildNode, StudentInfo{2,1,nil})
	//stu.ChildNode=append(stu.ChildNode, StudentInfo{3,1,nil})
	//stu.ChildNode[1].ChildNode=append(stu.ChildNode[1].ChildNode, StudentInfo{3,1,nil})
	//a,_:=json.Marshal(stu)
	//fmt.Println(string(a))
	//
	//data:=[]StudentInfo{}

}
