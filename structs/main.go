package main


type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insert(d string) {
	newNode := Node{data: d, next: nil}
	if list.head == nil {
		list.head = &newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = &newNode
	}
}

func (list *LinkedList) print() {
	current := list.head
	for current != nil {
		print(current.data+" ")
		current = current.next
	}
	println()
}


func main() {
	list := LinkedList{}
	list.insert("Elian")
	list.insert("Renteria")
	list.insert("is")
	list.insert("cool")
	list.print()
}