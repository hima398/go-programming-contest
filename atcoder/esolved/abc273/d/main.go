package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type Coordinate struct {
	r, c int
}

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w, sr, sc := nextInt(), nextInt(), nextInt(), nextInt()
	n := nextInt()
	var r, c []int
	for i := 0; i < n; i++ {
		r = append(r, nextInt())
		c = append(c, nextInt())
	}
	q := nextInt()
	var d []string
	var l []int
	for i := 0; i < q; i++ {
		d = append(d, nextString())
		l = append(l, nextInt())
	}
	ans := solve(h, w, sr, sc, n, r, c, q, d, l)
	PrintVertically(ans)

}

func solve(h, w, sr, sc, n int, r, c []int, q int, d []string, l []int) []Coordinate {
	const INF = 1 << 60
	rp, rm := make(map[int][]int), make(map[int][]int)
	cp, cm := make(map[int][]int), make(map[int][]int)

	//壁を登録
	for i := 0; i < n; i++ {
		rp[c[i]] = append(rp[c[i]], r[i])
		rm[c[i]] = append(rm[c[i]], -r[i])

		cp[r[i]] = append(cp[r[i]], c[i])
		cm[r[i]] = append(cm[r[i]], -c[i])
	}
	for k := range rm {
		rm[k] = append(rm[k], 0)
		sort.Ints(rm[k])
	}
	for k := range rp {
		rp[k] = append(rp[k], h+1)
		sort.Ints(rp[k])
	}
	for k := range cm {
		cm[k] = append(cm[k], 0)
		sort.Ints(cm[k])
	}
	for k := range cp {
		cp[k] = append(cp[k], w+1)
		sort.Ints(cp[k])
	}

	var ans []Coordinate
	curR, curC := sr, sc
	for k := 0; k < q; k++ {
		//fmt.Println(k, d[k], l[k])
		switch d[k] {
		case "U":
			idx := sort.Search(len(rm[curC]), func(i int) bool {
				return -curR < rm[curC][i]
			})
			if idx >= len(rm[curC]) {
				curR = Max(curR-l[k], 1)
			} else {
				curR = Max(curR-l[k], -rm[curC][idx]+1)
			}
		case "D":
			idx := sort.Search(len(rp[curC]), func(i int) bool {
				return curR < rp[curC][i]
			})
			if idx >= len(rp[curC]) {
				curR = Min(curR+l[k], h)
			} else {
				curR = Min(curR+l[k], rp[curC][idx]-1)
			}
		case "L":
			idx := sort.Search(len(cm[curR]), func(i int) bool {
				return -curC < cm[curR][i]
			})
			if idx >= len(cm[curR]) {
				curC = Max(curC-l[k], 1)
			} else {
				curC = Max(curC-l[k], -cm[curR][idx]+1)
			}
		case "R":
			idx := sort.Search(len(cp[curR]), func(i int) bool {
				return curC < cp[curR][i]
			})
			if idx >= len(cp[curR]) {
				curC = Min(curC+l[k], w)
			} else {
				curC = Min(curC+l[k], cp[curR][idx]-1)
			}
		}
		ans = append(ans, Coordinate{curR, curC})
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintVertically(x []Coordinate) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v.r, v.c)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
