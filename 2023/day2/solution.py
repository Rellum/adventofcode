from typing import List
from os import path
import re
import math

def part1(input) -> int:
    res = []
    for game in input:
        possible = True
        for draw in game["draws"]:
            if draw["red"] > 12:
                possible = False
                break
            if draw["green"] > 13:
                possible = False
                break
            if draw["blue"] > 14:
                possible = False
                break
        if possible:
            res += [game["game"]]
    return sum(res)

def part2(input: List[str]) -> int:
    res = []
    for game in input:
        maximums = {colour: 0 for colour in COLOURS}
        for draw in game["draws"]:
            for colour in COLOURS:
                if draw[colour] > maximums[colour]:
                    maximums[colour] = draw[colour]
        res += [math.prod(maximums.values())]
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
        match = re.match(r'^Game (\d+): (.+)$', line)
        draws = []
        for draw_string in match.groups()[1].split(";"):
            draw = {}
            for colour in COLOURS:
                colour_match = re.search(f"(\d+) {colour}", draw_string)
                draw[colour] = 0 if colour_match is None else int(colour_match.groups()[0])
            draws += [draw]

        res += [{
            "game": int(match.groups()[0]),
            "draws": draws,
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

COLOURS = ["red", "green", "blue"]

if __name__ == "__main__":
    test()