namespace CpTool;

public class CompressionTool
{
    public Dictionary<char, int> GetFrequencies(string input)
    {
        return input
            .Where(x => !char.IsWhiteSpace(x))
            .GroupBy(x => x)
            .ToDictionary(x => x.Key, x => x.Count());
    }
}