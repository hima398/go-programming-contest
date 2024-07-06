package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	var a [][]int
	for i := 0; i < h; i++ {
		a = append(a, nextIntSlice(w))
	}
	var b [][]int
	for i := 0; i < h; i++ {
		b = append(b, nextIntSlice(w))
	}

	ans, err := solve(h, w, a, b)
	if err != nil {
		Print(-1)
	} else {
		Print(ans)
	}
}

func SwapGrid(r, c []int, a [][]int) [][]int {
	res := make([][]int, len(r))
	for i := 0; i < len(r); i++ {
		res[i] = make([]int, len(c))
	}
	for i, ii := range r {
		for j, jj := range c {
			res[i][j] = a[ii][jj]
		}
	}
	return res
}

func IsSame(a, b [][]int) bool {
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func ComputeInversion(s []int) int {
	var res int
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				res++
			}
		}
	}
	return res
}

func solve(h, w int, a, b [][]int) (int, error) {
	const INF = math.MaxInt
	row := make([]int, h)
	for i := 0; i < h; i++ {
		row[i] = i
	}
	ans := INF
	for {
		col := make([]int, w)
		for j := 0; j < w; j++ {
			col[j] = j
		}
		for {
			t := SwapGrid(row, col, a)
			if IsSame(t, b) {
				ans = Min(ans, ComputeInversion(row)+ComputeInversion(col))
			}
			if !NextPermutation(sort.IntSlice(col)) {
				break
			}
		}
		if !NextPermutation(sort.IntSlice(row)) {
			break
		}
	}
	if ans == INF {
		return -1, errors.New("Impossible")
	}
	return ans, nil
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
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
