package ai

import "fmt"

//从网上抄、修改了一份GOLANG实现的Queue
//原网址：http://blog.csdn.net/a13601100861water/article/details/74421093

type Node struct {
	data *Cell
	next *Node
}

type Queue struct {
	head *Node
	end  *Node
}

func NewQueue() *Queue {
	q := &Queue{nil, nil}
	return q
}

func (q *Queue) push(data *Cell) {
	//fmt.Println("入列：", data)
	n := &Node{data, nil}
	if q.end == nil {
		q.head = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}
	return
}

func (q *Queue) pop() (*Cell, bool) {
	if q.head == nil {
		return nil, false
	}

	data := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.end = nil
	}
	return data, true
}

//可以取出F值最小的元素
func (q *Queue) popMinF() (*Cell, bool) {
	if q.head == nil {
		return nil, false
	}

	now := q.head
	min := q.head.data.GetF()
	for {
		if now.next != nil {
			if now.data.GetF() > now.next.data.GetF() {
				min = now.next.data.GetF()
			}
		} else {
			break
		}
		now = now.next
	}
	now = q.head
	var last *Node
	var data *Cell
	for {
		if now.data.GetF() == min {
			data = now.data
			if now == q.head {
				q.head = q.head.next
			} else {
				last.next = now.next
			}
			break
		}
		if now.next == nil {
			panic("这不可能发生")
		}
		last = now
		now = now.next
	}
	//fmt.Println("出列：", data)
	if q.head == nil {
		q.end = nil
	}
	return data, true
}

func (q *Queue) exist(pos [2]int) (exist bool) {
	now := q.head
	if now == nil {
		return false
	}
	for {
		if now.data.Pos == pos {
			return true
		}
		if now.next == nil {
			return false
		}
		now = now.next
	}
}

func (q *Queue) get(pos [2]int) *Cell {
	now := q.head
	for {
		if now.data.Pos == pos {
			return now.data
		}
		if now.next == nil {
			return nil
		}
		now = now.next
	}
}

func (q *Queue) String() string {
	s := ""
	now := q.head
	for now != nil {
		s += fmt.Sprint(now.data.Pos)
		now = now.next
	}
	return s
}
