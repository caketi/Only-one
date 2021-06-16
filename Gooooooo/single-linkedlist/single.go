package main

import "fmt"

func main(){
	dl := NewDoubleLinkList()
	dl.Add(8)
	dl.Add(9999)
	dl.Show()
}



type DoubleLinkList struct {
	Head *Node
	Len  int
}
type Node struct {
	Item interface{}
	Pre  *Node
	Next *Node
}

func NewDoubleLinkList() *DoubleLinkList {
	return &DoubleLinkList{
		Head: &Node{
			Item: nil,
			Pre:  nil,
			Next: nil,
		},
		Len: 0,
	}

}

func (l *DoubleLinkList) Delete(item interface{}) (bool, *Node) {
	pre := l.Head
	curr := l.Head.Next
	for curr != nil {
		// 找到匹配的了
		if item == curr.Item {
			baseDelete(pre, curr.Next)
			l.Len--
			return true, curr
		}

		// 不匹配，pre和curr向后走
		pre = curr
		curr = curr.Next
	}
	// 循环结束，依然未找到
	fmt.Println("无法删除，未查到")
	return false, nil
}

func (l *DoubleLinkList) Search(item interface{}) (bool, *Node) {
	temp := l.Head.Next
	for temp != nil {
		if item == temp.Item {
			return true, temp
		}
		temp = temp.Next
	}
	fmt.Println("未查到")
	return false, nil
}

func (l *DoubleLinkList) Add(item interface{}) {
	newNode := &Node{
		Item: item,
	}
	last := l.GetLast()
	baseAdd(last, newNode, nil)
	l.Len++
}
func (l *DoubleLinkList) Show() {
	temp := l.Head.Next
	for temp != nil {
		fmt.Printf("%v\n", temp.Item)
		temp = temp.Next
	}

	fmt.Println("---")
}
func (l *DoubleLinkList) IsEmpty() bool {
	return l.Len == 0
}

func (l *DoubleLinkList) GetLast() *Node {
	temp := l.Head
	for temp.Next != nil {
		temp = temp.Next
	}
	return temp
}

func baseDelete(pre, next *Node) {
	pre.Next = next
	if next != nil {
		next.Pre = pre
	}

}
func baseAdd(pre, curr, next *Node) {
	pre.Next = curr
	curr.Pre = pre

	if next != nil {
		curr.Next = next
		next.Pre = curr
	}

}