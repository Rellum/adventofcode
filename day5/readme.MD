<a name="readme-top"></a>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## Advent of Code Day 5

This is my attempt at a solution to the daily challenge. See the original challenge [here][Challenge-url].

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- TASK DESCRIPTION -->
## My solution to the Advent of Code Day 5 challenge

--- Day 5: Supply Stacks ---
The expedition can depart as soon as the final supplies have been unloaded 
from the ships. Supplies are stored in stacks of marked **crates**, but because 
the needed supplies are buried under many other crates, the crates need to 
be rearranged.

The ship has a **giant cargo crane** capable of moving crates between stacks. 
To ensure none of the crates get crushed or fall over, the crane operator 
will rearrange them in a series of carefully-planned steps. After the 
crates are rearranged, the desired crates will be at the top of each stack.

The Elves don't want to interrupt the crane operator during this delicate 
procedure, but they forgot to ask her **which** crate will end up where, and 
they want to be ready to unload them as soon as possible so they can 
embark.

They do, however, have a drawing of the starting stacks of crates **and** the 
rearrangement procedure (your puzzle input). For example:

```
    [D]    
[N] [C]    
[Z] [M] [P]
1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
```

In this example, there are three stacks of crates. Stack 1 contains two 
crates: crate `Z` is on the bottom, and crate `N` is on top. Stack 2 contains 
three crates; from bottom to top, they are crates `M`, `C`, and `D`. Finally, 
stack 3 contains a single crate, `P`.

Then, the rearrangement procedure is given. In each step of the procedure, 
a quantity of crates is moved from one stack to a different stack. In the 
first step of the above rearrangement procedure, one crate is moved from 
stack 2 to stack 1, resulting in this configuration:

```
[D]        
[N] [C]    
[Z] [M] [P]
1   2   3
```

In the second step, three crates are moved from stack 1 to stack 3. Crates 
are moved **one at a time**, so the first crate to be moved (`D`) ends up below 
the second and third crates:

```
        [Z]
        [N]
    [C] [D]
    [M] [P]
1   2   3
```

Then, both crates are moved from stack 2 to stack 1. Again, because crates 
are moved **one at a time**, crate `C` ends up below crate `M`:

```
        [Z]
        [N]
[M]     [D]
[C]     [P]
1   2   3
```

Finally, one crate is moved from stack 1 to stack 2:

```
        [Z]
        [N]
        [D]
[C] [M] [P]
1   2   3
```

The Elves just need to know **which crate will end up on top of each stack**; 
in this example, the top crates are `C` in stack 1, `M` in stack 2, and `Z` in 
stack 3, so you should combine these together and give the Elves the 
message `CMZ`.

**After the rearrangement procedure completes, what crate ends up on top of each stack?**

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- BUILT WITH -->
### Built With

[![Go][Go-shield]][Go-url]

I am using Go for this challenge as I have not being using it as much as I would like in the past year and I want to 
keep my muscle memory.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE -->
### Usage

```bash
cat example-input.txt | go run main.go
# output: Top crates: CMZ
```

<!-- MARKDOWN LINKS & IMAGES -->
[Challenge-url]: https://adventofcode.com/2022/day/5
[Go-shield]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/