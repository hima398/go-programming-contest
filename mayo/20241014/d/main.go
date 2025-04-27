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

	ans := solve(n, a)

	switch ans {
	case 0:
		Print("pairwise coprime")
	case 1:
		Print("setwise coprime")
	case 2:
		Print("not coprime")
	}
}

func dividePrimes(x int) map[int]struct{} {
	res := make(map[int]struct{})
	for x%2 == 0 {
		res[2] = struct{}{}
		x /= 2
	}
	for i := 3; i*i <= x; i += 2 {
		for x%i == 0 {
			res[i] = struct{}{}
			x /= i
		}
	}
	if x > 1 {
		res[x] = struct{}{}
	}
	return res
}

func solve(n int, a []int) int {
	gcd := a[0]
	for i := 1; i < n; i++ {
		gcd = Gcd(gcd, a[i])
	}
	if gcd > 1 {
		return 2
	}

	primes := make(map[int]struct{})
	for _, ai := range a {
		ps := dividePrimes(ai)
		//fmt.Println(ps)
		for k := range ps {
			if _, found := primes[k]; found {
				return 1
			}
			primes[k] = struct{}{}
		}
	}

	return 0
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
