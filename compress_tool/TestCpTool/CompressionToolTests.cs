using FluentAssertions;
using CpTool;
namespace TestCpTool;

public class CompressionToolTests
{
    [Fact]
    public void CountFrequenciesShouldGiveCorrectCount()
    {
        const string input = "abbccc" +
                             "ddd d   ";
        var cpt = new CompressionTool();
        var expected = new Dictionary<char, int>()
        {
            { 'a', 1 },
            { 'b', 2 },
            { 'c', 3 },
            { 'd', 4 },
        };
        var result = cpt.GetFrequencies(input);
        
        result.Should().BeEquivalentTo(expected);
    }
}