from utils.linked_list import LinkedList, Node


def delete(linklist, node):
    if node == linklist.head:
        del linklist.head
        del linklist.tail
        
    if node.get_next() is not None:
        next = node.get_next()
        nnext = next.get_next()

        # 后节点前移
        node.set_data(next.get_data())
        node.set_next(nnext)

        # 其实这里node4的内存是释放不了的，因为之前还存在引用
        # del next
    else:
        # 要删除的是尾节点，从头遍历
        c = linklist.head
        while c.get_next() != node:
            c = c.get_next()
        c.set_next(None)

if __name__ == '__main__':
    ll = LinkedList()
    node1 = Node(1)
    node2 = Node(2)
    node3 = Node(3)
    node4 = Node(4)
    node5 = Node(5)
    ll.append_node(node1)
    ll.append_node(node2)
    ll.append_node(node3)
    ll.append_node(node4)
    ll.append_node(node5)
    delete(ll, node3)
    ll.print()
    print("======")
    delete(ll, node5)
    ll.print()
    print("======")
    ll2 = LinkedList()
    ll2.append_node(node1)
    delete(ll2, node1)