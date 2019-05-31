"""
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符路径。
路径可以从矩阵任意一格开始，每一步可以在矩阵中向左、右、上下移动一格。
如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。
a b t g
c f c s 
j d e h 
包含一条字符串"bfce"的路径。但矩阵中不包含"abfb"的路径。
"""
import unittest


def has_path(matrix, cols, rows, string):
    visited = [False] * (rows*cols)
    path_len = 0
    for row in range(rows):
        for col in range(cols):
            # 从任意点出发
            if has_path_core(matrix, visited, cols, col, rows, row, string, path_len):
                return True

    return False

def has_path_core(matrix, visited, cols, col, rows, row, string, path_len):
    if path_len == len(string):
        return True

    has_path = False
    if col < 0 or row < 0 or col >= cols or row >= rows:
        return False

    if not visited[cols*row + col] and matrix[row][col] == string[path_len]:
        visited[cols*row + col] = True
        path_len += 1
        has_path = has_path_core(matrix, visited, cols, col+1, rows, row, string, path_len) or \
                   has_path_core(matrix, visited, cols, col, rows, row+1, string, path_len) or \
                   has_path_core(matrix, visited, cols, col-1, rows, row, string, path_len) or \
                   has_path_core(matrix, visited, cols, col, rows, row-1, string, path_len)
        if not has_path:
            # 不符合即回溯
            visited[cols*row + col] = False
            path_len -= 1

    return has_path


class HasPathTestCase(unittest.TestCase):
    def test_1(self):
        matrix = [['a', 'b', 't', 'g'],
                  ['c', 'f', 'c', 's'],
                  ['j', 'd', 'e', 'h']]

        self.assertEqual(has_path(matrix, 4, 3, 'bfce'), True)
        self.assertEqual(has_path(matrix, 4, 3, 'abfb'), False)

    def test_2(self):
        matrix = [['a', 'a', 'a', 'a'],
                  ['a', 'a', 'a', 'a'],
                  ['a', 'a', 'a', 'a']]

        self.assertEqual(has_path(matrix, 4, 3, 'aaaaaaaaaaaaaaaa'), False)


if __name__ == '__main__':
    unittest.main()
