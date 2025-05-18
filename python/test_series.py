import unittest
import series

class TestSeries(unittest.TestCase):
    def test_hello(self):
        self.assertEqual(series.hello(), "Hello, world!")

if __name__ == "__main__":
    unittest.main()
