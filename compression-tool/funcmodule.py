def CountCharOccurencies(file):
    char_occurrences = {}
    try:
        with open(file, 'r') as f:
            for line in f:
                for char in line:
                    if char in [' ', '\n']:
                        continue
                    if char in char_occurrences:
                        char_occurrences[char] += 1
                    else:
                        char_occurrences[char] = 1             
    except FileNotFoundError:
        print(f"File {file} not found.")
        
    return char_occurrences