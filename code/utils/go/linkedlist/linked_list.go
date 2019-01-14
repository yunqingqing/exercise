package linkedlist

type Node struct {
	next *Node
	data interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Append(value interface{}) {
	node := &Node{
		data: value,
	}
	if l.tail != nil {
		l.tail.next = node
		l.tail = node
	} else {
		l.head = node
		l.tail = node
	}
}

func (l *List) Push(value interface{}) {
	node := &Node{
		data: value,
		next: l.head,
	}
	l.head = node
}

func (l *List) Pop() {

}

func (l *List) IsEmpty() bool {
	if l.head == nil {
		return true
	}
	return false
}

func (l *List) Len() int {
	len := 0
	current := l.head
	for current != nil {
		len++
		current = current.next
	}
	return len
}

func (l *List) Delete(value interface{}) {
	current := l.head
	prev := &Node{}
	for current.next != nil {
		if current.data == value {
			if prev != nil {
				prev.next = current.next
			} else {
				l.head = current.next
			}
		}
		prev = current		
		current = current.next
	}
}
