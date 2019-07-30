import queue
import unittest


class Node(object):
    def __init__(self, num, step):
        self.num = num
        self.step = step


def num_squares(n):
    q = queue.Queue()
    q.put(Node(n, 1))
    visited = [False] * n

    while not q.empty():
        node = q.get()
        step = node.step
        num = node.num
        for i in range(1, n+1):
            next_num = num - i*i
            if next_num < 0:
                break
            if next_num == 0:
                return step

            if not visited[next_num]:
                q.put(Node(next_num, step + 1))
                visited[next_num] = True


class TestNumSquares(unittest.TestCase):
    def test_num_squares(self):
        self.assertEqual(num_squares(12), 3)
        self.assertEqual(num_squares(13), 2)
        self.assertEqual(num_squares(1), 1)


if __name__ == '__main__':
    unittest.main()
