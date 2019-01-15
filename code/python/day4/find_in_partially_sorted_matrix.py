import unittest

def find(l : list, x :int) -> bool:
    m = len(l) - 1
    n = len(l[0]) - 1
    row = 0
    column = n
    while column >= 0 and row <= m:
        value = l[row][column]
        if value == x:
            return True
        elif value > x:
            column = column - 1
        else:
            row = row + 1
    return False


class TestFind(unittest.TestCase):

    def test_in_list(self):
        l = [
                [1, 2, 8, 9],
                [2, 4, 9, 12],
                [4, 7, 10, 13],
                [6, 8, 11, 15]
            ]
        self.assertEqual(find(l, 10), True)

    def test_not_in_list(self):
        l = [
                [1, 2, 8, 9],
                [2, 4, 9, 12],
                [4, 7, 10, 13],
                [6, 8, 11, 15]
            ]
        self.assertEqual(find(l, 20), False)

    def test_find_min_number(self):
        l = [
                [1, 2, 8, 9],
                [2, 4, 9, 12],
                [4, 7, 10, 13],
                [6, 8, 11, 15]
            ]
        self.assertEqual(find(l, 1), True)

    def test_find_max_number(self):
        l = [
                [1, 2, 8, 9],
                [2, 4, 9, 12],
                [4, 7, 10, 13],
                [6, 8, 11, 15]
            ]
        self.assertEqual(find(l, 15), True)


if __name__ == "__main__":
    unittest.main()