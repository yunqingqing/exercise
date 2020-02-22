"""
Given s="leetcode", return "leotcede"
"""

import unittest

def reverse_vowels(string):
    vowels = ['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U']

    i, j = 0, len(string)-1
    char_list = list(string)
    while i < j:
        if string[i] in vowels and string[j] in vowels:
            char_list[i], char_list[j] = char_list[j], char_list[i]
            i += 1
            j -= 1
        
        elif string[i] in vowels:
            j -= 1
        elif string[j] in vowels:
            i += 1
        else:
            j -= 1
            i += 1
    return "".join(char_list)


class ReverseVowelsTestCase(unittest.TestCase):
    def test_reverse_vowels(self):
        self.assertEqual(reverse_vowels("leetcode"), "leotcede")

    def test_reverse_vowels_1(self):
        self.assertEqual(reverse_vowels("a"), "a")

if __name__ == '__main__':
    unittest.main()
