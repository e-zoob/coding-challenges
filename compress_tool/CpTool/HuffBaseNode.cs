namespace CpTool;

public abstract class HuffBaseNode
{
    public bool IsLeaf { get; set; }
    public int Weight { get; protected init; }
}

public class HuffLeafNode : HuffBaseNode
{
    public char Element { get; set; }

    public HuffLeafNode(char element, int weight)
    {
        IsLeaf = true;
        Element = element;
        Weight = weight;
    }
}

public class HuffInternalNode : HuffBaseNode
{
    public HuffBaseNode Left { get; set; }
    public HuffBaseNode Right { get; set; }
    
    public HuffInternalNode(HuffBaseNode left, HuffBaseNode right, int weight)
    {
        IsLeaf = false;
        Left = left;
        Right = right;
        Weight = weight;
    }
}