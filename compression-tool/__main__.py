import argparse
import functions_module
from huff_nodes import HuffLeafNode, HuffInternalNode
from collections import deque
# parser = argparse.ArgumentParser()

# parser.add_argument("file")

# args = parser.parse_args()

# file = args.file

chars = functions_module.count_char_occurencies("text.txt")

tree = functions_module.build_tree(chars)



