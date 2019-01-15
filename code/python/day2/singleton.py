import unittest

class Singleton(object):
    _instances = {}

    def __new__(cls, *args, **kw):
        if cls not in cls._instances:
            cls._instances[cls] = super(Singleton, cls).__new__(cls, *args, **kw)
        return cls._instances[cls]

class MyClass(Singleton):
    a = 1

class TestReplaceSpace(unittest.TestCase):

    def test_my_class(self):
        self.assertEqual(MyClass(), MyClass())

if __name__ == "__main__":
    unittest.main()