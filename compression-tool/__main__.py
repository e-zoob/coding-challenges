import argparse

parser = argparse.ArgumentParser()

parser.add_argument("file")

args = parser.parse_args()

file = args.file

print(file)