package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const INPUT_MAX = 1000000000

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		l := sc.Text()
		h := strings.Split(l, " ")
		o := [2]int{INPUT_MAX, INPUT_MAX}

		// Only check all odd number
		for _, tmp := range h {
			n, _ := strconv.Atoi(tmp)
			if n%2 == 0 {
				o = [2]int{}
				break
			}

			if len(o) < 2 {
				o[len(o)-1] = n
			} else {
				if o[0] < o[1] {
					if o[1] > n {
						o[1] = n
					}
				} else {
					if o[0] > n {
						o[0] = n
					}
				}
			}
		}

		if len(o) == 2 {
			fmt.Println(o[0] * o[1])
		} else {
			fmt.Println(0)
		}
	}
}
