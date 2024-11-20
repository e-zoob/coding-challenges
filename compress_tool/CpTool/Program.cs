
using CpTool;

Console.WriteLine("Enter the file name");
var fileName = Console.ReadLine();

var currentDirectory = Directory.GetCurrentDirectory();
//string rootDirectory = Path.GetFullPath(Path.Combine(currentDirectory, @"../../../.."));
var path = Path.Combine(currentDirectory, fileName!);
var input = File.ReadAllText(path);
var cpTool = new CompressionTool();

var result = cpTool.GetFrequencies(input);

foreach (var kv in result)
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}