package main

import "fmt"

type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

type LinkList struct {
	Head   *LinkNode
	Length int
}

//链表初始化为空
func (list *LinkList) InitLinkList() {
	list.Head = new(LinkNode)
	list.Length = 0
}

//链表判空
func (list *LinkList) Empty() bool {
	return list.Length == 0
}

//获取链表大小
func (list *LinkList) Size() int {
	return list.Length
}

//链表头插
func (list *LinkList) PushFront(data interface{}) {
	newNode := &LinkNode{Data: data}
	newNode.Next = list.Head.Next
	list.Head.Next = newNode
	list.Length++
}

//链表尾插
func (list *LinkList) PushBack(data interface{}) {
	p := list.Head
	for p.Next != nil {
		p = p.Next
	}
	newNode := &LinkNode{Data: data}
	p.Next = newNode
	list.Length++
}

//链表指定中间位置插入
func (list *LinkList) Insert(data interface{}, location int) bool {
	if location <= 0 || location > list.Length {
		fmt.Println("Illegal location!")
		return false
	}
	p := list.Head
	newNode := &LinkNode{Data: data}
	for i := 1; i < location; i++ {
		p = p.Next
	}
	newNode.Next = p.Next
	p.Next = newNode
	list.Length++
	return true
}

//链表删除指定元素
func (list *LinkList) Remove(data interface{}) bool {
	pre := list.Head
	p := list.Head.Next
	for p != nil && p.Data != data {
		pre = p
		p = p.Next
	}
	if p == nil {
		fmt.Println("链表无指定元素")
		return false
	} else {
		pre.Next = p.Next
		list.Length--
		return true
	}
}

//打印链表
func (list *LinkList) Print() {
	if list.Empty() {
		fmt.Println("空")
	} else {
		p := list.Head.Next
		for p != nil {
			fmt.Print(p.Data)
			fmt.Print(" ")
			p = p.Next
		}
		fmt.Print("\n")
	}
}
