package main

import (
	"fmt"
)

func main() {
	var cmd string
	fmt.Println("Welcome to Menu!")
	for {
		fmt.Scanln(&cmd)
		switch cmd {
		case "version":
			fmt.Println("Menu 1.0")
		case "hello":
			fmt.Println("world")
		case "help":
			fmt.Println("Try hello")
		case "quit":
			fmt.Println("bye")
			return
		case "testLinkList":
			TestLinkList()
		default:
			fmt.Println("emmmm...I can't understand.")
		}
	}
}

func TestLinkList() {
	list := &LinkList{}
	list.InitLinkList()
	fmt.Println("插入1-10")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 10; i++ {
		list.PushFront(nums[i])
	}
	list.Print()
	fmt.Println("删除4")
	list.Remove(4)
	list.Print()
	fmt.Println("删除7")
	list.Remove(7)
	fmt.Println("尾插100")
	list.PushBack(100)
	list.Print()
	fmt.Println("第5个元素处插入77")
	list.Insert(77, 5)
	list.Print()
}
