from __future__ import annotations
from typing import List, Optional, TypeVar
import sys

T = TypeVar('T')

class MinHeap:

    nodes = List[T]

    def __init__(self, nodes: List[T] = []):
        self.nodes = []
        for node in nodes:
            self.add(node)

    def __len__(self):
        return len(self.nodes)
    
    def __get_left_child_index(self, parent_index: int) -> int:
        return 2 * parent_index + 1

    def __get_right_child_index(self, parent_index: int) -> int:
        return 2 * parent_index + 2 
    
    def __get_parent_index(self, child_index: int) -> int:
        return (child_index -1 ) // 2
    
    def __has_left_child(self, index: int) -> bool:
        return self.__get_left_child_index(index) < len(self.nodes)
    
    def __has_right_child(self, index: int) -> bool:
        return self.__get_right_child_index(index) < len(self.nodes)
    
    def __has_parent(self, index: int) -> bool:
        return self.__get_parent_index(index) >= 0
    
    def __left_child(self, index: int) -> Optional[T]:
        return self.nodes[self.__get_left_child_index(index)] if self.__has_left_child(index) else -sys.maxsize

    def __right_child(self, index: int) -> Optional[T]:
        return self.nodes[self.__get_right_child_index(index)] if self.__has_right_child(index) else -sys.maxsize
    
    def __parent(self, index: int) -> Optional[T]:
        return self.nodes[self.__get_parent_index(index)] if self.__has_parent(index) else None 

    def __swap(self, first_index: int, second_index: int):
        if first_index >= len(self.nodes) or second_index >= len(self.nodes):
            print(f'first ({first_index}) or second ({first_index}) are invalid')
            return
        self.nodes[first_index], self.nodes[second_index] = self.nodes[second_index], self.nodes[first_index]

    def  __heapify_up(self, child: Optional[T] = None):
        if not child:
            child = len(self.nodes) - 1
        parent_index = self.__get_parent_index(child)
        if self.__parent(child) and self.nodes[child] < self.__parent(child):
            self.__swap(child, parent_index)
            self.__heapify_up(parent_index)
        return self.nodes
   
    def add(self, item:T):
        self.nodes.append(item)
        self.__heapify_up()
    
    def __heapify_down(self, index: Optional[T] = 0):
        if index >= len(self.nodes) or (not self.__has_left_child(index)and not self.__has_right_child(index)):
            return self.nodes
        smaller_child_index = self.__get_left_child_index(index)
        if self.__has_right_child(index) and self.__right_child(index) < self.__left_child(index):
            smaller_child_index = self.__get_right_child_index(index)
        if self.nodes[index] < self.nodes[smaller_child_index]:
            return self.nodes
        if self.nodes[index] > self.nodes[smaller_child_index]:
            self.__swap(index, smaller_child_index)
            return self.__heapify_down(smaller_child_index)
        
    def poll(self) -> Optional[T]:
        if self.is_empty():
            print('Heap is empty')
            return None
        
        removed_node = self.nodes[0]
        self.nodes[0] = self.nodes[- 1]
        
        del self.nodes[-1]
        self.__heapify_down()
        
        return removed_node
    
    def is_empty(self) -> bool:
        return not self.nodes
    
    def peek(self) -> Optional[T]:
        if self.is_empty():
            return None
        return self.nodes[0]
    
    def heapsort_aux(unsorted_input: List[T]) -> List[T]:
        heap = MinHeap(unsorted_input)
        sorted_input = []
        for _ in range(len(heap)):
            sorted_input.append(heap.poll())
        return sorted_input