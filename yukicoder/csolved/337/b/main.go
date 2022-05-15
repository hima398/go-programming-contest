package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(n int, a []int) string {
	sort.Ints(a)
	d := make(map[int]struct{})
	var ds []int
	for i := 1; i < n; i++ {
		if a[i] > 0 && a[i-1] > 0 {
			ds = append(ds, a[i]-a[i-1])
			d[a[i]-a[i-1]] = struct{}{}
		}
	}
	//数字が書かれているカードがすでに等差数列になっている
	if len(d) <= 1 {
		return "Yes"
	}
	//等差が0と異なる等差の値がある場合は数列Bを作成できない
	for k := range d {
		if k == 0 {
			return "No"
		}
	}

	var gcd int
	for k := range d {
		if gcd == 0 {
			gcd = k
		} else {
			gcd = Gcd(gcd, k)
		}
	}
	//fmt.Println(ds)
	var nz int
	for _, v := range a {
		if v == 0 {
			nz++
		}
	}
	var need int
	for i := 1; i < len(ds); i++ {
		if gcd > 0 {
			need += ds[i]/gcd - 1
		}
	}
	if nz >= need {
		return "Yes"
	} else {
		return "No"
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	ans := Solve(n, a)
	fmt.Println(ans)
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

func Gcd(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	/*
		if x < y {
			x, y = y, x
		}
	*/
	return Gcd(y, x%y)
}

func Lcm(x, y int) int {
	// x*yのオーバーフロー対策のため先にGcdで割る
	// Gcd(x, y)はxの約数のため割り切れる
	ret := x / Gcd(x, y)
	ret *= y
	return ret
}
