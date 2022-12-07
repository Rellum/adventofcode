<a name="readme-top"></a>

<!-- TITLE -->
## Advent of Code 2015 Day 4

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

**--- Day 4: The Ideal Stocking Stuffer ---**

Santa needs help [mining][descLink1] some AdventCoins (very similar to [bitcoins][descLink2]) to use 
as gifts for all the economically forward-thinking little girls and boys.

To do this, he needs to find [MD5][descLink3] hashes which, in [hexadecimal][descLink4], start with 
at least **five zeroes**.  The input to the MD5 hash is some secret key (your 
puzzle input, given below) followed by a number in decimal. To mine 
AdventCoins, you must find Santa the lowest positive number (no leading 
zeroes: `1`, `2`, `3`, ...) that produces such a hash.

For example:

* If your secret key is `abcdef`, the answer is `609043`, because the MD5 
  hash of `abcdef609043` starts with five zeroes (`000001dbbfa...`), and it 
  is the lowest such number to do so.
* If your secret key is `pqrstuv`, the lowest number it combines with to 
  make an MD5 hash starting with five zeroes is `1048970`; that is, the 
  MD5 hash of `pqrstuv1048970` looks like `000006136ef...`.

**--- Part Two ---**

Now find one that starts with **six zeroes**.

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
echo "abcdef" | go run main.go
# output: Answer (5 zero prefix): 609043
#         Answer (6 zero prefix): 6742839

```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
[Challenge-url]: https://adventofcode.com/2015/day/4
[descLink1]: https://en.wikipedia.org/wiki/Bitcoin#Mining
[descLink2]: https://en.wikipedia.org/wiki/Bitcoin
[descLink3]: https://en.wikipedia.org/wiki/MD5
[descLink4]: https://en.wikipedia.org/wiki/Hexadecimal
[Go-shield]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/