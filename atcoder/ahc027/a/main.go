package main

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/liyue201/gostl/ds/queue"
	"github.com/liyue201/gostl/ds/stack"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)
var start int64

func input() (int, []string, []string, [][]int) {
	n := nextInt()
	var h []string
	for i := 0; i < n-1; i++ {
		h = append(h, nextString())
	}
	var v []string
	for i := 0; i < n; i++ {
		v = append(v, nextString())
	}
	var d [][]int
	for i := 0; i < n; i++ {
		d = append(d, nextIntSlice(n))
	}
	return n, h, v, d
}

func ExportImage(dir, name string, n int, field [][]int) {
	dest := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{n, n}})
	mx := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mx = Max(mx, field[i][j])
		}
	}
	//fmt.Println("mx = ", mx)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			v := field[y][x] * 255 / mx
			v = Abs(v - 255)
			//fmt.Println(y, x, field[y][x], mx)
			dest.Set(x, y, color.Gray{uint8(v)})
		}
	}
	path := filepath.Join(dir, name)
	file, err := os.Create(path)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	png.Encode(file, dest)
}

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, h, v, d := input()

	//ans := solve00(n, h, v, d)
	//ans := solve01(n, h, v, d)
	ans := solve02(n, h, v, d)

	Print(ans)
}

func computeScore(n int, d [][]int, ans string) (int, error) {
	l := len(ans)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	var ci, cj int
	for _, c := range ans {
		switch c {
		case 'U':
			ci--
		case 'D':
			ci++
		case 'L':
			cj--
		case 'R':
			cj++
		}
		if ci < 0 || ci >= n || cj < 0 || cj >= n {
			return -1, errors.New("Impossible")
		}
		//掃除
		f[ci][cj] = 0
		//汚れ
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == ci && j == cj {
					continue
				}
				f[i][j] += d[i][j]
			}
		}
	}
	var score int
	for _, c := range ans {
		switch c {
		case 'U':
			ci--
		case 'D':
			ci++
		case 'L':
			cj--
		case 'R':
			cj++
		}
		if ci < 0 || ci >= n || cj < 0 || cj >= n {
			return -1, errors.New("Impossible")
		}
		//掃除
		f[ci][cj] = 0
		//汚れ
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == ci && j == cj {
					continue
				}
				f[i][j] += d[i][j]
			}
		}
		var st int
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				st += f[i][j]
			}
		}
		score += st
	}
	return score / l, nil
}

// スコアをO(Min(N**2, L))くらいで求める
func computeScore2(n int, d [][]int, ans string) (int, error) {
	l := len(ans)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	var sd int
	for i := range d {
		for j := range d[i] {
			sd += d[i][j]
		}
	}

	var ci, cj int
	var st int
	for i, c := range ans {
		switch c {
		case 'U':
			ci--
		case 'D':
			ci++
		case 'L':
			cj--
		case 'R':
			cj++
		}
		if ci < 0 || ci >= n || cj < 0 || cj >= n {
			return -1, errors.New("Impossible")
		}
		f[ci][cj] = i + 1
		st += sd
		st -= (i + 1 - f[ci][cj]) * d[ci][cj]
	}
	var score int
	for i, c := range ans {
		switch c {
		case 'U':
			ci--
		case 'D':
			ci++
		case 'L':
			cj--
		case 'R':
			cj++
		}
		if ci < 0 || ci >= n || cj < 0 || cj >= n {
			return -1, errors.New("Impossible")
		}
		f[ci][cj] = i + 1
		st += sd
		st -= (i + 1 - f[ci][cj]) * d[ci][cj]
		score += st
	}
	return score / l, nil
}

func filterField(n int, h, v []string, f [][]int) [][]int {
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	filtered := make([][]int, n)
	for i := range filtered {
		filtered[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			num := 1
			sum := f[i][j]
			for k := 0; k < 4; k++ {
				ni, nj := i+di[k], j+dj[k]
				//はみ出していないかチェック
				if ni < 0 || ni >= n || nj < 0 || nj >= n {
					continue
				}
				//壁のチェック
				if di[k] == 0 && v[i][Min(j, nj)] == '1' {
					continue
				}
				if dj[k] == 0 && h[Min(i, ni)][j] == '1' {
					continue
				}

				sum += f[ni][nj]
				num++
			}
			filtered[i][j] = sum / num
		}
	}
	return filtered
}

