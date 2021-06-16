package main

import "fmt"

func main() {
	q := NewQueue()
	fmt.Printf("q:%+v", q)
	q.Push(8)
	q.Push(88)
    q.Show()
	
	fmt.Printf("q:%+v", q)
}

type Queue struct {
	Arr  [8]int
	Head int
	Rear int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(n int) {
	last := q.Rear
	q.Arr[last] = n
	q.Rear++
}

func (q *Queue) Pop() {

}

func (q *Queue) Show() {
	for i := q.Head; i != q.Rear; i ++{
		fmt.Println("q:\n", q)
	}

}

func (q *Queue) IsEmpty() {

}