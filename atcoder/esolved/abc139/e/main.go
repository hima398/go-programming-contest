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

func solve(n int, a [][]int) (int, error) {
	isConnected := make([]map[int]struct{}, n)
	q := make([][]int, n)
	isPlayed := make([][]bool, n)
	for i := 0; i < n; i++ {
		isConnected[i] = make(map[int]struct{})
		isPlayed[i] = make([]bool, n)
		isPlayed[i][i] = true
	}
	var ans int
	var cnt int
	for {
		//ans日目に行われる試合数
		s := 0
		//ans日目に試合予定があるかどうか
		isPlayToday := make([]bool, n)
		if ans < n-1 {
			for i := 0; i < n; i++ {
				isConnected[i][a[i][ans]-1] = struct{}{}
				q[i] = append(q[i], a[i][ans]-1)
			}
		}
		for i := 0; i < n; i++ {
			cnt++
			if len(q[i]) == 0 {
				continue
			}
			if len(q[q[i][0]]) == 0 {
				continue
			}
			k1 := q[i][0]
			k2 := q[k1][0]
			if i == k2 && !isPlayed[k1][k2] && !isPlayed[k2][k1] && !isPlayToday[k1] && !isPlayToday[k2] {
				s++
				q[i] = q[i][1:]
				q[k1] = q[k1][1:]
				isPlayed[k1][k2] = true
				isPlayed[k2][k1] = true
				isPlayToday[k1] = true
				isPlayToday[k2] = true
			}
		}

		if s == 0 {
			break
		}
		ans++
	}
	//fmt.Println("cnt = ", cnt)
	ok := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ok = ok && isPlayed[i][j]
		}
	}
	if ok {
		return ans, nil
	} else {
		return 0, errors.New("Impossible")
	}
}

//試合の依存関係をグラフ化して矛盾がないか調べる
func solve02(n int, a [][]int) (int, error) {
	//試合番号
	idxes := make([][]int, n)
	for i := range idxes {
		idxes[i] = make([]int, n)
	}
	idx := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			idxes[i][j] = idx
			idxes[j][i] = idx
			idx++
		}
	}
	e := make([][]int, n*(n-1)/2)
	ie := make([][]int, n*(n-1)/2)
	degree := make([]int, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := 1; j < n-1; j++ {
			g1, g2 := idxes[i][a[i][j-1]], idxes[i][a[i][j]]
			e[g1] = append(e[g1], g2)
			degree[g1]++
			ie[g2] = append(ie[g2], g1)
		}
	}
	visited := make([]bool, n*(n-1)/2)
	var q []int
	for i, di := range degree {
		if di == 0 {
			q = append(q, i)
			visited[i] = true
		}
	}
	//if len(q) == 0 {
	//	return -1, errors.New("Impossible")
	//}
	dist := make([]int, n*(n-1)/2)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, next := range ie[cur] {
			degree[next]--
			if degree[next] > 0 {
				continue
			}
			if visited[next] {
				continue
			}
			q = append(q, next)
			dist[next] = dist[cur] + 1
		}
	}
	ok := true
	for _, d := range degree {
		ok = ok && d == 0
	}
	var ans int
	for _, d := range dist {
		ans = Max(ans, d)
	}
	ans++
	if ok {
		return ans, nil
	} else {
		return -1, errors.New("Impossible")
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextIntSlice(n - 1)
		for j := 0; j < n-1; j++ {
			a[i][j]--
		}
	}
	//ans, err := solve(n, a)
	ans, err := solve02(n, a)
	if err != nil {
		PrintInt(-1)
	} else {
		PrintInt(ans)
	}
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

func PrintHorizonaly(x []bool) {
	defer out.Flush()
	fmt.Fprintf(out, "%v", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %v", x[i])
	}
	fmt.Fprintln(out)
}

func PrintVertically(x [][]bool) {
	defer out.Flush()
	for _, v := range x {
		//fmt.Fprintln(out, v)
		PrintHorizonaly(v)
	}
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
