from utils.linked_list import LinkedList

# def PrintListReversinglyRecurisively
def print_list_reversingly_recurisively(head):
    if head is None:
        return

    current = head
    next = current.get_next()
    if next is not None:
        print_list_reversingly_recurisively(next)
    
    print(current.get_data())


if __name__ == '__main__':
    ll = LinkedList()
    ll.append(1)
    ll.append(2)
    ll.append(3)
    ll.append(4)
    ll.append(5)
    # ll.print()

    print_list_reversingly_recurisively(ll.head)
