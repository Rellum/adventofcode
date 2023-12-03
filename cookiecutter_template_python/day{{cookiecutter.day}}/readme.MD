<a name="readme-top"></a>

<!-- TITLE -->
## Advent of Code {{cookiecutter.year}} Day {{cookiecutter.day}}

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

paste {{cookiecutter.year}} day {{cookiecutter.day}} description here

```bash
curl -v --cookie "session=$ADVENT_OF_CODE_SESSION" https://adventofcode.com/{{cookiecutter.year}}/day/{{cookiecutter.day}}/input > challenge-input.txt
elf --session=$ADVENT_OF_CODE_SESSION --year={{cookiecutter.year}} --day={{cookiecutter.day}} | pbcopy
```

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
# output: Answer
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
[Challenge-url]: https://adventofcode.com/{{cookiecutter.year}}/day/{{cookiecutter.day}}
[Go-shield]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/
[Python-shield]: https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=white
[Python-url]: https://www.python.org/