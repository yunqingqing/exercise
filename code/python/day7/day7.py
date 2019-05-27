"""
in order seq: {4, 2, 8, 5, 9, 1, 6, 3, 7}
tree like:
          1
	  2       3
    4   5   6    7
       8 9    

case1: 有右子树，遍历最左节点
case2: 无右子树，该节点父节点为父节点的的左孩子，父节点即为下一个节点
case3: 无右子树，该节点为父节点的右孩子，则向上查找父节点，
       找到一个节点为父节点左孩子的节点，及为下一个节点
    
"""


class Node:
    def __init__(self, value, left=None, right=None, parent=None):
        self.value = value
        self.left = left
        self.right = right
        self.parent = parent


def connect(parent, left, right):
    parent.left = left
    parent.right = right
    if left is not None:
        left.parent = parent
    if right is not None:
        right.parent = parent


def get_next(node):
    if node.right is not None:
        # case1
        cur = node.right
        while cur.left:
            cur = cur.left
        return cur.value
    elif node.parent:
        # case2,3
        cur = node

        while cur.parent is not None and cur != cur.parent.left:
            cur = cur.parent

        res = cur.parent.value if cur.parent else None
        return res


if __name__ == '__main__':
    node1 = Node(1)
    node2 = Node(2)
    node3 = Node(3)
    node4 = Node(4)
    node5 = Node(5)
    node6 = Node(6)
    node7 = Node(7)
    node8 = Node(8)
    node9 = Node(9)

    connect(node1, node2, node3)
    connect(node2, node4, node5)
    connect(node3, node6, node7)
    connect(node5, node8, node9)

    print(get_next(node5))  # 9
    print(get_next(node7))  # None
    print(get_next(node9))  # 1
