package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type Node struct {
	i, nx, sum int
	s          string
}

func computeScore(s string) int {
	var res int
	m := make([]int, 10)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'X' {
			for j := 0; j < 10; j++ {
				res += m[j] * j
			}
		} else {
			m[s[i]-'0']++
		}
	}
	return res
}
func buildNode(idx int, s string) Node {
	var nx, sum int
	for _, r := range s {
		if r == 'X' {
			nx++
		} else {
			sum += int(r - '0')
		}
	}
	return Node{idx, nx, sum, s}
}

func solve(n int, s []string) int {
	const INF = 1 << 60
	var nodes []Node
	for i, si := range s {
		nodes = append(nodes, buildNode(i, si))
	}
	sort.Slice(nodes, func(i, j int) bool {
		var x1, x2 int
		if nodes[i].nx == 0 {
			x1 = INF
		} else {
			x1 = nodes[i].sum * nodes[j].nx
		}
		if nodes[j].nx == 0 {
			x2 = INF
		} else {
			x2 = nodes[j].sum * nodes[i].nx
		}
		return x1 < x2
	})
	var t []string
	for _, nd := range nodes {
		t = append(t, nd.s)
	}
	ans := computeScore(strings.Join(t, ""))
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}
	ans := solve(n, s)
	PrintInt(ans)
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
