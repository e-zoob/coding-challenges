import pytest
from functions_module import count_char_occurencies, build_tree
import os

class TestCountCharOccurencies:
    @pytest.fixture(autouse=True)
    def setup_and_teardown(self):
        self.test_file = 'test.txt'
        with open(self.test_file, 'w') as f:
            f.write('abcabcxxx   ')
        yield
        os.remove(self.test_file)

    def test_count_char_occurencies(self):
        expected = {'a': 2, 'b': 2, 'c': 2, 'x': 3 }
        actual = count_char_occurencies(self.test_file)
        assert actual == expected



class TestBuildTree:
    def test_build_tree(self):
        char_occurrences = {'a': 5, 'b': 9, 'c': 12, 'd': 13, 'e': 16, 'f': 45}
        tree = build_tree(char_occurrences)

        assert tree.get_weight() == sum(char_occurrences.values())