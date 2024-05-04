package main

import (
	"bufio"
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

// 探索の方向
// 左->上->右->下
var dirs = [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

// 水路を表す構造体
type waterway struct {
	h, w     int
	waterway [][][4]bool
}

// 場所(i, j)にいるとき、k方向(0<=k<=3)に移動可能かを返す
func (ww *waterway) canNotMove(i, j, k int) bool {
	return ww.waterway[i][j][k]
}

// 土地
type land struct {
	h, w int
	land [][]int
}

// 場所(i, j)にすでに作物があるかを返す
func (l *land) isPlanted(i, j int) bool {
	return l.land[i][j] > 0
}

// 作物を表す構造体
type crop struct {
	k, s, d int
}

// 解答の構造体
type answer struct {
	k    int //type
	i, j int //pos
	s    int //month
}

// 解答を出力
func print(ans []answer) {
	PrintInt(len(ans))
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintf(out, "%d %d %d %d\n", v.k, v.i, v.j, v.s)
	}
}

// 入力h_i_j, v_i_jを元に土地を表すデータを作成する
func createWaterway(lh, lw int, h, v []string) waterway {
	var res waterway
	res.h = lh
	res.w = lw
	res.waterway = make([][][4]bool, lh)
	for i := range res.waterway {
		res.waterway[i] = make([][4]bool, lw)
	}
	for i := 0; i < lh-1; i++ {
		for j := 0; j < lw; j++ {
			if h[i][j] == '0' {
				continue
			}
			//i行、j列目の下方向
			res.waterway[i][j][3] = true
			//i+1行、j列目の上方向
			res.waterway[i+1][j][1] = true
		}
	}
	for i := 0; i < lh; i++ {
		for j := 0; j < lw-1; j++ {
			if v[i][j] == '0' {
				continue
			}
			//i行、j列目の右方向
			res.waterway[i][j][2] = true
			//i行、j+1列目の左方向
			res.waterway[i][j+1][0] = true
		}
	}

	return res
}

// Deprecated: 入力h_i_j, v_i_jを元に土地を表すデータを作成する
func createWaterwaySlice(lh, lw int, h, v []string) [][][4]bool {
	res := make([][][4]bool, lh)
	for i := range res {
		res[i] = make([][4]bool, lw)
	}
	for i := 0; i < lh-1; i++ {
		for j := 0; j < lw; j++ {
			if h[i][j] == '0' {
				continue
			}
			//i行、j列目の下方向
			res[i][j][3] = true
			//i+1行、j列目の上方向
			res[i+1][j][1] = true
		}
	}
	for i := 0; i < lh; i++ {
		for j := 0; j < lw-1; j++ {
			if v[i][j] == '0' {
				continue
			}
			//i行、j列目の右方向
			res[i][j][2] = true
			//i行、j+1列目の左方向
			res[i][j+1][0] = true
		}
	}
	return res
}

// 土地を新規作成する
func createLand(h, w int) land {
	var res land
	res.h, res.w = h, w
	res.land = make([][]int, h)
	for i := range res.land {
		res.land[i] = make([]int, w)
	}
	return res
}

// 作物を植える
func plant(land *land, i, j, k int) {
	land.land[i][j] = k
}

// 作物を収穫する
func harvest(land *land, i, j int) {
	land.land[i][j] = 0
}

// 入力形式のデータ構造から作物の構造体を作成する
func createCrops(k int, s, d []int) []crop {
	var res []crop
	for i := 0; i < k; i++ {
		res = append(res, crop{i + 1, s[i], d[i]})
	}
	return res
}

// 土地からはみ出しているかを判定する
func isOutOfLand(h, w, i, j int) bool {
	return i < 0 || i >= h || j < 0 || j >= w
}

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t, h, w, i0 := nextInt(), nextInt(), nextInt(), nextInt()
	var r []string
	for i := 0; i < h-1; i++ {
		r = append(r, nextString())
	}
	var c []string
	for i := 0; i < h; i++ {
		c = append(c, nextString())
	}

	k := nextInt()
	var s, d []int
	for i := 0; i < k; i++ {
		s = append(s, nextInt())
		d = append(d, nextInt())
	}
	ans := solve02(t, h, w, i0, r, c, k, s, d)
	//ans := solve01(t, h, w, i0, r, c, k, s, d)

	print(ans)
}

