package main

import "fmt"

func main() {
	head := NewUserHeadNode()
	u1 := &User{
		UserID: 1,
		Name:   "caketi",
	}
	u2 := &User{
		UserID: 1,
		Name:   "caketi",
	}
	u3 := &User{
		UserID: 1,
		Name:   "caketi",
	}
	head.FrontInsert(u1)
	head.FrontInsert(u2)
	head.FrontInsert(u3)
	head.Show()
}

type UserNode struct {
	Data *User
	Next *UserNode
}

type User struct {
	UserID int
	Name   string
}

func NewUserHeadNode() *UserNode {
	return &UserNode{}
}

func (head *UserNode) FrontInsert(u *User) {
	userNode := &UserNode{
		Data: u,
		Next: head.Next,
	}

	head.Next = userNode
}

func (node *UserNode) RearInsert() {

}

func (head *UserNode) Show() {
	temp := head.Next

	for temp != nil {
		fmt.Println(temp)
		temp = temp.Next
	}
}