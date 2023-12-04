from typing import List
from os import path
import re

def part1(input: List[str]) -> int:
    solution = 0
    for line in input:
        digits = re.findall(r'\d', line)
        solution += int(digits[0]) * 10
        solution += int(digits[-1])
        # print(f"{int(digits[0]) * 10} + {int(digits[-1])} = {int(digits[0]) * 10 + int(digits[-1])}")
    return solution

def part2(input: List[str]) -> int:
    solution = 0
    for line in input:
        first_index = len(line)
        first_number = 0
        last_index = -1
        last_number = 0

        for (term, digit) in [(i, i) for i in range(10)] + list(NUMBER_WORDS.items()):
            index = line.find(str(term))
            if index == -1:
                continue
            if index < first_index:
                first_index = index
                first_number = digit

            index = line.rfind(str(term))
            if last_index < index:
                last_index = index
                last_number = digit

        solution += int(first_number) * 10
        solution += int(last_number)
        # print(int(first_number) * 10 + int(last_number), line)
    return solution

def file_input(filename: str) -> List[str]:
    fn = path.join(path.dirname(__file__), filename)
    with open(fn) as i:
        return parse(i.read())

def challenge_input() -> List[str]:
    return file_input("challenge-input.txt")

def example1_input() -> List[str]:
    return file_input("example1-input.txt")

def example2_input() -> List[str]:
    return file_input("example2-input.txt")


def parse(i: str) -> List[str]:
    return i.splitlines()


def run():
    input = challenge_input()
    print(part1(input))
    print(part2(input))

def test():
    input1 = example1_input()
    print(part1(input1))
    input2 = example2_input()
    print(part2(input2))

NUMBER_WORDS = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}

if __name__ == "__main__":
    test()