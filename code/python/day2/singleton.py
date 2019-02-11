import unittest
import time
from threading import Lock, Thread
from concurrent.futures import ThreadPoolExecutor, as_completed

class Singleton(object):
    _instances = {}
    _lock = Lock()

    def __new__(cls, *args, **kw):
        # 如果不存在实例，则需要加锁创建新的实例,减少程序不必要的加锁
        if cls not in cls._instances:
            # sleep 1s 现象更容易出现
            # time.sleep(1)
            with cls._lock:
                # Double-checked_locking pattern.
                if cls not in cls._instances:
                    cls._instances[cls] = super(Singleton, cls).__new__(cls, *args, **kw)
        return cls._instances[cls]

class MyClass(Singleton):
    a = 1


def get_myclass_instance():
    return MyClass()

class TestSingleton(unittest.TestCase):

    def test_thread_safe(self):
        all_instances = []
        has_diff_ins = False

        with ThreadPoolExecutor(max_workers=4) as executor:
            all_task = [executor.submit(get_myclass_instance) for i in range(100)]
            for future in as_completed(all_task):
                all_instances.append(future.result())

        for index1, ins1 in enumerate(all_instances):
            for index2, ins2 in enumerate(all_instances):
                if index1 != index2 and ins1 is not ins2:
                    has_diff_ins = True
                    break

        self.assertEqual(has_diff_ins, False)



if __name__ == "__main__":
    unittest.main()