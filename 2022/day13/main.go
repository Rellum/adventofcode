package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	var a, b *treeNode
	var allPacketRoots []*treeNode

	var indexSum int
	var lineNum int

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		if len(line) == 0 {
			if lineNum%3 != 0 {
				return fmt.Errorf("unexpected empty line")
			}
			continue
		}

		p := parse(line)

		allPacketRoots = append(allPacketRoots, p)
		if lineNum%3 == 1 {
			a = p
		}
		if lineNum%3 == 2 {
			b = p

			if ordered(a, b) >= 0 {
				indexSum += (lineNum / 3) + 1
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	fmt.Fprintf(stdout, "Answer (part 1): %d\n", indexSum)

	divider1, divider2 := parse("[[2]]"), parse("[[6]]")
	allPacketRoots = append(allPacketRoots, divider1, divider2)
	sort.Slice(allPacketRoots, func(i, j int) bool {
		return ordered(allPacketRoots[i], allPacketRoots[j]) >= 0
	})

	dividerIndexProduct := 1
	for i := 0; i < len(allPacketRoots); i++ {
		if allPacketRoots[i] == divider1 || allPacketRoots[i] == divider2 {
			dividerIndexProduct *= (i + 1)
		}
	}

	fmt.Fprintf(stdout, "Answer (part 2): %d\n", dividerIndexProduct)

	return nil
}

type treeNode struct {
	val      *int
	parent   *treeNode
	children []*treeNode
}

func parse(s string) *treeNode {
	var node *treeNode
	var buf string

	for _, c := range s {
		if '0' <= c && c <= '9' {
			buf += string(c)
			continue
		} else if len(buf) > 0 {
			i, err := strconv.Atoi(buf)
			if err != nil {
				panic(fmt.Errorf("read rune: %w", err))
			}

			node.children = append(node.children, &treeNode{parent: node, val: &i})
			buf = ""
		}

		switch c {
		case '[':
			if node == nil {
				node = &treeNode{}
				continue
			}

			child := treeNode{parent: node}
			node.children = append(node.children, &child)
			node = &child
			continue
		case ']':
			if node.parent == nil {
				return node

			}

			node = node.parent
			continue
		}
	}

	return nil
}

func ordered(a, b *treeNode) int {
	if a.val != nil && b.val != nil {
		return compare(*a.val, *b.val)
	}

	if a.val == nil && b.val == nil {
		for i := 0; i < len(a.children); i++ {
			if i == len(b.children) {
				return -1
			}

			if o := ordered(a.children[i], b.children[i]); o != 0 {
				return o
			}
		}

		return compare(len(a.children), len(b.children))
	}

	if a.val == nil {
		return ordered(a, &treeNode{children: []*treeNode{b}})
	}

	return ordered(&treeNode{children: []*treeNode{a}}, b)
}

func compare(a, b int) int {
	if a < b {
		return 1
	}
	if b < a {
		return -1
	}
	return 0
}
