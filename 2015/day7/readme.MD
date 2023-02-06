<a name="readme-top"></a>

<!-- TITLE -->
## Advent of Code 2015 Day 7

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ul>
    <li><a href="#about">About</a></li>
    <li><a href="#task-description">Task description</a></li>
    <li><a href="#built-with">Built With</a></li>
    <li><a href="#usage">Usage</a></li>
  </ul>
</details>

<!-- ABOUT -->
## About

This is my attempt at a solution to the daily challenge. See the original challenge [here][Challenge-url].

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- TASK DESCRIPTION -->
## Task description

**--- Day 7: Some Assembly Required ---**

This year, Santa brought little Bobby Tables a set of wires and [bitwise logic gates][descLink1]!  Unfortunately, little Bobby is a little under the recommended age range, and he needs help assembling the circuit.

Each wire has an identifier (some lowercase letters) and can carry a [16-bit][descLink2] signal (a number from `0` to `65535`).  A signal is provided to each wire by a gate, another wire, or some specific value. Each wire can only get a signal from one source, but can provide its signal to multiple destinations.  A gate provides no signal until all of its inputs have a signal.

The included instructions booklet describes how to connect the parts together: `x AND y -> z` means to connect wires `x` and `y` to an AND gate, and then connect its output to wire `z`.

For example:

* `123 -> x` means that the signal `123` is provided to wire `x`.
* `x AND y -> z` means that the [bitwise AND][descLink3] of wire `x` and wire `y` is provided to wire `z`.
* `p LSHIFT 2 -> q` means that the value from wire `p` is [left-shifted][descLink4] by `2` and then provided to wire `q`.
* `NOT e -> f` means that the [bitwise complement][descLink5] of the value from wire `e` is provided to wire `f`.

Other possible gates include `OR` ([bitwise OR][descLink6]) and `RSHIFT` ([right-shift][descLink7]).  If, for some reason, you'd like to **emulate** the circuit instead, almost all programming languages (for example, [C][descLink8], [JavaScript][descLink9], or [Python][descLink10]) provide operators for these gates.

For example, here is a simple circuit:

```
123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i
```

After it is run, these are the signals on the wires:

```
d: 72
e: 507
f: 492
g: 114
h: 65412
i: 65079
x: 123
y: 456
```

In little Bobby's kit's instructions booklet (provided as your puzzle input), what signal is ultimately provided to **wire `a`**?

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- BUILT WITH -->
## Built With

[![Go][Python-shield]][Go-url]

I am using Python for this solution to brush up my skills.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE -->
## Usage

```bash
cat example-input.txt | go run main.go
# output: Answer
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
[Challenge-url]: https://adventofcode.com/2015/day/7
[descLink1]: https://en.wikipedia.org/wiki/Bitwise_operation
[descLink2]: https://en.wikipedia.org/wiki/16-bit
[descLink3]: https://en.wikipedia.org/wiki/Bitwise_operation#AND
[descLink4]: https://en.wikipedia.org/wiki/Logical_shift
[descLink5]: https://en.wikipedia.org/wiki/Bitwise_operation#NOT
[descLink6]: https://en.wikipedia.org/wiki/Bitwise_operation#OR
[descLink7]: https://en.wikipedia.org/wiki/Logical_shift
[descLink8]: https://en.wikipedia.org/wiki/Bitwise_operations_in_C
[descLink9]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Bitwise_Operators
[descLink10]: https://wiki.python.org/moin/BitwiseOperators
[Go-shield]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/
[Python-shield]: https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=white
[Python-url]: https://www.python.org/