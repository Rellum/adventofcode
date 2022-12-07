<a name="readme-top"></a>

<!-- TITLE -->
## Advent of Code 2015 Day 5

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

**--- Day 5: Doesn't He Have Intern-Elves For This? ---**

Santa needs help figuring out which strings in his text file are naughty or nice.

A **nice string** is one with all of the following properties:

* It contains at least three vowels (`aeiou` only), like `aei`, `xazegov`, or `aeiouaeiouaeiou`.
* It contains at least one letter that appears twice in a row, like `xx`, `abcdde` (`dd`), or `aabbccdd` (`aa`, `bb`, `cc`, or `dd`).
* It does **not** contain the strings `ab`, `cd`, `pq`, or `xy`, even if they are part of one of the other requirements.

For example:

* `ugknbfddgicrmopn` is nice because it has at least three vowels (`u...i...o...`), a double letter (`...dd...`), and none of the disallowed substrings.
* `aaa` is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
* `jchzalrnumimnmhp` is naughty because it has no double letter.
* `haegwjzuvuyypxyu` is naughty because it contains the string `xy`.
* `dvszwmarrgswjxmb` is naughty because it contains only one vowel.

How many strings are nice?

**--- Part Two ---**

Realizing the error of his ways, Santa has switched to a better model of determining whether a string is naughty or nice.  None of the old rules apply, as they are all clearly ridiculous.

Now, a nice string is one with all of the following properties:

* It contains a pair of any two letters that appears at least twice in the string without overlapping, like `xyxy` (`xy`) or `aabcdefgaa` (`aa`), but not like `aaa` (`aa`, but it overlaps).
* It contains at least one letter which repeats with exactly one letter between them, like `xyx`, `abcdefeghi` (`efe`), or even `aaa`.

For example:

* `qjhvhtzxzqqjkmpb` is nice because is has a pair that appears twice (`qj`) and a letter that repeats with exactly one letter between them (`zxz`).
* `xxyxx` is nice because it has a pair that appears twice and a letter that repeats with one between, even though the letters used by each rule overlap.
* `uurcxstgmygtbstg` is naughty because it has a pair (`tg`) but no repeat with a single letter between them.
* `ieodomkazucvgmuy` is naughty because it has a repeating letter with one between (`odo`), but no pair that appears twice.

How many strings are nice under these new rules?

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
# output:  Nice count (part 1): 2
#          Nice count (part 2): 2
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
[Challenge-url]: https://adventofcode.com/2015/day/5
[Go-shield]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/