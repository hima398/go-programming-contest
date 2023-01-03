package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	var s []string
	for i := 0; i < 9; i++ {
		s = append(s, nextString())
	}
	ans := solve(s)
	PrintInt(ans)
}

func solve(s []string) int {
	type point struct {
		r, c int
	}
	var ps []point
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s[i][j] == '#' {
				ps = append(ps, point{i, j})
			}
		}
	}
	n := len(ps)
	var ans int
	computeDist2 := func(r1, c1, r2, c2 int) int {
		rr := (r2 - r1)
		cc := (c2 - c1)
		return rr*rr + cc*cc
	}
	isSqare := func(r1, c1, r2, c2, r3, c3, r4, c4 int) bool {
		m := make(map[int]int)
		m[computeDist2(r1, c1, r2, c2)]++
		m[computeDist2(r1, c1, r3, c3)]++
		m[computeDist2(r1, c1, r4, c4)]++
		m[computeDist2(r2, c2, r3, c3)]++
		m[computeDist2(r2, c2, r4, c4)]++
		m[computeDist2(r3, c3, r4, c4)]++
		if len(m) != 2 {
			return false
		}
		var ks []int
		for k := range m {
			ks = append(ks, k)
		}
		return (m[ks[0]] == 2 && m[ks[1]] == 4) || (m[ks[0]] == 4 && m[ks[1]] == 2)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				for l := k + 1; l < n; l++ {
					if isSqare(ps[i].r, ps[i].c, ps[j].r, ps[j].c, ps[k].r, ps[k].c, ps[l].r, ps[l].c) {
						ans++
					}
				}
			}
		}
	}
	return ans
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
