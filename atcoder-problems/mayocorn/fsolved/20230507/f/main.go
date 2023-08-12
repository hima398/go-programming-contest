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

	n, k := nextInt(), nextInt()
	//ans := solveHonestly(n, k)
	ans := solve(n, k)
	PrintInt(ans)
}

func solveHonestly(n, k int) int {
	const p = int(1e9) + 7
	memo := make(map[int]int)
	var ans int
	var dfs func(a []int)
	dfs = func(a []int) {
		if len(a) == n {
			gcd := a[0]
			for i := 1; i < n; i++ {
				gcd = Gcd(gcd, a[i])
			}
			ans += gcd
			memo[gcd]++
			return
		}
		for i := 1; i <= k; i++ {
			na := make([]int, len(a))
			copy(na, a)
			na = append(na, i)
			dfs(na)
		}
	}
	var a []int
	dfs(a)
	fmt.Println(memo)
	digest := make(map[int]int)
	for _, v := range memo {
		digest[v]++
	}
	fmt.Println(digest)
	return ans
}

func computeDivisor(x int) []int {
	m := make(map[int]struct{})
	for i := 1; i*i <= x; i++ {
		if x%i == 0 {
			m[i] = struct{}{}
			m[x/i] = struct{}{}
		}
	}
	var res []int
	for k := range m {
		res = append(res, k)
	}
	return res
}
func solve(n, k int) int {
	const p = int(1e9) + 7

	e := NewSieveOfEratosthenes(k)
	var primes []int
	for i := 2; i <= k; i++ {
		if !e.IsPrime(i) {
			continue
		}
		primes = append(primes, i)
	}
	f := make([]int, k+1)
	for i := 1; i <= k; i++ {
		//k以下でiの倍数の数字の数
		num := Floor(k, i)
		//num種類の数をn個並べて作れる数列の数
		f[i] = Pow(num, n, p)
	}
	//f[x]の中には重複して数えられているものがあるので後から引く数
	//例えばf[2]のとき、g[2]=f[4]+f[6]+f[8]+....
	for i := k; i > 1; i-- {
		d := computeDivisor(i)
		//fmt.Println(i, d)
		for _, v := range d {
			if i == v {
				continue
			}
			f[v] = (f[v] - f[i] + p) % p
		}
	}

	//fmt.Println(f)
	var ans int
	for i := 1; i <= k; i++ {
		ans += i * f[i]
		ans %= p
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

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
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

func Lcm(x, y int) int {
	// x*yのオーバーフロー対策のため先にGcdで割る
	// Gcd(x, y)はxの約数のため割り切れる
	ret := x / Gcd(x, y)
	ret *= y
	return ret
}

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}

func Inv(x, p int) int {
	return Pow(x, p-2, p)
}
