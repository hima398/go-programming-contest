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

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	a := nextIntSlice(n)

	ans := solve(n, m, a)

	Print(len(ans))
	for _, v := range ans {
		Print(v)
	}
}

// 数列aに含まれる素数を列挙する
func dividePrimes(a []int) map[int]int {
	ret := make(map[int]int)
	for _, ai := range a {
		for ai%2 == 0 {
			ret[2]++
			ai /= 2
		}
		for i := 3; i*i <= ai; i += 2 {
			if ai == 1 {
				break
			}
			for ai%i == 0 {
				ret[i]++
				ai /= i
			}
		}
		if ai != 1 {
			ret[ai]++
		}
	}
	return ret
}
func solve(n, m int, a []int) []int {
	e := make([]bool, m+1)
	e[0] = true

	primes := dividePrimes(a)
	sort.Ints(a)
	for p := range primes {
		for i := p; i <= m; i += p {
			e[i] = true
		}
	}

	var ans []int
	for i := 1; i <= m; i++ {
		if !e[i] {
			ans = append(ans, i)
		}
	}
	return ans
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
