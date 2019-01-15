import unittest

def replace_space(input_string : str) -> str:
    if not isinstance(input_string, str):
        print("input not string.")
        return None
    return input_string.replace(" ", "%20")

class TestReplaceSpace(unittest.TestCase):
    def test_int_input(self):
        replace_space(12345)

    def test_has_space_string_input(self):
        self.assertEqual(replace_space("We Are Happy"), "We%20Are%20Happy")

    def test_no_space_string_input(self):
        self.assertEqual(replace_space("WeAreHappy"), "WeAreHappy")

if __name__ == "__main__":
    unittest.main()