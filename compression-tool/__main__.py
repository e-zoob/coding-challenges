import argparse
import funcmodule
from huff_nodes import HuffLeafNode, HuffInternalNode
from collections import deque
# parser = argparse.ArgumentParser()

# parser.add_argument("file")

# args = parser.parse_args()

# file = args.file

chars = funcmodule.count_char_occurencies("text.txt")

tree = funcmodule.build_tree(chars)



