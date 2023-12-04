from typing import List
from os import path
import re
import math

def part1(input: List[str]) -> int:
    res = 0
    for card in input:
        score = 0
        for i in set(card["winning"]).intersection(set(card["selected"])):
            if score == 0:
                score = 1
            else:
                score *= 2
        res += score

    return res

def part2(input: List[str]) -> int:
    res = [1 for card in input]
    for idx, card in enumerate(input):
        for i, _ in enumerate(set(card["winning"]).intersection(set(card["selected"]))):
            card_index = idx + i + 1
            if card_index > len(res):
                break
            res[card_index] = res[card_index] + res[idx]

    return sum(res)

def file_input(filename: str) -> List[str]:
    fn = path.join(path.dirname(__file__), filename)
    with open(fn) as i:
        return parse(i.read())

def challenge_input() -> List[str]:
    return file_input("challenge-input.txt")

def example_input() -> List[str]:
    return file_input("example-input.txt")

def parse(input: str) -> List:
    res = []
    for line in input.splitlines():
        match = re.match(r'^Card +([0-9]+): ([0-9 ]+) \| ([0-9 ]+)$', line)
        res += [{
            "card": int(match.groups()[0]),
            "winning": [int(s) for s in match.groups()[1].split()],
            "selected": [int(s) for s in match.groups()[2].split()],
        }]

    return res

def run():
    input = challenge_input()
    print(part1(input))
    print(part2(input))

def test():
    input = example_input()
    print(part1(input))
    print(part2(input))

if __name__ == "__main__":
    test()