import pytest
from min_heap import MinHeap

def test_add():
    heap = MinHeap()
    heap.add(5)
    assert heap.peek() == 5

    heap.add(3)
    assert heap.peek() == 3

    heap.add(7)
    assert heap.peek() == 3

def test_poll():
    heap = MinHeap()
    heap.add(5)
    heap.add(3)
    heap.add(7)

    assert heap.poll() == 3
    assert heap.poll() == 5
    assert heap.poll() == 7

def test_is_empty():
    heap = MinHeap()
    assert heap.is_empty()

    heap.add(5)
    assert not heap.is_empty()

    heap.poll()
    assert heap.is_empty()

def test_heapsort_aux():
    unsorted_list = [5, 3, 7, 1, 2]
    sorted_list = MinHeap.heapsort_aux(unsorted_list)
    assert sorted_list == [1, 2, 3, 5, 7]

def is_heap_invariant_satisfied(heap):
    for i in range(len(heap.nodes)):
        if heap._MinHeap__has_left_child(i):
            assert heap.nodes[i] <= heap.nodes[heap._MinHeap__get_left_child_index(i)]
        if heap._MinHeap__has_right_child(i):
            assert heap.nodes[i] <= heap.nodes[heap._MinHeap__get_right_child_index(i)]

def test_heap_invariant():
    heap = MinHeap()
    heap.add(5)
    heap.add(3)
    heap.add(7)
    heap.add(1)
    heap.add(6)

    is_heap_invariant_satisfied(heap)

    heap.poll()
    is_heap_invariant_satisfied(heap)