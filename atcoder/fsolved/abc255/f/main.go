package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type Node struct {
	l, r int
}

func solve(n int, p, li []int) ([]Node, error) {
	errMsg := "Does not construct"

	if p[0] != 1 {
		return nil, errors.New(errMsg)
	}
	//0-indexed
	for i := 0; i < n; i++ {
		p[i]--
		li[i]--
	}

	//Iの値xがi番目(0-indexed)にあることを保持する
	mi := make([]int, n)
	for i := 0; i < n; i++ {
		mi[li[i]] = i
	}
	//fmt.Println(mi)

	ans := make([]Node, n)

	var f func(ps, pt, is, it int) (int, error)
	f = func(ps, pt, is, it int) (int, error) {
		current := p[ps]
		mid := mi[current]
		//Iの部分区間(is - it)の中に根になる候補が計算できない
		if mid < is || it < mid {
			return 0, errors.New(errMsg)
		}
		if is < mid {
			left, err := f(ps+1, ps+mid-is, is, mid-1)
			if err != nil {
				return 0, err
			}
			ans[current].l = left
		}
		if mid < it {
			right, err := f(ps+mid-is+1, pt, mid+1, it)
			if err != nil {
				return 0, err
			}
			ans[current].r = right
		}
		return current + 1, nil
	}
	_, err := f(0, n-1, 0, n-1)
	if err != nil {
		return nil, err
	}
	return ans, nil

}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	p := nextIntSlice(n)
	li := nextIntSlice(n)
	ans, err := solve(n, p, li)
	if err != nil {
		PrintInt(-1)
	} else {
		PrintVertically(ans)
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

func PrintVertically(x []Node) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v.l, v.r)
	}
}
