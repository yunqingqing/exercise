"""
地上有一个m行n列的方格。一个机器人从坐标（0，0）的格子开始移动，每次可以向左右上下移动一格，
但不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格(35,37)
因为3+5+3+7=18但是不能进入(35,38)因为3+5+3+8=19。请问机器人能够到达多少个格子。

和day10练习思路一样，起始位置固定为(0,0)，如果能进入(i,j)的格子，
则再判断是否能进入(i, j-1) (i-1, j) (i+1, j), (i, j+1)的格子.
"""
import unittest


def moving_count(k, rows, cols):
    visited = [False] * cols*rows

    return moving_count_core(k, rows, cols, 0, 0, visited)


def moving_count_core(k, rows, cols, row, col, visited):
    count = 0

    if check(k, cols, rows, col, row, visited):
        visited[row*cols+col] = True
        count = 1 + moving_count_core(k, rows, cols, row+1, col, visited) + \
        moving_count_core(k, rows, cols, row, col+1, visited) + \
        moving_count_core(k, rows, cols, row-1, col, visited) + \
        moving_count_core(k, rows, cols, row, col-1, visited)
        # count = count + moving_count_core(k, rows, cols, row, col+1, visited)
        # count = count + moving_count_core(k, rows, cols, row-1, col, visited)
        # count = count + moving_count_core(k, rows, cols, row, col-1, visited)

    return count


def check(k, cols, rows, col, row, visited):
    if col < 0 or row < 0 or (get_sum(col) + get_sum(row)) > k or col >= cols or row >= rows or visited[row*cols + col]:
        return False
    return True


def get_sum(number):
    """获取数位和"""
    sum = 0
    while number > 0:
        sum += number % 10
        number //= 10

    return sum

class GetMovingCountTestCase(unittest.TestCase):
    def test1(self):
        self.assertEqual(moving_count(5, 10, 10), 21)
        self.assertEqual(moving_count(10, 1, 100), 29)
        self.assertEqual(moving_count(15, 100, 1), 79)
    
    def test_get_sum(self):
        self.assertEqual(get_sum(12), 3)
        self.assertEqual(get_sum(10), 1)
        self.assertEqual(get_sum(0), 0)


if __name__ == '__main__':
    unittest.main()