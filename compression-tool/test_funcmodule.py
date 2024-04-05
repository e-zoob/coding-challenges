import pytest
from funcmodule import CountCharOccurencies
import os

class TestCountCharOccurencies:
    @pytest.fixture(autouse=True)
    def setup_and_teardown(self):
        self.test_file = 'test.txt'
        with open(self.test_file, 'w') as f:
            f.write('abcabcxxx   ')
        yield
        os.remove(self.test_file)

    def test_CountCharOccurencies(self):
        expected = {'a': 2, 'b': 2, 'c': 2, 'x': 3 }
        actual = CountCharOccurencies(self.test_file)
        assert actual == expected