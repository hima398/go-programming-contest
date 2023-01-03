package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	s := nextString()
	ans := solve(n, x, y, s)
	PrintString(ans)
}

func solve(n int, x, y []int, s string) string {
	//m := make(map[int][]int)
	leftMax, rightMin := make(map[int]int), make(map[int]int)
	for i := 0; i < n; i++ {
		if s[i] == 'L' {
			if v, found := rightMin[y[i]]; found {
				if v < x[i] {
					return "Yes"
				}
			}
		} else {
			if v, found := leftMax[y[i]]; found {
				if x[i] < v {
					return "Yes"
				}
			}
		}
		if s[i] == 'L' {
			if _, found := leftMax[y[i]]; found {
				leftMax[y[i]] = Max(leftMax[y[i]], x[i])
			} else {
				leftMax[y[i]] = x[i]
			}
		} else {
			if _, found := rightMin[y[i]]; found {
				rightMin[y[i]] = Min(rightMin[y[i]], x[i])
			} else {
				rightMin[y[i]] = x[i]
			}
		}
	}
	return "No"
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
