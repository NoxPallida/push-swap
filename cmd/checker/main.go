package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(args []string) ([]int, bool) {
	if len(args) == 0 {
		return nil, true
	}
	var nums []int
	seen := make(map[int]bool)
	for _, arg := range args {
		fields := strings.Fields(arg)
		if len(fields) == 0 {
			return nil, false
		}
		for _, f := range fields {
			val, err := strconv.ParseInt(f, 10, 32)
			if err != nil {
				return nil, false
			}
			n := int(val)
			if seen[n] {
				return nil, false
			}
			seen[n] = true
			nums = append(nums, n)
		}
	}
	return nums, true
}

func isSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func execOp(op string, a, b *[]int) bool {
	switch op {
	case "sa":
		if len(*a) >= 2 {
			(*a)[0], (*a)[1] = (*a)[1], (*a)[0]
		}
	case "sb":
		if len(*b) >= 2 {
			(*b)[0], (*b)[1] = (*b)[1], (*b)[0]
		}
	case "ss":
		if len(*a) >= 2 {
			(*a)[0], (*a)[1] = (*a)[1], (*a)[0]
		}
		if len(*b) >= 2 {
			(*b)[0], (*b)[1] = (*b)[1], (*b)[0]
		}
	case "pa":
		if len(*b) > 0 {
			val := (*b)[0]
			*b = (*b)[1:]
			*a = append([]int{val}, *a...)
		}
	case "pb":
		if len(*a) > 0 {
			val := (*a)[0]
			*a = (*a)[1:]
			*b = append([]int{val}, *b...)
		}
	case "ra":
		if len(*a) >= 2 {
			first := (*a)[0]
			copy(*a, (*a)[1:])
			(*a)[len(*a)-1] = first
		}
	case "rb":
		if len(*b) >= 2 {
			first := (*b)[0]
			copy(*b, (*b)[1:])
			(*b)[len(*b)-1] = first
		}
	case "rr":
		if len(*a) >= 2 {
			first := (*a)[0]
			copy(*a, (*a)[1:])
			(*a)[len(*a)-1] = first
		}
		if len(*b) >= 2 {
			first := (*b)[0]
			copy(*b, (*b)[1:])
			(*b)[len(*b)-1] = first
		}
	case "rra":
		if len(*a) >= 2 {
			last := (*a)[len(*a)-1]
			copy((*a)[1:], (*a)[:len(*a)-1])
			(*a)[0] = last
		}
	case "rrb":
		if len(*b) >= 2 {
			last := (*b)[len(*b)-1]
			copy((*b)[1:], (*b)[:len(*b)-1])
			(*b)[0] = last
		}
	case "rrr":
		if len(*a) >= 2 {
			last := (*a)[len(*a)-1]
			copy((*a)[1:], (*a)[:len(*a)-1])
			(*a)[0] = last
		}
		if len(*b) >= 2 {
			last := (*b)[len(*b)-1]
			copy((*b)[1:], (*b)[:len(*b)-1])
			(*b)[0] = last
		}
	default:
		return false
	}
	return true
}

func main() {
	if len(os.Args) == 1 {
		return
	}

	a, ok := parseInput(os.Args[1:])
	if !ok {
		os.Stderr.WriteString("Error\n")
		os.Exit(1)
	}

	var b []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if !execOp(line, &a, &b) {
			os.Stderr.WriteString("Error\n")
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		os.Stderr.WriteString("Error\n")
		os.Exit(1)
	}

	if isSorted(a) && len(b) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
