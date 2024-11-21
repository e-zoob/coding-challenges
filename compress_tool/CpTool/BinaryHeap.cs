namespace CpTool;

public class BinaryHeap
{
    private List<(char, int)> Heap { get; } = [];
    
    public int Count => Heap.Count;

    public void Push((char c, int frequency) item)
    {
        Heap.Add(item);
        HeapifyUp(Heap.Count - 1);
    }

    private void HeapifyUp(int index)
    {
        while (index > 0)
        {
            var parentIndex = (index - 1) / 2;
            if (Heap[parentIndex].Item2 >= Heap[index].Item2)
                break;
            Swap(index, parentIndex);
            index = parentIndex;
        }
    }

    private void Swap(int index1, int index2)
    {
        (Heap[index1], Heap[index2]) = (Heap[index2], Heap[index1]);
    }
    
    //Add HeapifyDown 
    //Add remove
    //Add peek
}