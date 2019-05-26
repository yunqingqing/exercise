"""
prev order: 1, 2, 4, 7, 3, 5, 6, 8
in order: 4, 7, 2, 1, 5, 3, 8, 6
"""

class Node:
    def __init__(self, value, left=None, right=None):
        self.value = value
        self.left = left
        self.right = right


def print_in_order(node):
    if node is None:
        return

    print_in_order(node.left)
    print(node.value)
    print_in_order(node.right)


def print_pre_orde(node):
    if node is None:
        return

    print(node.value)
    print_pre_orde(node.left)
    print_pre_orde(node.right)


def get_sub_in_order(in_order_list):
    pass


def get_sub_prev_order(prev_order_list, root_value):
    """
    return sub_prev_left_subtree, sub_prev_rigth_subtree
    """
    root_index = prev_order_list.index(root_value)
    return prev_order_list[:root], prev_order_list[root+1:]


def contruct_tree(in_order_list, prev_order_list):
    if len(prev_order_list) == 0:
        return

    # 先序第一个为root节点
    root_value = prev_order_list[0]
    root_node = Node(root_value)

    root_index_in_prev = in_order_list.index(root_value)
    left_subtree_prev_order = in_order_list[:root_index_in_prev]
    right_subtree_prev_order = in_order_list[root_index_in_prev+1:]

    left_subtree_in_order = prev_order_list[1:len(left_subtree_prev_order)+1]
    right_subtree_in_order = prev_order_list[len(left_subtree_prev_order)+1:]
    
    root_node.left = contruct_tree(left_subtree_prev_order, left_subtree_in_order)
    root_node.right = contruct_tree(right_subtree_prev_order, right_subtree_in_order)
    return root_node


if __name__ == '__main__':
    prev_order = [1, 2, 4, 7, 3, 5, 6, 8]
    in_order = [4, 7, 2, 1, 5, 3, 8, 6]
    node = contruct_tree(in_order, prev_order)
    print("inorder>>>")
    print_in_order(node)
    print()
    print("preorder>>>")
    print_pre_orde(node)