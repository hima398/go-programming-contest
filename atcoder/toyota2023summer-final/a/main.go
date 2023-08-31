package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/liyue201/gostl/ds/queue"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	d, n := nextInt(), nextInt()
	var ri, rj []int
	for i := 0; i < n; i++ {
		ri = append(ri, nextInt())
		rj = append(rj, nextInt())
	}
	solve01(d, n, ri, rj)

	//for _, v := range ans {
	//	Print(v[0], v[1])
	//}
}

const h = 9
const w = 9
const INF = math.MaxInt

func computeDist(i1, j1, i2, j2 int) int {
	return Abs(i2-i1) + Abs(j2-j1)
}

func solve01(d, n int, ri, rj []int) {
	//var t []int
	var plan [h][w]int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			plan[i][j] = -1
		}
	}
	//障害物を配置
	for i := 0; i < n; i++ {
		plan[ri[i]][rj[i]] = INF
	}
	q := queue.New[[2]int]()
	q.Push([2]int{0, 4})
	var idx int
	di := [4]int{0, -1, 0, 1}
	dj := [4]int{-1, 0, 1, 0}
	for !q.Empty() {
		cur := q.Pop()
		for k := 0; k < 4; k++ {
			ni, nj := cur[0]+di[k], cur[1]+dj[k]
			if ni < 0 || ni >= h || nj < 0 || nj >= w {
				continue
			}
			//訪問済みもしくは障害物
			if plan[ni][nj] >= 0 {
				continue
			}
			if ni == 0 && nj == 4 {
				continue
			}

			q.Push([2]int{ni, nj})
			plan[ni][nj] = idx
			idx++
		}
	}
	var field [h][w]int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			field[i][j] = -1
		}
	}
	//障害物を配置
	for i := 0; i < n; i++ {
		field[ri[i]][rj[i]] = INF
	}

	//for _, v := range field {
	//	fmt.Println(v)
	//}
	//check := func(f [h][w]int) bool {
	check := func() bool {
		visited := make([][]bool, h)
		for i := range visited {
			visited[i] = make([]bool, w)
			for j := range visited[i] {
				if field[i][j] >= 0 {
					visited[i][j] = true
				}
			}
		}
		q := queue.New[[2]int]()
		q.Push([2]int{0, 4})
		visited[0][4] = true
		di := [4]int{0, -1, 0, 1}
		dj := [4]int{-1, 0, 1, 0}
		for !q.Empty() {
			cur := q.Pop()
			for k := 0; k < 4; k++ {
				ni, nj := cur[0]+di[k], cur[1]+dj[k]
				if ni < 0 || ni >= h || nj < 0 || nj >= w {
					continue
				}
				//訪問済みもしくは障害物
				if visited[ni][nj] {
					continue
				}

				q.Push([2]int{ni, nj})
				visited[ni][nj] = true
			}
		}
		for i := range visited {
			for _, ok := range visited[i] {
				if !ok {
					return false
				}
			}
		}
		return true
	}
	for k := 0; k < d*d-1-n; k++ {
		t := nextInt()
		var ti, tj int
		var fixed bool
		for i := 0; i < h; i++ {
			if fixed {
				break
			}
			for j := 0; j < w; j++ {
				if fixed {
					break
				}
				if plan[i][j] == t {
					ti, tj = i, j
					field[ti][tj] = t
					if check() {
						Print(ti, tj)
						fixed = true
					} else {
						field[ti][tj] = -1
						var cnd [][2]int
						//var cnd2 [][2]int
						for ii := 0; ii < h; ii++ {
							for jj := 0; jj < w; jj++ {
								if ii == 0 && jj == 4 {
									continue
								}
								if field[ii][jj] >= 0 {
									continue
								}
								var cnt [2]int
								for k := 0; k < 4; k++ {
									ni, nj := ii+di[k], jj+dj[k]
									if ni < 0 || ni >= h || nj < 0 || nj >= w {
										continue
									}
									if field[ni][nj] != -1 {
										continue
									}
									cnt[k%2]++
								}
								if cnt[0]+cnt[1] <= 1 { //|| (cnt[0]+cnt[1] <= 2 && cnt[0] > 0 && cnt[1] > 0) {
									cnd = append(cnd, [2]int{ii, jj})
								}
							}
						}
						if len(cnd) == 0 {
							for ii := 0; ii < h; ii++ {
								for jj := 0; jj < w; jj++ {
									if ii == 0 && jj == 4 {
										continue
									}

									if field[ii][jj] >= 0 {
										continue
									}
									var cnt [2]int
									for k := 0; k < 4; k++ {
										ni, nj := ii+di[k], jj+dj[k]
										if ni < 0 || ni >= h || nj < 0 || nj >= w {
											continue
										}
										if field[ni][nj] != -1 {
											continue
										}
										cnt[k%2]++
									}
									if cnt[0]+cnt[1] <= 2 && cnt[0] > 0 && cnt[1] > 0 {
										cnd = append(cnd, [2]int{ii, jj})
									}
								}
							}

						}
						//fmt.Println(t, ti, tj, cnd)
						dist := INF
						ui, uj := -1, -1
						for _, c := range cnd {
							d := computeDist(ti, tj, c[0], c[1])
							if dist > d && d > 0 {
								dist = d
								ui, uj = c[0], c[1]
							}
						}
						plan[ti][tj], plan[ui][uj] = plan[ui][uj], plan[ti][tj]
						ti, tj = ui, uj
						field[ti][tj] = t
						Print(ti, tj)
						fixed = true
					}
				}
			}
		}
	}
	v := make([][]bool, h)
	for i := range v {
		v[i] = make([]bool, w)
	}
	pq := &PriorityQueue{}
	heap.Init(pq)
	v[0][4] = true
	heap.Push(pq, node{0, 3, field[0][3]})
	v[0][3] = true
	heap.Push(pq, node{0, 5, field[0][5]})
	v[0][5] = true
	heap.Push(pq, node{1, 4, field[1][4]})
	v[1][4] = true
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(node)
		Print(cur.i, cur.j)
		for k := 0; k < 4; k++ {
			ni, nj := cur.i+di[k], cur.j+dj[k]
			if ni < 0 || ni >= h || nj < 0 || nj >= w {
				continue
			}
			if v[ni][nj] {
				continue
			}
			heap.Push(pq, node{ni, nj, field[ni][nj]})
			v[ni][nj] = true
		}
	}
}

type node struct {
	i, j, idx int
}

type PriorityQueue []node

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].idx < pq[j].idx
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(node))
}

func (pq *PriorityQueue) Pop() interface{} {
	es := *pq // EdgeのSlice
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

//func Print(x any) {
//	defer out.Flush()
//	fmt.Fprintln(out, x)
//}

func Print(i, j any) {
	defer out.Flush()
	fmt.Fprintf(out, "%d %d\n", i, j)
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
