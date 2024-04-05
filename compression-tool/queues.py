from collections import deque
from heapq import heappop, heappush

class PriorityQueue:
    def __init__(self):
        self._elements = []

    def enqueue_with_priority(self, value, priority):
        heappush(self._elements, (priority, value))
    
    def dequeue(self):
        return heappop(self._elements)