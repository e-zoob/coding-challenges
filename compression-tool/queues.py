from min_heap import MinHeap

class HuffPriorityQueue:
    def __init__(self):
        self._elements = MinHeap() 

    def enqueue_with_priority(self, value):
        self._elements.add(value)
    
    def dequeue(self):
        return self._elements.poll()
    
    def is_empty(self):
        return self._elements.is_empty()