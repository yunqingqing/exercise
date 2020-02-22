"""
Input: "abca"
Output: True
Explanation: you could delete the character 'c'.
"""

import unittest

def valid_palindrome(string):
    i, j = 0, len(string)-1
    while i < j:
        if string[i] != string[j]:
            return is_palidrome(string, i+1, j) or is_palidrome(string, i, j-1)
        else:
            i += 1
            j -= 1
    return True

def is_palidrome(string, i, j):
    while i < j:
        if string[i] != string[j]:
            return False
        else:
            return is_palidrome(string, i+1, j-1)
    return True

class ValidPalindromeTestCase(unittest.TestCase):
    def test_valid_palindrome(self):
        self.assertEqual(valid_palindrome('abca'), True)

    def test_valid_palindrome_1(self):
        self.assertEqual(valid_palindrome('abcab'), False)

    def test_valid_palindrome_2(self):
        self.assertEqual(valid_palindrome('a'), True)

if __name__ == '__main__':
    unittest.main()
