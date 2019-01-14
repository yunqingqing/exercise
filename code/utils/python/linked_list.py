class Node(object):
    def __init__(self, data=None, next=None):
        self.data = data
        self.next = next

    def get_data(self):
        return self.data

    def set_data(self, data):
        self.data = data

    def get_next(self):
        return self.next

    def set_next(self, next):
        self.next = next

    def __str__(self):
        return str(self.data)


class LinkedList(object):
    def __init__(self):
        self.head = None

    def __len__(self):
        count = 0
        current = self.head
        while current:
            count += 1
            current = current.get_next()
        return count

    def __str__(self):
        current = self.head
        output = ""
        while current:
            output += str(current) + " -> "
            current = current.get_next()
        return output

    def set_head(self, head_node):
        self.head = head_node

    def pop(self):
        if self.head:
            self.head = self.head.get_next()
        else:
            raise IndexError("Unable to pop from empty list.")
    
    def delete(self, value):
        current = self.head
        prev = None
        while current:
            if current.get_data() == value:
                if prev:
                    prev.set_next(current.get_next())
                else:
                    self.head = current.get_next()

            prev = current
            current = current.get_next()

    def push(self, value):
        node = Node(value)
        node.set_next(self.head)
        self.set_head(node)

    def append(self, value):
        node = Node(value)
        current = self.head
        if not current:
            self.head = node
            return

        while current.get_next():
            current = current.get_next()

        current.set_next(node)

    def reverse(self):
        

if __name__ == "__main__":
    ll = LinkedList()

    print("Initial size: ", len(ll))
    ll.push(24)
    print("New size: ", len(ll))
    print("List content: ", ll)
    print("Pushing more")
    ll.push(6)
    ll.push(783)
    print("List content: ", ll)
    print("popping...")
    ll.pop()
    print("List content: ", ll)
    print("Deleting 24")
    ll.delete(24)
    print("List content: ", ll)
    print("Pushing another onto end.")
    ll.append(365)
    print("List content: ", ll)