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

func solve(n int, a []int) int {
	gcd := a[0]
	for i := 1; i < n; i++ {
		gcd = Gcd(gcd, a[i])
	}
	if gcd != 1 {
		//not coprime
		return 2
	}
	sieve := NewSieveOfEratosthenes(int(1e6))
	m := make(map[int]int)
	for _, ai := range a {
		d := sieve.Divide(ai)
		for k, v := range d {
			if _, found := m[k]; found {
				if k != 1 {
					return 1
				}
			}
			m[k] = v
		}
	}
	return 0
}

// エラトステネスの篩
type SieveOfEratosthenes struct {
	n          int
	isNotPrime []bool
	v          []int
}

func New(n int) *SieveOfEratosthenes {
	return NewSieveOfEratosthenes(n)
}

func NewSieveOfEratosthenes(n int) *SieveOfEratosthenes {
	sieve := new(SieveOfEratosthenes)
	sieve.init(n)
	return sieve
}

func (sieve *SieveOfEratosthenes) init(n int) {
	sieve.n = n + 1
	sieve.v = make([]int, n+1)
	sieve.isNotPrime = make([]bool, sieve.n)
	sieve.isNotPrime[0] = true
	sieve.isNotPrime[1] = true
	for j := 4; j < sieve.n; j += 2 {
		sieve.isNotPrime[j] = true
		sieve.v[j] = 2
	}
	for i := 3; i < sieve.n; i += 2 {
		if sieve.isNotPrime[i] {
			continue
		}
		for j := i + i; j < sieve.n; j += i {
			sieve.isNotPrime[j] = true
			if sieve.v[j] == 0 {
				sieve.v[j] = i
			}
		}
	}
}

func (sieve *SieveOfEratosthenes) IsPrime(x int) bool {
	return !sieve.isNotPrime[x]
}

func (sieve *SieveOfEratosthenes) Divide(x int) map[int]int {
	ret := make(map[int]int)
	for {
		if sieve.v[x] == 0 {
			ret[x]++
			break
		}
		ret[sieve.v[x]]++
		x /= sieve.v[x]
	}
	return ret
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
