import unittest

from utils.linked_list import LinkedList, Node


def find_kth_to_tail(linklist, k):
    p1 = p2 = linklist.head
    if p1 is None:
        return

    for i in range(k-1):
        p1 = p1.get_next()

    while p1.get_next() is not None:
        p1 = p1.get_next()
        p2 = p2.get_next()

    return p2.get_data()


class TestFindKthToTail(unittest.TestCase):
    def test_1(self):
        ll = LinkedList()
        ll.append(1)
        ll.append(2)
        ll.append(3)
        ll.append(4)
        ll.append(5)
        self.assertEqual(find_kth_to_tail(ll, 2), 4)

    def test_2(self):
        ll = LinkedList()
        ll.append(1)
        self.assertEqual(find_kth_to_tail(ll, 1), 1)


if __name__ == '__main__':
    unittest.main()

    