func computeG(n int, f [][]int) [][]int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
	}
	type cell struct {
		i, j, a int
	}
	var cs []cell
	for i := range f {
		for j := range f[i] {
			cs = append(cs, cell{i, j, f[i][j]})
		}
	}
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].a > cs[j].a
	})
	for k := 0; k < n; k++ {
		i, j := cs[k].i, cs[k].j
		g[i][j] = 1
	}
	return g
}

var dirs = DirSlice{
	{0, "U", -1, 0},
	{1, "L", 0, -1},
	{2, "D", 1, 0},
	{3, "R", 0, 1},
}

type Robot struct {
	i, j int
	plan *stack.Stack[int]
}

func (r *Robot) move() string {
	p := r.plan.Pop()
	dir := dirs[p]
	r.i, r.j = r.i+dir.di, r.j+dir.dj
	return dir.c
}

// 汚い場所と塗っていない場所を優先する
func solve02(n int, h, v []string, d [][]int) string {
	start = time.Now().UnixNano()

	type cell struct {
		i, j int
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	var g [][]int

	//最も近い汚れに向かう次の位置を返す
	bfs := func(n, i, j int) (int, int, [][]int, *stack.Stack[int]) {
		const INF = 1 << 60
		dist := make([][]int, n)
		for i := range dist {
			dist[i] = make([]int, n)
			for j := range dist[i] {
				dist[i][j] = INF
			}
		}
		q := queue.New[cell]()
		q.Push(cell{i, j})
		dist[i][j] = 0
		for !q.Empty() {
			cur := q.Pop()
			for k := 0; k < 4; k++ {
				ni, nj := cur.i+dirs[k].di, cur.j+dirs[k].dj
				//はみ出していないかチェック
				if ni < 0 || ni >= n || nj < 0 || nj >= n {
					continue
				}
				//壁のチェック
				if dirs[k].di == 0 && v[cur.i][Min(cur.j, nj)] == '1' {
					continue
				}
				if dirs[k].dj == 0 && h[Min(cur.i, ni)][cur.j] == '1' {
					continue
				}
				//訪問すみかチェック
				if dist[ni][nj] < INF {
					continue
				}
				q.Push(cell{ni, nj})
				dist[ni][nj] = dist[cur.i][cur.j] + 1
				if g[ni][nj] == 1 {
					res := stack.New[int]()
					for dist[ni][nj] > 0 {
						for kk := 0; kk < 4; kk++ {
							nni, nnj := ni+dirs[kk].di, nj+dirs[kk].dj
							//はみ出していないかチェック
							if nni < 0 || nni >= n || nnj < 0 || nnj >= n {
								continue
							}
							//壁のチェック
							if dirs[kk].di == 0 && v[ni][Min(nj, nnj)] == '1' {
								continue
							}
							if dirs[kk].dj == 0 && h[Min(ni, nni)][nj] == '1' {
								continue
							}

							if dist[nni][nnj] == dist[ni][nj]-1 {
								ni, nj = nni, nnj
								res.Push((kk + 2) % 4)
								break
							}
						}
					}
					//res.Push(k)
					return ni, nj, dist, res
				}
			}
		}
		return -1, -1, nil, nil
	}

	robot := Robot{0, 0, stack.New[int]()}
	var ans []string

	visited := make([][]int, n)
	for i := range visited {
		visited[i] = make([]int, n)
		for j := range visited[i] {
			visited[i][j] = 1
		}
	}
	visited[0][0] = 0
	goal := make([][]int, n)
	for i := range goal {
		goal[i] = make([]int, n)
	}
	goal[0][0] = 1
	rem := n*n - 1
	for {
		//fmt.Println(time.Now().UnixNano() - start)
		//for cnt := 0; cnt < int(1e5); cnt++ {
		//fmt.Println(ci, cj)
		//次の汚れを計算
		var ave int
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == robot.i && j == robot.j {
					continue
				}
				f[i][j] += d[i][j]
				ave += f[i][j]
			}
		}

		if robot.plan.Empty() {
			if rem > 0 {
				//fmt.Println("まず全面掃除")
				g = visited
			} else if time.Now().UnixNano()-start < 15*int64(1e8) {
				//fmt.Println("汚れているところを掃除")
				g = computeG(n, f)
			} else {
				//fmt.Println("0, 0に向かう")
				g = goal
				if robot.i == 0 && robot.j == 0 {
					break
				}
			}
			_, _, _, res := bfs(n, robot.i, robot.j)
			robot.plan = res
		}
		r := robot.move()
		if visited[robot.i][robot.j] == 1 {
			visited[robot.i][robot.j] = 0
			rem--
		}
		//fmt.Println(r)
		ans = append(ans, r)
		f[robot.i][robot.j] = 0

		/*
			if cnt == 500 {
				ExportImage("./image", "0000.png", n, g)
				for i := range dist {
					for j := range dist[i] {
						if dist[i][j] == 1<<60 {
							dist[i][j] = 0
						}
					}
					fmt.Println(dist[i])
				}
				fmt.Println()
				for i := range f {
					fmt.Println(f[i])
				}
			}
		*/

		//fmt.Println(robot.i, robot.j, r, robot.plan)
	}

	return strings.Join(ans, "")
}

