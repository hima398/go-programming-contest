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

	ans := solve(n)

	PrintInt(ans)
}

func solve(n int) int {
	e := NewSieveOfEratosthenes(int(1e6) + 1)
	var ps []int
	//ps := []int{0} //番兵
	for i := 1; i <= 1e6; i++ {
		if e.IsPrime(i) {
			ps = append(ps, i)
		}
	}

	m := len(ps)
	var ans int
	for i := 0; i < m; i++ {
		a := ps[i]
		if a*a > n {
			continue
		}
		for j := i + 1; j < m; j++ {
			b := ps[j]
			if a*a*b > n {
				break
			}
			for k := j + 1; k < m; k++ {
				c := ps[k]
				if a*a*b*c > n {
					break
				}
				if a*a*b*c*c > n {
					break
				}
				ans++
			}
		}
	}
	return ans
}

// エラトステネスの篩
type SieveOfEratosthenes struct {
	n          int
	isNotPrime []bool
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
}

func (sieve *SieveOfEratosthenes) IsPrime(x int) bool {
	return !sieve.isNotPrime[x]
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
