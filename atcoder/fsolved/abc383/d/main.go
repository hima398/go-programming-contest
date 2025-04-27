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

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()

	ans := solve(n)

	Print(ans)
}

func divisors(x int) (map[int]struct{}, error) {
	res := make(map[int]struct{})
	for i := 1; i*i <= x; i++ {
		if x%i == 0 {
			res[i] = struct{}{}
			res[x/i] = struct{}{}
		}
		if len(res) > 9 {
			return nil, errors.New("Impossible")
		}
	}
	//fmt.Println("x = ", x, " res = ", res)
	return res, nil
}

func solve(n int) int {
	const max = 2 * int(1e6)

	sieve := NewSieveOfEratosthenes(max)
	sieve.init(max)
	var ps []int
	for i := 0; i <= max; i++ {
		if sieve.IsPrime(i) {
			ps = append(ps, i)
		}
	}
	//fmt.Println(ps)

	var ans int
	for i := 0; i < len(ps); i++ {
		if ps[i]*ps[i]*ps[i]*ps[i]*ps[i]*ps[i]*ps[i]*ps[i] > n {
			break
		}
		ans++
	}
	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			if ps[i]*ps[j] > n || ps[i]*ps[i]*ps[j] > n || ps[i]*ps[i]*ps[j]*ps[j] > n {
				break
			}
			//fmt.Println(ps[i], ps[j])
			ans++
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
