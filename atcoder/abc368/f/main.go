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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)

	ok := solve(n, a)

	if ok {
		Print("Anna")
	} else {
		Print("Bruno")
	}
}

func dividePrimes(x int) int {
	//ret := make(map[int]int)
	var res int
	for x%2 == 0 {
		res++
		x /= 2
	}
	for i := 3; i*i <= x; i += 2 {
		if x == 1 {
			break
		}
		for x%i == 0 {
			res++
			x /= i
		}
	}
	if x != 1 {
		res++
	}

	return res
}

func solve(n int, a []int) bool {
	b := make([]int, n)
	for i, ai := range a {
		b[i] = dividePrimes(ai)
	}
	//fmt.Println(b)
	var x int
	for _, bi := range b {
		x ^= bi
	}

	return x != 0
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
