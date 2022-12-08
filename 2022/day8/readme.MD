<a name="readme-top"></a>

<!-- TITLE -->
## Advent of Code 2022 Day 8

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

**--- Day 8: Treetop Tree House ---**

The expedition comes across a peculiar patch of tall trees all planted carefully in a grid. The Elves explain that a previous expedition planted these trees as a reforestation effort. Now, they're curious if this would be a good location for a [tree house][descLink1].

First, determine whether there is enough tree cover here to keep a tree house **hidden**. To do this, you need to count the number of trees that are **visible from outside the grid** when looking directly along a row or column.

The Elves have already launched a [quadcopter][descLink2] to generate a map with the height of each tree (your puzzle input). For example:

```
30373
25512
65332
33549
35390
```

Each tree is represented as a single digit whose value is its height, where `0` is the shortest and `9` is the tallest.

A tree is **visible** if all of the other trees between it and an edge of the grid are **shorter** than it. Only consider trees in the same row or column; that is, only look up, down, left, or right from any given tree.

All of the trees around the edge of the grid are **visible** - since they are already on the edge, there are no trees to block the view. In this example, that only leaves the **interior nine trees** to consider:

* The top-left `5` is **visible** from the left and top. (It isn't visible from the right or bottom since other trees of height `5` are in the way.)
* The top-middle `5` is **visible** from the top and right.
* The top-right `1` is not visible from any direction; for it to be visible, there would need to only be trees of height **0** between it and an edge.
* The left-middle `5` is **visible**, but only from the right.
* The center `3` is not visible from any direction; for it to be visible, there would need to be only trees of at most height `2` between it and an edge.
* The right-middle `3` is **visible** from the right.
* In the bottom row, the middle `5` is **visible**, but the `3` and `4` are not.

With 16 trees visible on the edge and another 5 visible in the interior, a total of `21` trees are visible in this arrangement.

Consider your map; **how many trees are visible from outside the grid?**

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- BUILT WITH -->
## Built With

[![Go][Go-shield]][Go-url]

I am using Go for this challenge as I have not being using it as much as I would like in the past year and I want to 
keep my muscle memory.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE -->
## Usage

```bash
cat example-input.txt | go run main.go
# output:
# vvvvv
# vvv.v
# vv.vv
# v.v.v
# vvvvv
# Answer: 21

```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
[Challenge-url]: https://adventofcode.com/2022/day/8
[descLink1]: https://en.wikipedia.org/wiki/Tree_house
[descLink2]: https://en.wikipedia.org/wiki/Quadcopter
[Go-shield]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/