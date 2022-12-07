<a name="readme-top"></a>

<!-- TITLE -->
## Advent of Code 2015 Day 2

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

**--- Day 2: I Was Told There Would Be No Math ---**

The elves are running low on wrapping paper, and so they need to submit an 
order for more.  They have a list of the dimensions (length `l`, width `w`, and 
height `h`) of each present, and only want to order exactly as much as they 
need.

Fortunately, every present is a box (a perfect right [rectangular prism][Wiki-cuboid]), 
which makes calculating the required wrapping paper for each gift a little 
easier: find the surface area of the box, which is `2*l*w + 2*w*h + 2*h*l`.  
The elves also need a little extra paper for each present: the area of the 
smallest side.

For example:

* A present with dimensions `2x3x4` requires `2*6 + 2*12 + 2*8 = 52` square 
  feet of wrapping paper plus `6` square feet of slack, for a total of `58` 
  square feet.
* A present with dimensions `1x1x10` requires `2*1 + 2*10 + 2*10 = 42` 
  square feet of wrapping paper plus `1` square foot of slack, for a total 
  of `43` square feet.

All numbers in the elves' list are in feet.  How many total **square feet of 
wrapping paper** should they order?

**--- Part Two ---**

The elves are also running low on ribbon.  Ribbon is all the same width, so 
they only have to worry about the length they need to order, which they 
would again like to be exact.

The ribbon required to wrap a present is the shortest distance around its 
sides, or the smallest perimeter of any one face.  Each present also 
requires a bow made out of ribbon as well; the feet of ribbon required for 
the perfect bow is equal to the cubic feet of volume of the present.  Don't 
ask how they tie the bow, though; they'll never tell.

For example:

* A present with dimensions `2x3x4` requires `2+2+3+3 = 10` feet of ribbon 
  to wrap the present plus `2*3*4 = 24` feet of ribbon for the bow, for a 
  total of `34` feet.
* A present with dimensions `1x1x10` requires `1+1+1+1 = 4` feet of ribbon 
  to wrap the present plus `1*1*10 = 10` feet of ribbon for the bow, for a 
  total of `14` feet.

How many total **feet of ribbon** should they order?

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
echo "2x3x4" | go run main.go
# output: 58 square feet of wrapping paper
#         34 feet of ribbon

echo "1x1x10" | go run main.go
# output: 43 square feet of wrapping paper
#         14 feet of ribbon


cat example-input.txt | go run main.go
# output: 101 square feet of wrapping paper
#         48 feet of ribbon
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
[Challenge-url]: https://adventofcode.com/2015/day/2
[Go-shield]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/
[Wiki-cuboid]: https://en.wikipedia.org/wiki/Cuboid#Rectangular_cuboid