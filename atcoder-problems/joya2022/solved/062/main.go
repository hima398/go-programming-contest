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
	a := nextIntSlice(n)
	ans := solve(n, a)
	PrintInt(ans)
}

func solve(n int, a []int) int {
	gcd := a[0]
	for i := 1; i < n; i++ {
		gcd = Gcd(gcd, a[i])
	}
	for i := range a {
		a[i] /= gcd
	}
	var ans int
	for i := 0; i < n; i++ {
		for a[i]%3 == 0 {
			a[i] /= 3
			ans++
		}
		for a[i]%2 == 0 {
			a[i] /= 2
			ans++
		}
	}
	//fmt.Println(gcd)
	//fmt.Println(a)
	ok := true
	for i := 0; i < n; i++ {
		ok = ok && a[i] == 1
	}
	if ok {
		return ans
	} else {
		return -1
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
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
