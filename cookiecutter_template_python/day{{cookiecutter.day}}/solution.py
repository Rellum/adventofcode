from typing import List
from os import path

def part1(input: List[str]) -> int:
    return len(input)

def part2(input: List[str]) -> int:
    return len(input) ** 2

def file_input(filename: str) -> List[str]:
    fn = path.join(path.dirname(__file__), filename)
    with open(fn) as i:
        return parse(i.read())

def challenge_input() -> List[str]:
    return file_input("challenge-input.txt")

def example_input() -> List[str]:
    return file_input("example-input.txt")

def parse(i: str) -> List[str]:
    return i.splitlines()

def run():
    i = challenge_input()
    print(part1(i))
    print(part2(i))

def test():
    i = example_input()
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()