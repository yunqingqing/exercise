import unittest


def max_area_of_island(grid):
    """
    :type grid: List[List[int]]
    :rtype: int
    """
    length_y = len(grid)
    length_x = len(grid[0])
    visited = [[False]*length_x]*length_y

    def _dfs(row, col, visited=visited):
        if row >= 0 and row < length_x and col >= 0 and col < length_y and grid[row][col] == 1 and not visited[row][col]:
            visited[row][col] = True
            res = _dfs(row-1, col) + _dfs(row+1, col) + _dfs(row, col-1) + _dfs(row, col+1) + 1
            return res
        else:
            return 0

    areas = [0]
    print(length_x)
    print(length_y)
    print(visited)
    for col in range(length_x):
        for row in range(length_y):
            if not visited[row][col] and grid[row][col]==1:
                area = _dfs(row, col, visited)
                areas.append(area)

    return max(areas)


class MaxAreaOfIslandTestCase(unittest.TestCase):
    def test_max_area_of_island(self):
        self.assertEqual(max_area_of_island([[1,1,0,0,0],[1,1,0,0,0],[0,0,0,1,1],[0,0,0,1,1]]), 4)
        self.assertEqual(max_area_of_island([[0,0,0,0,0]]), 0)


if __name__ == '__main__':
    unittest.main()