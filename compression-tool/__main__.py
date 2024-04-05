import argparse
import funcmodule

parser = argparse.ArgumentParser()

parser.add_argument("file")

args = parser.parse_args()

file = args.file

funcmodule.CountCharOccurencies(file)
