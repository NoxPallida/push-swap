package main

import (
	"fmt"
	"math"
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

func minIdx(a []int) int {
	idx := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[idx] {
			idx = i
		}
	}
	return idx
}

func maxIdx(a []int) int {
	idx := 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[idx] {
			idx = i
		}
	}
	return idx
}

func maxVal(a []int) int {
	m := a[0]
	for _, v := range a[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func minVal(a []int) int {
	m := a[0]
	for _, v := range a[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sa(a []int) {
	if len(a) >= 2 {
		a[0], a[1] = a[1], a[0]
	}
	fmt.Println("sa")
}

func sb(b []int) {
	if len(b) >= 2 {
		b[0], b[1] = b[1], b[0]
	}
	fmt.Println("sb")
}

func ss(a, b []int) {
	if len(a) >= 2 {
		a[0], a[1] = a[1], a[0]
	}
	if len(b) >= 2 {
		b[0], b[1] = b[1], b[0]
	}
	fmt.Println("ss")
}

func pa(a, b *[]int) {
	if len(*b) > 0 {
		val := (*b)[0]
		*b = (*b)[1:]
		*a = append([]int{val}, *a...)
	}
	fmt.Println("pa")
}

func pb(a, b *[]int) {
	if len(*a) > 0 {
		val := (*a)[0]
		*a = (*a)[1:]
		*b = append([]int{val}, *b...)
	}
	fmt.Println("pb")
}

func ra(a []int) {
	if len(a) >= 2 {
		first := a[0]
		copy(a, a[1:])
		a[len(a)-1] = first
	}
	fmt.Println("ra")
}

func rb(b []int) {
	if len(b) >= 2 {
		first := b[0]
		copy(b, b[1:])
		b[len(b)-1] = first
	}
	fmt.Println("rb")
}

func rr(a, b []int) {
	if len(a) >= 2 {
		first := a[0]
		copy(a, a[1:])
		a[len(a)-1] = first
	}
	if len(b) >= 2 {
		first := b[0]
		copy(b, b[1:])
		b[len(b)-1] = first
	}
	fmt.Println("rr")
}

func rra(a []int) {
	if len(a) >= 2 {
		last := a[len(a)-1]
		copy(a[1:], a[:len(a)-1])
		a[0] = last
	}
	fmt.Println("rra")
}

func rrb(b []int) {
	if len(b) >= 2 {
		last := b[len(b)-1]
		copy(b[1:], b[:len(b)-1])
		b[0] = last
	}
	fmt.Println("rrb")
}

func rrr(a, b []int) {
	if len(a) >= 2 {
		last := a[len(a)-1]
		copy(a[1:], a[:len(a)-1])
		a[0] = last
	}
	if len(b) >= 2 {
		last := b[len(b)-1]
		copy(b[1:], b[:len(b)-1])
		b[0] = last
	}
	fmt.Println("rrr")
}

func sortThree(a []int) {
	mIdx := maxIdx(a)
	if mIdx == 0 {
		ra(a)
	} else if mIdx == 1 {
		rra(a)
	}
	if a[0] > a[1] {
		sa(a)
	}
}

func sortTurk(a, b *[]int) {
	pb(a, b)
	if len(*a) > 3 {
		pb(a, b)
	}

	for len(*a) > 3 {
		bestCost := math.MaxInt
		bestI, bestJ := 0, 0
		bestDirA, bestDirB := 1, 1

		lenA := len(*a)
		lenB := len(*b)
		maxB := maxVal(*b)
		minB := minVal(*b)

		for i, valA := range *a {
			targetJ := 0
			if valA > maxB || valA < minB {
				targetJ = maxIdx(*b)
			} else {
				bestVal := math.MinInt
				for j, valB := range *b {
					if valB < valA && valB > bestVal {
						bestVal = valB
						targetJ = j
					}
				}
			}

			cost1 := maxInt(i, targetJ)
			cost2 := maxInt(lenA-i, lenB-targetJ)
			cost3 := i + (lenB - targetJ)
			cost4 := (lenA - i) + targetJ

			minC := cost1
			dirA, dirB := 1, 1
			if cost2 < minC {
				minC = cost2
				dirA, dirB = -1, -1
			}
			if cost3 < minC {
				minC = cost3
				dirA, dirB = 1, -1
			}
			if cost4 < minC {
				minC = cost4
				dirA, dirB = -1, 1
			}

			if minC < bestCost {
				bestCost = minC
				bestI = i
				bestJ = targetJ
				bestDirA = dirA
				bestDirB = dirB
			}
		}

		i, j := bestI, bestJ
		lenA = len(*a)
		lenB = len(*b)

		if bestDirA == 1 && bestDirB == 1 {
			common := minInt(i, j)
			for k := 0; k < common; k++ {
				rr(*a, *b)
			}
			for k := 0; k < i-common; k++ {
				ra(*a)
			}
			for k := 0; k < j-common; k++ {
				rb(*b)
			}
		} else if bestDirA == -1 && bestDirB == -1 {
			rotA := lenA - i
			rotB := lenB - j
			common := minInt(rotA, rotB)
			for k := 0; k < common; k++ {
				rrr(*a, *b)
			}
			for k := 0; k < rotA-common; k++ {
				rra(*a)
			}
			for k := 0; k < rotB-common; k++ {
				rrb(*b)
			}
		} else if bestDirA == 1 && bestDirB == -1 {
			for k := 0; k < i; k++ {
				ra(*a)
			}
			for k := 0; k < lenB-j; k++ {
				rrb(*b)
			}
		} else {
			for k := 0; k < lenA-i; k++ {
				rra(*a)
			}
			for k := 0; k < j; k++ {
				rb(*b)
			}
		}

		pb(a, b)
	}

	sortThree(*a)

	for len(*b) > 0 {
		valB := (*b)[0]
		lenA := len(*a)
		maxA := maxVal(*a)
		minA := minVal(*a)

		targetI := 0
		if valB > maxA || valB < minA {
			targetI = minIdx(*a)
		} else {
			bestVal := math.MaxInt
			for i, valA := range *a {
				if valA > valB && valA < bestVal {
					bestVal = valA
					targetI = i
				}
			}
		}

		if targetI <= lenA/2 {
			for k := 0; k < targetI; k++ {
				ra(*a)
			}
		} else {
			for k := 0; k < lenA-targetI; k++ {
				rra(*a)
			}
		}
		pa(a, b)
	}

	mIdx := minIdx(*a)
	lenA := len(*a)
	if mIdx <= lenA/2 {
		for k := 0; k < mIdx; k++ {
			ra(*a)
		}
	} else {
		for k := 0; k < lenA-mIdx; k++ {
			rra(*a)
		}
	}
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

	if len(a) <= 1 || isSorted(a) {
		return
	}

	var b []int

	if len(a) == 2 {
		sa(a)
	} else if len(a) == 3 {
		sortThree(a)
	} else {
		sortTurk(&a, &b)
	}
}
