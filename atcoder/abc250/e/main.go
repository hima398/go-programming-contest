package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func solveHonestly(n, q int, a, b, x, y []int) []string {
	var ans []string
	for k := 0; k < q; k++ {
		mx, my := make(map[int]struct{}), make(map[int]struct{})
		for i := 0; i < x[k]; i++ {
			mx[a[i]] = struct{}{}
		}
		for i := 0; i < y[k]; i++ {
			my[b[i]] = struct{}{}
		}
		ok := true
		for key := range mx {
			_, has := my[key]
			ok = ok && has
		}
		for key := range my {
			_, has := mx[key]
			ok = ok && has
		}
		if ok {
			ans = append(ans, "Yes")
		} else {
			ans = append(ans, "No")
		}
	}
	return ans
}

func firstsolve(n, q int, a, b, x, y []int) []string {
	ans := make([]string, q)

	type query struct {
		i, x, y int
	}
	var qs []query
	for k := 0; k < q; k++ {
		qs = append(qs, query{k, x[k], y[k]})
	}
	sort.Slice(qs, func(i, j int) bool {
		if qs[i].x == qs[j].x {
			return qs[i].y < qs[j].y
		}
		return qs[i].x < qs[j].x
	})
	var xi, yi int
	mx1, mx2 := make(map[int]int), make(map[int]int)
	my1, my2 := make(map[int]int), make(map[int]int)
	for _, v := range qs {
		i := xi
		for ; i < v.x; i++ {
			mx2[a[i]]++
			if _, found := my2[a[i]]; found {
				my2[a[i]]--
				if my2[a[i]] == 0 {
					delete(my2, a[i])
				}
			}
		}
		xi = i
		i = yi
		for ; i < Min(xi, v.y); i++ {
			if _, found := mx2[b[i]]; found {
				mx1[b[i]]++
				mx2[b[i]]--
				if mx2[b[i]] == 0 {
					delete(mx2, b[i])
				}
			}
		}
		yi = i
		if len(mx2) == 0 && len(my2) == 0 {
			ans[v.i] = "Yes"
		} else {
			ans[v.i] = "No"
		}
	}

	sort.Slice(qs, func(i, j int) bool {
		if qs[i].y == qs[j].y {
			return qs[i].x < qs[j].x
		}
		return qs[i].y < qs[j].y
	})
	xi, yi = 0, 0
	mx1, mx2 = make(map[int]int), make(map[int]int)
	my1, my2 = make(map[int]int), make(map[int]int)
	for _, v := range qs {
		i := yi
		for ; i < v.y; i++ {
			if _, found := mx2[b[i]]; found {
				mx2[b[i]]--
				if mx2[b[i]] == 0 {
					delete(mx2, b[i])
				}
				my1[b[i]]++
			} else {
				my2[b[i]]++
			}
		}
		yi = i
		i = xi
		for {
			if _, found := my2[a[i]]; found {
				my1[a[i]]++
				my2[a[i]]--
				if my2[a[i]] == 0 {
					delete(my2, a[i])
				}
			} else {
				mx2[a[i]]++
			}
			i++
			fmt.Println(i, v.x, yi)
			if i >= Min(v.x, yi) {
				break
			}
		}
		xi = i
		if len(ans[v.i]) > 0 {
			continue
		} else if len(mx2) == 0 && len(my2) == 0 {
			ans[v.i] = "Yes"
		} else {
			ans[v.i] = "No"
		}
	}

	return ans
}

func solveCommentary(n, q int, a, b, x, y []int) (ans []string) {
	const INF = 1 << 60
	//0-indexed
	for i := 0; i < q; i++ {
		x[i]--
		y[i]--
	}
	va := make([]int, n)
	ma := make(map[int]int)
	for i, ai := range a {
		if _, found := ma[ai]; !found {
			ma[ai] = len(ma) + 1
		}
		va[i] = len(ma)
	}
	mx := 0
	mb := make(map[int]struct{})
	ms := make([]int, n)
	vb := make([]int, n)
	for i, bi := range b {
		mb[bi] = struct{}{}
		if v, found := ma[bi]; found {
			mx = Max(mx, v)
		} else {
			mx = Max(mx, INF)
		}
		ms[i] = mx
		vb[i] = len(mb)
	}
	//fmt.Println(va)
	//fmt.Println(vb)
	//fmt.Println(ms)
	for i := 0; i < q; i++ {
		if va[x[i]] == vb[y[i]] && ms[y[i]] == vb[y[i]] {
			ans = append(ans, "Yes")
		} else {
			ans = append(ans, "No")
		}
	}
	return ans

}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	b := nextIntSlice(n)
	q := nextInt()
	var x, y []int
	for i := 0; i < q; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	//ans := solveHonestly(n, q, a, b, x, y)
	ans := solveCommentary(n, q, a, b, x, y)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
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