type dir struct {
	id     int
	c      string
	di, dj int
}

type DirSlice []dir

func (x DirSlice) Len() int {
	return len(x)
}
func (x DirSlice) Less(i, j int) bool {
	return x[i].id < x[j].id
}
func (x DirSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// スタートから24種類のDFSをして戻ってくるうち最適なものを選択
func solve01(n int, h, v []string, d [][]int) string {
	type dir struct {
		id     int
		c      string
		di, dj int
	}
	dirs := DirSlice{
		{0, "U", -1, 0},
		{1, "D", 1, 0},
		{2, "L", 0, -1},
		{3, "R", 0, 1},
	}
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var candidate []string

	var dfs func(i, j int)
	dfs = func(i, j int) {
		visited[i][j] = true
		for k := 0; k < 4; k++ {
			ni, nj := i+dirs[k].di, j+dirs[k].dj
			//はみ出していないかチェック
			if ni < 0 || ni >= n || nj < 0 || nj >= n {
				continue
			}
			//壁のチェック
			if dirs[k].di == 0 && v[i][Min(j, nj)] == '1' {
				continue
			}
			if dirs[k].dj == 0 && h[Min(i, ni)][j] == '1' {
				continue
			}
			//訪問すみかチェック
			if visited[ni][nj] {
				continue
			}
			candidate = append(candidate, dirs[k].c)
			dfs(ni, nj)
			var r string
			switch dirs[k].c {
			case "U":
				r = "D"
			case "D":
				r = "U"
			case "L":
				r = "R"
			case "R":
				r = "L"
			}
			candidate = append(candidate, r)
		}
	}

	score := 1 << 60
	var ans string
	for {
		//訪問先を初期化
		for i := range visited {
			for j := range visited[i] {
				visited[i][j] = false
			}
		}
		//ルートの候補も初期化
		candidate = []string{}
		dfs(0, 0)
		varS, err := computeScore(n, d, strings.Join(candidate, ""))
		if err != nil {
			continue
		}
		if varS < score {
			score = varS
			ans = strings.Join(candidate, "")
		}
		if !NextPermutation(dirs) {
			break
		}
	}
	return ans
}

// スタートからDFSをして戻ってくるサンプル
func solve00(n int, h, v []string, d [][]int) string {
	//U, L, D, Rの順に方向を検索
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	dir := []string{"U", "L", "D", "R"}

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var ans []string

	var dfs func(i, j int)
	dfs = func(i, j int) {
		visited[i][j] = true
		for k := 0; k < 4; k++ {
			ni, nj := i+di[k], j+dj[k]
			//はみ出していないかチェック
			if ni < 0 || ni >= n || nj < 0 || nj >= n {
				continue
			}
			//壁のチェック
			if di[k] == 0 && v[i][Min(j, nj)] == '1' {
				continue
			}
			if dj[k] == 0 && h[Min(i, ni)][j] == '1' {
				continue
			}
			//訪問すみかチェック
			if visited[ni][nj] {
				continue
			}
			ans = append(ans, dir[k])
			dfs(ni, nj)
			ans = append(ans, dir[(k+2)%4])
		}
	}
	dfs(0, 0)

	return strings.Join(ans, "")
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