func solve02(t, h, w, i0 int, r, c []string, k int, s, d []int) []answer {
	//農機のスタート地点
	si, sj := i0, 0
	//水路
	waterway := createWaterway(h, w, r, c)
	//作物
	crops := createCrops(k, s, d)
	//土地
	land := createLand(h, w)

	//ここから解答
	sort.Slice(crops, func(i, j int) bool {
		if crops[i].s == crops[j].s {
			if crops[i].d == crops[j].d {
				return (crops[i].d - crops[i].s) > (crops[j].d - crops[j].s)
			}
			return crops[i].d > crops[j].d
		}
		return crops[i].s < crops[j].s
	})
	fmt.Println(crops)

	bfs := func() [][]int {
		//20という数値に意味を持たせる
		res := make([][]int, 20)
		for i := range res {
			res[i] = make([]int, 20)
			for j := range res[i] {
				res[i][j] = -1
			}
		}
		q := queue.New[[2]int]()
		q.Push([2]int{si, sj})
		res[si][sj] = 0
		for !q.Empty() {
			cur := q.Pop()
			i, j := cur[0], cur[1]
			for k, dir := range dirs {
				//水路があって移動不能
				if waterway.canNotMove(i, j, k) {
					continue
				}

				ni, nj := i+dir[0], j+dir[1]
				//移動さきが土地の範囲外は後続の処理をしない
				if isOutOfLand(h, w, ni, nj) {
					continue
				}
				//すでに探索済み、すでに作物が植えてある
				if res[ni][nj] >= 0 || land.isPlanted(ni, nj) {
					continue
				}
				res[ni][nj] = res[i][j] + 1
				q.Push([2]int{ni, nj})
			}
		}
		return res
	}
	var ans []answer
	for k := 1; k <= t; k++ {
		//作物を収穫
		for {
			var cnt int
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					ok := false
					for l, dir := range dirs {
						if waterway.canNotMove(i, j, l) {
							continue
						}
						ni, nj := i+dir[0], j+dir[1]
						if isOutOfLand(land.h, land.w, ni, nj) {
							continue
						}
						if land.land[ni][nj] == 0 && land.land[i][j] == k {
							ok = true
						}
					}
					if ok {
						cnt++
						harvest(&land, i, j)
					}
				}
			}
			//収穫できる作物がない
			if cnt == 0 {
				break
			}
		}

		//作物を植える
		for len(crops) > 0 {
			cur := crops[0]
			if cur.s != k {
				break
			}
			crops = crops[1:]
			//fmt.Println("cs = ", cs)
			//植えるタイミングを逃しているので植えない
			if cur.s < k {
				continue
			}
			//収穫できるタイミングを逃しているので植えない
			if cur.d < k {
				continue
			}
			dist := bfs()
			ti, tj := -1, -1
			d := 0
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					if dist[i][j] > d {
						ti, tj = i, j
						d = dist[i][j]
					}
				}
			}
			if ti >= 0 && tj >= 0 {
				//land[ti][tj] = cur.d
				plant(&land, ti, tj, cur.d)
				ans = append(ans, answer{cur.k, ti, tj, k})
			}
		}
	}

	return ans

}

func solve01(t, h, w, i0 int, r, c []string, k int, s, d []int) []answer {
	//農機のスタート地点
	si, sj := i0, 0
	//水路
	waterway := createWaterwaySlice(h, w, r, c)

	//作物
	cs := createCrops(k, s, d)
	//土地
	land := make([][]int, 20)
	for i := range land {
		land[i] = make([]int, 20)
	}

	//ここから解答
	sort.Slice(cs, func(i, j int) bool {
		if cs[i].d == cs[j].d {
			return cs[i].s < cs[j].s
		}
		return cs[i].d > cs[j].d
	})
	//cropsはdが遅いものから取り出せる
	//crops := stack.New[crop]()
	//for _, c := range cs {
	//	crops.Push(c)
	//}
	//fmt.Println(crops)

	bfs := func() [][]int {
		//20という数値に意味を持たせる
		res := make([][]int, 20)
		for i := range res {
			res[i] = make([]int, 20)
			for j := range res[i] {
				res[i][j] = -1
			}
		}
		q := queue.New[[2]int]()
		q.Push([2]int{si, sj})
		res[si][sj] = 0
		for !q.Empty() {
			cur := q.Pop()
			i, j := cur[0], cur[1]
			for k, dir := range dirs {
				//水路があって移動不能
				if waterway[i][j][k] {
					continue
				}

				ni, nj := i+dir[0], j+dir[1]
				//移動さきが土地の範囲外は後続の処理をしない
				if isOutOfLand(h, w, ni, nj) {
					continue
				}
				//すでに探索済み、TODO:作物ありの評価
				if res[ni][nj] >= 0 || land[ni][nj] > 0 {
					continue
				}
				res[ni][nj] = res[i][j] + 1
				q.Push([2]int{ni, nj})
			}
		}
		return res
	}
	var ans []answer
	for k := 1; k <= t; k++ {
		//作物を収穫
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if land[i][j] == k {
					land[i][j] = 0
				}
			}
		}

		//作物を植える
		for len(cs) > 0 {
			cur := cs[0]
			cs = cs[1:]
			//fmt.Println("cs = ", cs)
			//植えるタイミングを逃しているので植えない
			if cur.s < k {
				continue
			}
			//収穫できるタイミングを逃しているので植えない
			if cur.d < k {
				continue
			}
			dist := bfs()
			ti, tj := -1, -1
			d := 0
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					if dist[i][j] > d {
						ti, tj = i, j
						d = dist[i][j]
					}
				}
			}
			if ti >= 0 && tj >= 0 {
				land[ti][tj] = cur.d
				ans = append(ans, answer{cur.k, ti, tj, k})
			}
		}
	}

	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
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
