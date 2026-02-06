package main

import "fmt"

// 📝 Yêu cầu:

// Viết một chương trình tạo linked list đơn với các thao tác sau:

//     Thêm phần tử vào cuối danh sách (append)

//     In ra toàn bộ danh sách (print)

//     Xóa phần tử có giá trị cụ thể (deleteByValue)

//     Tìm phần tử theo giá trị (find)

//     Đếm số phần tử trong danh sách (length)

type node struct {
	data any
	next *node
}

type linkedList struct {
	head   *node
	lenght int
}

func (ll *linkedList) append(data any) {
	// tạo node mới từ data đã thêm
	newNode := &node{data: data}

	// Nếu head trống thì thêm node mới tạo vào head
	if ll.head == nil {
		ll.head = newNode
		// Nếu head không trống thì tìm tới node cuối cùng
	} else {
		// tạo node hiện tại bằng node head (node đầu tiên của linkedlist)
		current := ll.head
		// lặp việc gán current thành next của chính nó đến khi nào next == nil (nghĩa là đã đến node cuối cùng của linkedlist)
		for current.next != nil {
			current = current.next
		}

		// vì node hiện đại đã là node cuối nên chỉ cần thêm node mới tạo vào next là xong
		current.next = newNode
	}

	ll.lenght++
}

func (ll *linkedList) printAllLinkedList() {
	if ll.head == nil {
		fmt.Println("linked list is nil")
		return
	}

	current := ll.head
	for current.next != nil {
		fmt.Println(current.data)
		current = current.next
	}
	fmt.Println(current.data)

}

func (ll *linkedList) deleteByValue(data any) {
	// nếu head mà rổng thì trả về luôn là linked list bị rổng
	if ll.head == nil {
		fmt.Println("linked list is nil")
	}

	// nếu node đầu tiên luôn bằng với data thì xóa tới khi nào không giống nữa thì thôi
	for ll.head.data == data {
		ll.head = ll.head.next
		ll.lenght--
	}

	// xử lý các node tiếp theo ( lúc này head chắc chắn đã khác với data)
	current := ll.head
	for current.next != nil {
		if current.next.data == data {
			if current.next.next != nil {
				current.next = current.next.next
			} else {
				current.next = nil
			}
		}
		current = current.next
	}
}

func main() {
	var ll linkedList = linkedList{}
	ll.append(1)
	ll.append(2)
	ll.append(3)
	ll.append(4)
	ll.append(4)
	ll.append(4)
	ll.append(4)
	ll.append(4)
	ll.append("hello")
	ll.append(5)
	ll.deleteByValue(4)
	ll.printAllLinkedList()
	// fmt.Println(ll.head.data)
	// fmt.Println(ll.head.next.data)
	// fmt.Println(ll.head.next.next.data)
	// fmt.Println(ll.head.next.next.next.data)
	// fmt.Println(ll.head.next.next.next.next.data)
}
