package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/emirpasic/gods/trees/avltree"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k, q := nextInt(), nextInt(), nextInt()
	var x, y []int
	for i := 0; i < q; i++ {
		x = append(x, nextInt()-1)
		y = append(y, nextInt())
	}
	ans := solve(n, k, q, x, y)
	PrintVertically(ans)
}

func solve(n, k, q int, x, y []int) []int {
	a := make([]int, n)
	bFront, bRear := avltree.NewWithIntComparator(), avltree.NewWithIntComparator()
	bFront.Put(0, k)
	bRear.Put(0, n-k)
	bRear.Put(-1, 1) //n==kのとき用の番兵
	var fa int
	var ans []int
	for i := 0; i < q; i++ {
		//fmt.Println(i, a)
		//fmt.Printf("bFront: %v\n", bFront)
		//fmt.Printf("bRear: %v\n", bRear)
		prevAi := a[x[i]]
		a[x[i]] = y[i]
		//bの書き換え
		if prevAi > bRear.Right().Key.(int) {
			//k以内の数が書き換えられる
			v, _ := bFront.Get(prevAi)
			if v.(int) == 1 {
				bFront.Remove(prevAi)
			} else {
				bFront.Put(prevAi, v.(int)-1)
			}
			fa -= prevAi

			//新しく更新される方がk以内か、そうでないか
			if a[x[i]] > bRear.Right().Key.(int) { //fから消えてfを更新
				//a[x[i]]で更新
				v, found := bFront.Get(a[x[i]])
				if found {
					bFront.Put(a[x[i]], v.(int)+1)
				} else {
					bFront.Put(a[x[i]], 1)
				}
				fa += a[x[i]]
			} else { //fから消えてrを更新
				//kより後ろの最大値
				right := bRear.Right()
				v, found := bFront.Get(right.Key)
				if found {
					bFront.Put(right.Key, v.(int)+1)
				} else {
					bFront.Put(right.Key, 1)
				}
				fa += right.Key.(int)
				//(念のため)rightを使い終わってから消す
				if right.Value.(int) == 1 {
					bRear.Remove(right.Key)
				} else {
					bRear.Put(right.Key, right.Value.(int)-1)
				}
				v2, found := bRear.Get(a[x[i]])
				if found {
					bRear.Put(a[x[i]], v2.(int)+1)
				} else {
					bRear.Put(a[x[i]], 1)
				}
			}
		} else {
			//kより後ろの候補が書き換えられる
			v, _ := bRear.Get(prevAi)
			if v.(int) == 1 {
				bRear.Remove(prevAi)
			} else {
				bRear.Put(prevAi, v.(int)-1)
			}

			//新しく更新される値がk以内か、そうでないか
			if a[x[i]] > bRear.Right().Key.(int) { //rから消えてfを更新
				v, found := bFront.Get(a[x[i]])
				if found {
					bFront.Put(a[x[i]], v.(int)+1)
				} else {
					bFront.Put(a[x[i]], 1)
				}
				fa += a[x[i]]
				//k以内の最小値を持ってくる
				left := bFront.Left()
				v2, found := bRear.Get(left.Key)
				if found {
					bRear.Put(left.Key, v2.(int)+1)
				} else {
					bRear.Put(left.Key, 1)
				}
				fa -= left.Key.(int)
				if left.Value == 1 {
					bFront.Remove(left.Key)
				} else {
					bFront.Put(left.Key, left.Value.(int)-1)
				}
			} else {
				v, found := bRear.Get(a[x[i]])
				if found {
					bRear.Put(a[x[i]], v.(int)+1)
				} else {
					bRear.Put(a[x[i]], 1)
				}
			}
		}
		ans = append(ans, fa)
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
