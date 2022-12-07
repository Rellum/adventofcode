<a name="readme-top"></a>

<!-- TITLE -->
## Advent of Code 2015 Day 1

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

**--- Day 1: Not Quite Lisp ---**

Santa was hoping for a white Christmas, but his weather machine's "snow"
function is powered by stars, and he's fresh out!  To save Christmas, he
needs you to collect **fifty stars** by December 25th.

Collect stars by helping Santa solve puzzles.  Two puzzles will be made
available on each day in the Advent calendar; the second puzzle is unlocked
when you complete the first.  Each puzzle grants **one star**. Good luck!

Here's an easy puzzle to warm you up.

Santa is trying to deliver presents in a large apartment building, but he
can't find the right floor - the directions he got are a little confusing.
He starts on the ground floor (floor `0`) and then follows the instructions
one character at a time.

An opening parenthesis, `(`, means he should go up one floor, and a closing
parenthesis, `)`, means he should go down one floor.

The apartment building is very tall, and the basement is very deep; he will
never find the top or bottom floors.

For example:

* `(())` and `()()` both result in floor `0`.
* `(((` and `(()(()(` both result in floor `3`.
* `))(((((` also results in floor `3`.
* `())` and `))(` both result in floor `-1` (the first basement level).
* `)))` and `)())())` both result in floor `-3`.

To **what floor** do the instructions take Santa?

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
echo "))((((" | go run main.go
# output: Floor: 2
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
[Challenge-url]: https://adventofcode.com/2015/day/1
[Go-shield]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/