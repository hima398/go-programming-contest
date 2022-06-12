package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func divide(x int) (res []int) {
	m := make(map[int]struct{})
	for i := 1; i*i <= x; i++ {
		if x%i == 0 {
			m[i] = struct{}{}
			m[x/i] = struct{}{}
		}
	}
	for k := range m {
		res = append(res, k)
	}
	return res
}

func solver(n int) int {
	s := strconv.Itoa(n)
	d := divide(len(s))
	//fmt.Println(d)
	res := 0
	//sを、v文字が繰り返される周期的な数と考える
	for _, v := range d {
		if v > len(s)/2 {
			continue
		}
		var ss []int
		for i := 0; i+v <= len(s); i += v {
			si, _ := strconv.Atoi(s[i : i+v])
			ss = append(ss, si)
		}
		//fmt.Println(ss)
		pre := ss[0]
		c1, _ := strconv.Atoi(strings.Repeat(strconv.Itoa(pre), len(s)/v))
		if c1 <= n {
			res = Max(res, c1)
		}
		c2, _ := strconv.Atoi(strings.Repeat(strconv.Itoa(pre-1), len(s)/v))
		if c2 <= n {
			res = Max(res, c2)
		}
	}
	c3 := 1
	for c3*10 <= n {
		c3 *= 10
	}
	c3--
	if c3 <= n {
		res = Max(res, c3)
	}

	return res
}

func solve(t int, n []int) (ans []int) {
	for i := 0; i < t; i++ {
		ans = append(ans, solver(n[i]))
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := nextInt()
	var n []int
	for i := 0; i < t; i++ {
		n = append(n, nextInt())
	}

	ans := solve(t, n)

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
