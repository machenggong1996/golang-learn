package main

import (
	"container/list"
	"fmt"
)

func print(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func main() {
	l := list.New()
	l.PushBack(1) //尾插
	l.PushBack(2)
	print(l)

	fmt.Println("=========")

	l.PushFront(0) //头插
	print(l)

	fmt.Println("=========")

	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == 1 {
			l.InsertAfter(1.1, e)
		}

		if e.Value == 2 {
			l.InsertBefore(1.2, e)
		}
	}

	print(l)

	fmt.Println("=========")

	fmt.Println(l.Front().Value) //返回链表的第一个元素
	fmt.Println("=========")

	fmt.Println(l.Back().Value) //返回链表的最后一个元素
	fmt.Println("=========")

	l.MoveToBack(l.Front())
	print(l)

	fmt.Println("=========")

	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
