import unittest

from utils.linked_list import LinkedList, Node


def reverse_linklist(linklist):
    prev = None
    cur = linklist.head
    if cur is None:
        return linklist

    next = cur.get_next()
    while cur is not None:
        next = cur.get_next()
        cur.set_next(prev)

        prev = cur
        cur = next
    linklist.head = prev
    return linklist


class TestReverseLinklist(unittest.TestCase):
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
        reverse_linklist(ll)
        ll.print()

    # def test_2(self):
    #     ll = LinkedList()
    #     ll.append(1)
    #     self.assertEqual(get_length_of_loop_link(ll), -1)


if __name__ == '__main__':
    unittest.main()
