package doublelinkedlist

type Node struct {
    key   string
    value string
    next  *Node
    prev  *Node
}

type List struct {
	root Node
	len  int    // 链表长度
}

func NewList() *List {
	l := new(List)
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func NewNode(key, value string) *Node {
	n := new(Node)
	n.key = key
	n.value = value
	return n
}

func (n *Node) GetKey() string {
	return n.key
}

func (n *Node) SetValue(value string) {
	n.value = value
}

func (n *Node) GetValue() string {
	return n.value
}

func (l *List) Len() int {
	return l.len
}

// 返回双向链表头节点
func (l *List) Front() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// 在at节点插入e节点
func (l *List) Insert(e, at *Node) *Node {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	l.len++
	return e
}

// 移除e节点
func (l *List) Remove(e *Node) *Node {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	l.len--
	return e
}

// 把e节点移到链表末尾
func (l *List) MoveToBack(e *Node) {
	if l.root.prev == e {
		return
	}
	l.Insert(l.Remove(e), l.root.prev)
}

func (l *List) AppendToBack(e *Node) {
	l.Insert(e, l.root.prev)
}