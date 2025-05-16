// package main

// type Message struct {
// 	Topic string
// 	Payload interface{}
// }

// type Subscriber struct {
// 	Channel chan interface{}
// 	Unsubscribe chan bool
// }

//-------------------------------------------------------------------------------------

// Linked Lists:

// type node struct {
// 	data int
// 	next *node
// }

// type linkedList struct {
// 	head   *node
// 	length int
// }

// func (l *linkedList) prepend(n *node) {
// 	second := l.head
// 	l.head = n
// 	l.head.next = second
// 	l.length++
// }

// func (l linkedList) printListData() {
// 	toPrint := l.head
// 	for l.length != 0 {
// 		fmt.Printf("%d ", toPrint.data)

// 		toPrint = toPrint.next
// 		l.length--
// 	}
// 	fmt.Printf("\n")
// }

// func (l *linkedList) deleteWithValue(value int) {
// 	previousToDelete := l.head
// 	for previousToDelete.next.data != value {
// 		previousToDelete = previousToDelete.next
// 	}
// 	previousToDelete.next = previousToDelete.next.next
// 	l.length--
// }

// func main() {
// 	mylist := linkedList{}
// 	node1 := &node{data: 18}
// 	node2 := &node{data: 20}
// 	node3 := &node{data: 30}
// 	mylist.prepend(node1)
// 	mylist.prepend(node2)
// 	mylist.prepend(node3)

// 	mylist.printListData()
// 	mylist.deleteWithValue(20)
// 	mylist.printListData()
// }

//-------------------------------------------------------------------------------------

// type node struct {
// 	data int
// 	next *node
// }

// type linkedList struct {
// 	head   *node
// 	length int
// }

// func (l *linkedList) prependLinkedList(n *node) {
// 	second := l.head
// 	l.head = n
// 	n.next = second
// 	l.length++
// }

// func (l *linkedList) popLinkedList(n *node) {
// 	l.head = n.next
// 	l.length--
// }

// func (l *linkedList) deleteNode(value int) {
// 	if l.length == 0 {
// 		return
// 	}
// 	if l.head.data == value {
// 		l.head = l.head.next
// 		l.length--
// 	}

// 	toDelete := l.head

// 	for value != toDelete.next.data {
// 		if toDelete.next.next == nil {
// 			return
// 		}
// 		toDelete = toDelete.next
// 	}
// 	toDelete.next = toDelete.next.next
// 	l.length--
// }

// func (l linkedList) printLinkedList() {
// 	toPrint := l.head
// 	for l.length != 0 {
// 		fmt.Printf("%d ", toPrint.data)
// 		toPrint = toPrint.next
// 		l.length--
// 	}
// 	fmt.Println()
// }

// func main() {
// 	myList := linkedList{}
// 	node1 := node{data: 1}
// 	node2 := node{data: 2}
// 	node3 := node{data: 3}
// 	node4 := node{data: 4}
// 	node5 := node{data: 5}
// 	node6 := node{data: 6}
// 	myList.prependLinkedList(&node1)
// 	myList.prependLinkedList(&node2)
// 	myList.prependLinkedList(&node3)
// 	myList.prependLinkedList(&node4)
// 	myList.prependLinkedList(&node5)
// 	myList.prependLinkedList(&node6)

// 	fmt.Printf("%d is the length of the linked list.\n", myList.length)
// 	myList.printLinkedList()

// 	myList.popLinkedList(&node6)
// 	myList.printLinkedList()
// 	fmt.Printf("%d is the length of the linked list.\n", myList.length)

// 	myList.deleteNode(3)
// 	myList.printLinkedList()
// 	myList.deleteNode(5)
// 	myList.printLinkedList()
// 	myList.deleteNode(234)
// 	myList.printLinkedList()

// }

// ------------------------------------------------------
// structs

// type circle struct {
// 	radius float64
// }

// type rectangle struct {
// 	length  float64
// 	breadth float64
// }
//----------------------------------------------------------
// MAPS
// package main

// import "fmt"

// func main() {
// 	mp := make(map[string]int)
// 	mp["one"] = 1
// 	mp["two"] = 2
// 	mp["three"] = 3
// 	fmt.Println(mp)
// 	delete(mp, "two")
// 	fmt.Println(mp)
// 	fmt.Println(mp["one"])
// 	val, ok := mp["three"]
// 	fmt.Println(val, ok)

// }

//------------------------------------------------------------
// ProgrammingWithPercy project:

package main

import (
	"log"
	"net/http"
)

func main() {
	setupAPI()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupAPI() {
	http.Handle("/", http.FileServer(http.Dir("./Client")))
}
