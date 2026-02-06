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

func (ll *linkedList) getLastElementOfLinkedList() *linkedList {
	current := ll
	for {
		if current == nil {
			// tránh panic: xử lý trường hợp list rỗng
			return nil
		}
		if current.next == nil {
			return current
		} else {
			current = current.next
		}
	}
}

func (ll *linkedList) popLinkedList() *linkedList {
	if ll == nil || ll.next == nil {
		return nil
	}
	current := ll

	for current.next.next != nil {
		current = current.next
	}

	current.next = nil

	return ll
}

func revertLinkedList(ll *linkedList) linkedList {
	var revert linkedList
	for {
		if ll.getLastElementOfLinkedList() == nil {
			break
		} else {
			revert.appendLinkedList(*ll.getLastElementOfLinkedList())
			ll = ll.popLinkedList()
		}
	}

	return revert
}

func reverseLinkedListOfficial(head *linkedList) *linkedList {
	var prev *linkedList = nil
	current := head

	for current != nil {
		next := current.next // lưu lại next
		current.next = prev  // đảo chiều liên kết
		prev = current       // dịch prev sang current
		current = next       // dịch current sang next
	}

	return prev // prev giờ là node đầu của list đã đảo
}

func main() {
	ll := newLinkedList(1)

	ll.appendLinkedList(linkedList{value: 2})
	ll.appendLinkedList(linkedList{value: 3})
	ll.appendLinkedList(linkedList{value: 4})
	ll.appendLinkedList(linkedList{value: "ahihi"})

	ll.printOutAll()
	revertLinkedList := reverseLinkedListOfficial(ll)

	revertLinkedList.printOutAll()

}
