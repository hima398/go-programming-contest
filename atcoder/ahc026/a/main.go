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
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type result struct {
	v, i int
}

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	//すべてのテストケースにおいて n = 200, m = 10で固定
	n, m := nextInt(), nextInt()
	var b [][]int
	for i := 0; i < m; i++ {
		b = append(b, nextIntSlice(n/m))
	}
	ans := solveBeamSearch(n, m, b)
	//ans := solveGreedily(n, m, b)
	//ans := solveRandomly(n, m, b)
	for _, v := range ans {
		fmt.Println(v.v, v.i)
	}
}

type stack struct {
	data []int
}

func NewStack() *stack {
	return new(stack)
}

func (s *stack) append(x []int) {
	s.data = append(s.data, x...)
}

// 番号xの箱を山から運び出せるかどうか
func (s *stack) canCarry(x int) (bool, bool) {
	if len(s.data) == 0 {
		return false, false
	}
	var found bool
	for _, v := range s.data {
		found = found || v == x
	}
	l := len(s.data)
	canCarry := s.data[l-1] == x
	return canCarry, found
}

func (s *stack) print() {
	fmt.Println(s.data)
}

func (s *stack) liftManyBoxes(x int) []int {
	idx := -1
	for i, v := range s.data {
		if v == x {
			idx = i
		}
	}
	res := s.data[idx+1:]
	s.data = s.data[:idx+1]
	return res
}

func (s *stack) carryOut() {
	s.data = s.data[:len(s.data)-1]
}

type state struct {
	score      int
	operations []result
	stacks     []stack
	carryOut   int
}

func solveBeamSearch(n, m int, b [][]int) []result {
	const beamWidth = 500

	//スタックのデータ型を変換
	ss := make([]stack, m)
	for i := range ss {
		ss[i] = *NewStack()
		ss[i].append(b[i])
	}
	var dp []state
	dp = append(dp, state{0, []result{}, ss, 0})
	mn := 10000
	var ans []result
	for len(dp) > 0 {
		//fmt.Println(len(dp), dp[0])
		var next []state
		for _, cur := range dp {
			//fmt.Println(cur.carryOut)
			if cur.carryOut == 200 {
				if mn > cur.score {
					mn = cur.score
					ans = cur.operations
					fmt.Println(mn)
				}
			}
			stacks := cur.stacks
			nextBox := cur.carryOut + 1
			for i := 0; i < m; i++ {
				canCarry, found := stacks[i].canCarry(nextBox)
				if !found {
					continue
				}
				nextScore := cur.score

				//箱がスタックの中にある
				if !canCarry {
					for j := 0; j < m; j++ {
						if j == i {
							continue
						}
						var nextStacks []stack
						for _, v := range stacks {
							var u stack
							u.data = make([]int, len(v.data))
							copy(u.data, v.data)
							nextStacks = append(nextStacks, u)
						}
						boxes := nextStacks[i].liftManyBoxes(nextBox)
						nextStacks[j].append(boxes)

						var nextResult []result
						nextResult = append(nextResult, cur.operations...)
						nextResult = append(nextResult, result{boxes[0], j + 1})

						nextStacks[i].carryOut()
						nextResult = append(nextResult, result{nextBox, 0})
						nextScore += len(boxes) + 1
						next = append(next, state{nextScore, nextResult, nextStacks, nextBox})
					}
				} else {
					var nextStacks []stack
					for _, v := range stacks {
						var u stack
						u.data = make([]int, len(v.data))
						copy(u.data, v.data)

						nextStacks = append(nextStacks, u)
					}
					var nextResult []result
					nextResult = append(nextResult, cur.operations...)

					nextStacks[i].carryOut()
					nextResult = append(nextResult, result{nextBox, 0})
					//nextScore++
					next = append(next, state{nextScore, nextResult, nextStacks, nextBox})
				}
			}
		}
		if len(next) > beamWidth {
			next = next[:beamWidth]
		}
		sort.Slice(next, func(i, j int) bool {
			return next[i].score < next[j].score
		})
		dp = next
	}
	return ans
}

func solveGreedily(n, m int, b [][]int) []result {
	//スタックのデータ型を変換
	ss := make([]stack, m)
	for i := range ss {
		ss[i] = *NewStack()
		ss[i].append(b[i])
	}

	var ans []result
	for i := 1; i <= n; i++ {
		for j := 0; j < m; j++ {
			canCarry, found := ss[j].canCarry(i)
			if !found {
				continue
			}
			//箱がスタックの中にある
			if !canCarry {
				boxes := ss[j].liftManyBoxes(i)
				to := -1
				mn := n
				for k := 0; k < m; k++ {
					if k == j {
						continue
					}
					if mn > len(ss[k].data) {
						mn = len(ss[k].data)
						to = k
					}
				}
				ss[to].append(boxes)
				ans = append(ans, result{boxes[0], to + 1})
			}
			ss[j].carryOut()
			ans = append(ans, result{i, 0})
		}
	}
	return ans
}

func solveRandomly(n, m int, b [][]int) []result {
	//スタックのデータ型を変換
	ss := make([]stack, m)
	for i := range ss {
		ss[i] = *NewStack()
		ss[i].append(b[i])
	}

	var ans []result
	for i := 1; i <= n; i++ {
		//if i == 11 {
		//	for _, v := range ss {
		//		v.print()
		//	}
		//}
		for j := 0; j < m; j++ {
			canCarry, found := ss[j].canCarry(i)
			if !found {
				continue
			}
			//箱がスタックの中にある
			if !canCarry {
				boxes := ss[j].liftManyBoxes(i)
				to := j
				for to == j {
					to = rand.Intn(m)
				}
				ss[to].append(boxes)
				ans = append(ans, result{boxes[0], to + 1})
			}
			ss[j].carryOut()
			ans = append(ans, result{i, 0})
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

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
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

// 重複組み合わせ H
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
