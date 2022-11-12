package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type SieveOfEratosthenes struct {
	n          int
	isNotPrime []bool
}

func New(n int) *SieveOfEratosthenes {
	return NewSieveOfEratosthenes(n)
}

func NewSieveOfEratosthenes(n int) *SieveOfEratosthenes {
	sieve := new(SieveOfEratosthenes)
	sieve.n = n + 1
	sieve.isNotPrime = make([]bool, sieve.n)
	sieve.isNotPrime[0] = true
	sieve.isNotPrime[1] = true
	for j := 4; j < sieve.n; j += 2 {
		sieve.isNotPrime[j] = true
	}
	for i := 3; i*i < sieve.n; i += 2 {
		if sieve.isNotPrime[i] {
			continue
		}
		for j := i + i; j < sieve.n; j += i {
			sieve.isNotPrime[j] = true
		}
	}
	return sieve
}

func (sieve *SieveOfEratosthenes) IsPrime(x int) bool {
	return !sieve.isNotPrime[x]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	q := nextInt()
	x := nextIntSlice(q)
	ans := solve(q, x)
	PrintVertically(ans)
}

func solve(q int, x []int) []string {
	const maxX = 3 * int(1e5)
	sieve := NewSieveOfEratosthenes(maxX)

	var ans []string
	for _, xi := range x {
		if sieve.IsPrime(xi) {
			ans = append(ans, "Yes")
		} else {
			ans = append(ans, "No")
		}
	}
	return ans
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

func PrintVertically(x []string) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
