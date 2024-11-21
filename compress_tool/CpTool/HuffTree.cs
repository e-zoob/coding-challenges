namespace CpTool;

public class HuffTree: IComparable
{
    private HuffBaseNode Root { get; set; }
    private int Weight { get; set; }

    public HuffTree(char element, int weight)
    {
        Root = new HuffLeafNode(element, weight);
        Weight = Root.Weight;
    }

    public HuffTree(HuffBaseNode left, HuffBaseNode right, int weight)
    {
        Root = new HuffInternalNode(left, right, weight);
        Weight = Root.Weight;
    }

    public int CompareTo(object? obj)
    {
        if (obj is not HuffTree huffTree) 
            throw new ArgumentException("Object is not a HuffTree");
        return Weight.CompareTo(huffTree.Weight);
    }
}