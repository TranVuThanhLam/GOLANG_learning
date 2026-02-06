package main

import "fmt"

type linkedList struct {
	value any
	next  *linkedList
}

func newLinkedList(value any) *linkedList {
	var ll linkedList
	ll.value = value
	return &ll
}

func (ll *linkedList) printOutAll() {
	temp := ll
	for {
		fmt.Println(temp.value)
		if temp.next == nil {
			break
		} else {
			temp = temp.next
		}
	}
}

func (ll *linkedList) appendLinkedList(newLl linkedList) {
	current := ll
	for {
		if current.next == nil {
			current.next = &newLl
			break
		} else {
			current = current.next
		}
	}
}

func main() {
	ll := newLinkedList(1)

	ll.appendLinkedList(linkedList{value: 2})
	ll.appendLinkedList(linkedList{value: 3})
	ll.appendLinkedList(linkedList{value: 4})
	ll.appendLinkedList(linkedList{value: "ahihi"})

	ll.printOutAll()
}
