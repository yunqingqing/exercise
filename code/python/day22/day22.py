import unittest

from utils.linked_list import LinkedList, Node


def get_length_of_loop_link(linklist):
    # setp1 是否有环， p1走一步， p2走两步，若相遇则有环
    p1 = p2 = linklist.head
    while p2 is not None:
        p2 = p2.get_next()
        if p2 is None:
            return -1

        p2 = p2.get_next()
        if p2 is None:
            return -1
        p1 = p1.get_next()

        if p1 == p2:
            has_ring = True
            break

    # setp2 求环长度length
    length = 1
    while p2.get_next() != p1:
        p2 = p2.get_next()
        length += 1

    # setp3 p1走l+1步，p2 走1步， 两者相遇即为入口
    p1 = p2 = linklist.head
    for i in range(length):
        p1 = p1.get_next()

    while p1 != p2:
        p1 = p1.get_next()
        p2 = p2.get_next()

    return p1.get_data()

class TestFindKthToTail(unittest.TestCase):
    def test_1(self):
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
        ll.append_node(node3)
        self.assertEqual(get_length_of_loop_link(ll), 3)

    def test_2(self):
        ll = LinkedList()
        ll.append(1)
        self.assertEqual(get_length_of_loop_link(ll), -1)


if __name__ == '__main__':
    unittest.main()
