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

	q := nextInt()
	var com []string
	var x []int
	for i := 0; i < q; i++ {
		com = append(com, nextString())
		if com[i] != "DELETE" {
			x = append(x, nextInt())
		} else {
			x = append(x, 0)
		}
	}
	ans := solve(q, com, x)
	PrintHorizonaly(ans)
}

func solve(q int, com []string, x []int) []int {
	par := []int{0}
	v := []int{-1}
	cur := 0
	m := make(map[int]int)
	var ans []int
	for i := 0; i < q; i++ {
		switch com[i] {
		case "ADD":
			par = append(par, cur)
			v = append(v, x[i])
			cur = len(v) - 1
		case "SAVE":
			m[x[i]] = cur
		case "DELETE":
			cur = par[cur]
		case "LOAD":
			if value, found := m[x[i]]; found {
				cur = value
			} else {
				cur = 0
			}
		}
		ans = append(ans, v[cur])
	}
	return ans
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
