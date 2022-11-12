package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type Pos struct {
	i, x, y int
}

func selectRand(l, r int) int {
	return l + rand.Intn(int(1e6))%(r-l+1)
}

func computeDist(x1, y1, x2, y2 int) float64 {
	xx, yy := float64(x2-x1), float64(y2-y1)
	return math.Sqrt(xx*xx + yy*yy)
}

func computeDist2(x1, y1, x2, y2 int) int {
	xx, yy := x2-x1, y2-y1
	return xx*xx + yy*yy
}

func computeScore(p, x, y []int) int {
	res := 0
	n := len(p)
	for i := 1; i < n; i++ {
		res += computeDist2(x[p[i-1]-1], y[p[i-1]-1], x[p[i]-1], y[p[i]-1])
	}
	return res
}

func computeScoref(p, x, y []int) float64 {
	res := 0.0
	n := len(p)
	for i := 1; i < n; i++ {
		res += computeDist(x[p[i-1]-1], y[p[i-1]-1], x[p[i]-1], y[p[i]-1])
	}
	return res
}

func reverse(l, r int, p []int) []int {
	l--
	r--
	res := make([]int, len(p))
	copy(res, p)
	for i := 0; i < (r-l)/2; i++ {
		j := r - i
		res[l+i], res[j] = res[j], res[l+i]
	}
	return res
}

func solveSimulatedAnnealing(n int, x, y []int) (ans []int) {
	var p []int
	for i := 1; i <= n; i++ {
		p = append(p, i)
	}
	p = append(p, 1)
	//p := solveGreedy(n, x, y)
	currentScore := computeScoref(p, x, y)
	const maxCount = int(1e6)
	//const maxCount = 5
	for cnt := 0; cnt < maxCount; cnt++ {
		l, r := selectRand(2, n), selectRand(2, n)
		if l > r {
			l, r = r, l
		}
		//fmt.Println("l, r = ", l, r)
		np := reverse(l, r, p)
		//fmt.Println("p = ", np)
		newScore := computeScoref(np, x, y)

		t := 30.0 - 28.0*float64(cnt)/float64(maxCount)
		//fmt.Println("t = ", t)
		probability := math.Exp(math.Min(0.0, float64(currentScore-newScore)/t))
		newProbability := rand.Float64()
		//fmt.Println("cur, new = ", probability, newProbability)
		if newProbability < probability {
			currentScore = newScore
			p = np
		}

		if currentScore >= newScore {
			currentScore = newScore
			p = np
		}
	}
	for _, v := range p {
		ans = append(ans, v)
	}
	return ans
}

func solveLocalSearch(n int, x, y []int) (ans []int) {
	p := solveGreedy(n, x, y)
	currentScore := computeScore(p, x, y)
	const maxCount = int(1e6)
	//const maxCount = 5
	for cnt := 0; cnt < maxCount; cnt++ {
		l, r := selectRand(2, n), selectRand(2, n)
		if l > r {
			l, r = r, l
		}

		np := reverse(l, r, p)

		newScore := computeScore(np, x, y)
		if currentScore >= newScore {
			currentScore = newScore
			p = np
		}
	}
	for _, v := range p {
		ans = append(ans, v)
	}
	return ans
}

func solveGreedy(n int, x, y []int) (ans []int) {
	visited := make([]bool, n)
	current := 0
	ans = append(ans, current+1)
	visited[current] = true
	for i := 1; i < n; i++ {
		dist2 := int(1e8) + 1
		nextId := -1
		for j := 0; j < n; j++ {
			if visited[j] {
				continue
			}
			nextDist := computeDist2(x[current], y[current], x[j], y[j])
			if nextDist < dist2 {
				nextId = j
				dist2 = nextDist
			}
		}
		current = nextId
		ans = append(ans, current+1)
		visited[current] = true
	}
	ans = append(ans, 1)
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	rand.Seed(time.Now().UnixNano())

	n := nextInt()
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	//ans := solveGreedy(n, x, y)
	//ans := solveLocalSearch(n, x, y)
	ans := solveSimulatedAnnealing(n, x, y)
	PrintVertically(ans)
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

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
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

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}

func Sqrt(x int) int {
	x2 := int(math.Sqrt(float64(x))) - 1
	for (x2+1)*(x2+1) <= x {
		x2++
	}
	return x2
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

func Permutation(N, K int) int {
	v := 1
	if 0 < K && K <= N {
		for i := 0; i < K; i++ {
			v *= (N - i)
		}
	} else if K > N {
		v = 0
	}
	return v
}

func Factional(N int) int {
	return Permutation(N, N-1)
}

func Combination(N, K int) int {
	if K == 0 {
		return 1
	}
	if K == 1 {
		return N
	}
	return Combination(N, K-1) * (N + 1 - K) / K
}

type Comb struct {
	n, p int
	fac  []int // Factional(i) mod p
	finv []int // 1/Factional(i) mod p
	inv  []int // 1/i mod p
}

func NewCombination(n, p int) *Comb {
	c := new(Comb)
	c.n = n
	c.p = p
	c.fac = make([]int, n+1)
	c.finv = make([]int, n+1)
	c.inv = make([]int, n+1)

	c.fac[0] = 1
	c.fac[1] = 1
	c.finv[0] = 1
	c.finv[1] = 1
	c.inv[1] = 1
	for i := 2; i <= n; i++ {
		c.fac[i] = c.fac[i-1] * i % p
		c.inv[i] = p - c.inv[p%i]*(p/i)%p
		c.finv[i] = c.finv[i-1] * c.inv[i] % p
	}
	return c
}

func (c *Comb) Factional(x int) int {
	return c.fac[x]
}

func (c *Comb) Combination(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	ret := c.fac[n] * c.finv[k]
	ret %= c.p
	ret *= c.finv[n-k]
	ret %= c.p
	return ret
}

//重複組み合わせ H
func (c *Comb) DuplicateCombination(n, k int) int {
	return c.Combination(n+k-1, k)
}
func (c *Comb) Inv(x int) int {
	return c.inv[x]
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func DivideSlice(A []int, K int) ([]int, []int, error) {

	if len(A) < K {
		return nil, nil, errors.New("")
	}
	return A[:K+1], A[K:], nil
}
