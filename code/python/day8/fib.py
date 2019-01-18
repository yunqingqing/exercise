import unittest

def fib(num :int) -> int:
    if num == 0:
        return 0
    if num == 1:
        return 1
    
    fib_n_minus_one = 1
    fib_n_minus_two = 0
    fib_n = 0
    for i in range(num-1):
        fib_n = fib_n_minus_one + fib_n_minus_two
        fib_n_minus_two = fib_n_minus_one
        fib_n_minus_one = fib_n
    return fib_n

class TestFib(unittest.TestCase):

    def test_fib(self):
        self.assertEqual(fib(1), 1)
        self.assertEqual(fib(1), 1)
        self.assertEqual(fib(1), 1)
        self.assertEqual(fib(1), 1)
        self.assertEqual(fib(1), 1)

if __name__ == "__main__":
    unittest.main()