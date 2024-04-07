class HuffLeafNode:
    def __init__(self, elem, wgt):
        self.element = elem
        self.weight = wgt

    def get_value(self):
        return self.element
    
    def get_weight(self):
        return self.weight
    
    def isLeaf(self):
        return True
    

class HuffInternalNode:
    def __init__(self, lft, rgt, wgt):
        self.left = lft
        self.right = rgt
        self.weight = wgt

    def get_left(self):
        return self.left
    
    def get_right(self):
        return self.right
    
    def get_weight(self):
        return self.weight
    
    def isLeaf(self):
        return False
    

class HuffTree:
    def __init__(self, elem, wgt):
        self.root = HuffLeafNode(elem, wgt)
    
    @classmethod
    def from_leaf_nodes(cls, left_node, right_node, wgt):
        return cls(HuffInternalNode(left_node, right_node, wgt))
    
    def get_weight(self):
        return self.root.get_weight()

    def __lt__(self, other):
        return self.get_weight() < other.get_weight()
    
    def __eq__(self, other):
        return self.get_weight() == other.get_weight()
    
    def __gt__(self, other):
        return self.get_weight() > other.get_weight()
