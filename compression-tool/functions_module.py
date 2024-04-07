from queues import HuffPriorityQueue
from huff_nodes import HuffLeafNode, HuffInternalNode, HuffTree

def count_char_occurencies(file):
    char_occurrences = {}
    try:
        with open(file, 'r') as f:
            for line in f:
                for char in line:
                    if char in [' ', '\n']:
                        continue
                    if char in char_occurrences:
                        char_occurrences[char] += 1
                    else:
                        char_occurrences[char] = 1             
    except FileNotFoundError:
        print(f"File {file} not found.")
        
    return char_occurrences


def build_tree(char_occurences):
    queue = HuffPriorityQueue()
    for char, freq in char_occurences.items():
        leaf_node = HuffLeafNode(char, freq)
        huff_tree = HuffTree(leaf_node, freq)
        queue.enqueue_with_priority(huff_tree)

    while not queue.is_empty():
        tree1 = queue.dequeue()
        if queue.is_empty():
            return tree1
        tree2 = queue.dequeue()

        merged_weight = tree1.get_weight() + tree2.get_weight()
        merged_node = HuffInternalNode(tree1.root, tree2.root, merged_weight)
        merged_tree = HuffTree(merged_node, merged_weight)
        queue.enqueue_with_priority(merged_tree)
    return queue.dequeue()