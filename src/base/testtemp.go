package main

type set struct {
}

func (n *set) test() string {
	return ""
}
func main() {
	var _ = set{}.test()
	var s set
	var _ = s.test()
}
