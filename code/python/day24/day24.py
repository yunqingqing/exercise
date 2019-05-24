import unittest

from utils.linked_list import LinkedList, Node


def merge(ll1, ll2):
    ll = LinkedList()

    c1 = ll1.head
    c2 = ll2.head

    while c1 and c2:
        if c1.get_data() < c2.get_data():
            ll.append(c1.get_data())
            c1 = c1.get_next()
        else:
            ll.append(c2.get_data())
            c2 = c2.get_next()

    while c1:
        ll.append(c1.get_data())
        c1 = c1.get_next()

    while c2:
        ll.append(c2.get_data())
        c2 = c2.get_next()
    return ll


class TestMergeLinklist(unittest.TestCase):
    def test_1(self):
        ll1 = LinkedList()
        ll1.append(2)
        ll1.append(4)
        ll1.append(6)
        ll1.append(8)
        ll1.append(10)

        ll2 = LinkedList()
        ll2.append(1)
        ll2.append(3)
        ll2.append(5)
        ll2.append(7)
        ll2.append(9)
        ll = merge(ll1, ll2)
        ll.print()


if __name__ == '__main__':
    unittest.main()
