package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const COST_MAX = 1000000000

// This program is wrong.
// I have to consider about all slime slide when X action exec.
func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		// one line
		l := sc.Text()
		t := strings.Split(l, " ")
		n, _ := strconv.Atoi(t[0])
		x, _ := strconv.Atoi(t[1])

		// two line
		sc.Scan()
		l = sc.Text()
		t = strings.Split(l, " ")
		c := make([]int, len(t))
		for idx, u := range t {
			c[idx], _ = strconv.Atoi(u)
		}

		totalCost := 0
		for i := 0; i < n; i++ {
			cost := COST_MAX
			for l := 0; l < n; l++ {
				foo := 0
				if i == l {
					foo = c[l]
				} else {
					foo = abs(i-l)*x + c[l]
				}
				// Find min cost
				if cost > foo {
					cost = foo
				}
			}

			totalCost += cost
		}

		fmt.Println(totalCost)
	}
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
