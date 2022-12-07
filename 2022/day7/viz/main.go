package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	root := node{name: "/"}
	var dir *node

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			return fmt.Errorf("unexpected empty line")
		}

		if line == "$ ls" {
			continue
		}

		if line == "$ cd /" {
			dir = &root
			continue
		}

		if line == "$ cd .." {
			dir = dir.parent
			continue
		}

		if strings.HasPrefix(line, "$ cd ") {
			dir = dir.child(strings.TrimPrefix(line, "$ cd "))
			continue
		}

		if strings.HasPrefix(line, "dir ") {
			name := strings.TrimPrefix(line, "dir ")
			dir.addChildDir(name)
			continue
		}

		splits := strings.SplitN(line, " ", 2)

		i, err := strconv.ParseInt(splits[0], 10, 64)
		if err != nil {
			return fmt.Errorf("could not parse '%s' as number: %w", splits[0], err)
		}

		dir.addChildFile(splits[1], i)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	fmt.Fprint(stdout, render(&root, 0))

	return nil
}

type node struct {
	name     string
	parent   *node
	children []*node
	isFile   bool
	size     int64
}

func render(n *node, depth int) string {
	indent := strings.Repeat(" ", depth)

	typ := "dir"
	if n.isFile {
		typ = "file"
	}

	res := fmt.Sprintf("%s- %s (%s, size=%d)\n", indent, n.name, typ, n.size)

	for child := range n.children {
		res += render(n.children[child], depth+2)
	}

	return res
}

func (n *node) addChildDir(name string) {
	for i := range n.children {
		if n.children[i].name == name {
			panic("adding a dir twice")
		}
	}

	n.children = append(n.children, &node{
		name:   name,
		parent: n,
	})
}

func (n *node) child(name string) *node {
	for i := range n.children {
		if n.children[i].name == name {
			return n.children[i]
		}
	}

	panic("child not found")
}

func (n *node) addChildFile(name string, size int64) {
	for i := range n.children {
		if n.children[i].name == name {
			panic("adding a file twice")
		}
	}

	n.children = append(n.children, &node{
		name:   name,
		parent: n,
		isFile: true,
		size:   size,
	})

	nd := n
	for {
		if nd == nil {
			break
		}

		nd.size += size

		nd = nd.parent
	}
}
