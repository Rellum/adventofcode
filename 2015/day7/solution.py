from typing import List, Union
from os import path
import pprint

def part1(input: List[str]) -> int:
    split = (l.split(" -> ") for l in input)

    wires = evaluate_wires({e[1]: e[0] for e in split})

    pprint.pprint(wires)

def part2(input: List[str]) -> int:
    split = (l.split(" -> ") for l in input)
    wires = {e[1]: e[0] for e in split}
    wires['b'] = 16076

    wires = evaluate_wires(wires)

    pprint.pprint(wires)
def evaluate_wires(wires: dict[str, Union[int, str]]) -> dict[str, Union[int, str]]:
    updated = True
    while updated:
        updated = False
        for wire, value in wires.copy().items():
            if isinstance(value, str) == False:
                continue
            if is_int(value):
                wires[wire] = int(value)
                updated = True
                continue
            if " AND " in value:
                split = value.split(" AND ")

                input1 = current_value(wires, split[0])
                input2 = current_value(wires, split[1])

                if isinstance(input1, str) or isinstance(input2, str):
                    continue

                wires[wire] = input1 & input2
                updated = True
                continue
            if " OR " in value:
                split = value.split(" OR ")

                input1 = current_value(wires, split[0])
                input2 = current_value(wires, split[1])

                if isinstance(input1, str) or isinstance(input2, str):
                    continue

                wires[wire] = input1 | input2
                updated = True
                continue
            if " LSHIFT " in value:
                split = value.split(" LSHIFT ")

                input1 = current_value(wires, split[0])
                input2 = current_value(wires, split[1])

                if isinstance(input1, str) or isinstance(input2, str):
                    continue

                wires[wire] = input1 << input2
                updated = True
                continue
            if " RSHIFT " in value:
                split = value.split(" RSHIFT ")

                input1 = current_value(wires, split[0])
                input2 = current_value(wires, split[1])

                if isinstance(input1, str) or isinstance(input2, str):
                    continue

                wires[wire] = input1 >> input2
                updated = True
                continue
            if "NOT " in value:
                input1 = current_value(wires, value.removeprefix("NOT "))

                if isinstance(input1, str):
                    continue

                wires[wire] = 2**16 + ~input1
                updated = True
                continue

            input1 = current_value(wires, value)

            if isinstance(input1, str):
                continue

            wires[wire] = input1
            updated = True
            continue

    pprint.pprint(wires)

def is_int(s: Union[int, str]) -> bool:
    try:
        int(s)
        return True
    except ValueError:
        return False

def current_value(wires: dict[str, Union[int, str]], s: str) -> Union[int, str]:
    try:
        res = wires[s]
    except KeyError:
        res = s
    if is_int(res):
        return int(res)
    return res

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