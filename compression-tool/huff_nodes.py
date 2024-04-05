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
#create methods to get the left and right children of the node
    def get_left(self):
        return self.left
    
    def get_right(self):
        return self.right
    
    def get_weight(self):
        return self.weight
    
    def isLeaf(self):
        return False