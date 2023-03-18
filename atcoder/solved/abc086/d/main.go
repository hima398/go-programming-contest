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
	var x, y []int
	var c []string
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
		c = append(c, nextString())
	}
	ans := solve(n, k, x, y, c)
	PrintInt(ans)
}

type field struct {
	//f [][]int
	f [4001][4001]int32
}

func (f *field) query(si, sj, ti, tj int) int {
	res := f.f[ti][tj]
	if si > 0 {
		res -= f.f[si-1][tj]
	}
	if sj > 0 {
		res -= f.f[ti][sj-1]
	}
	if si > 0 && sj > 0 {
		res += f.f[si-1][sj-1]
	}
	return int(res)

}

func solve(n, k int, x, y []int, c []string) int {
	k2 := 2 * k
	b := &field{}
	w := &field{}
	//b.f = make([][]int, 2*k2+1)
	//w.f = make([][]int, 2*k2+1)
	//for i := 0; i <= 2*k2; i++ {
	//	b.f[i] = make([]int, 2*k2+1)
	//	w.f[i] = make([]int, 2*k2+1)
	//}
	for l := 0; l < n; l++ {
		i, j := y[l]%k2, x[l]%k2
		if c[l] == "B" {
			b.f[i][j]++
			b.f[i+k2][j]++
			b.f[i][j+k2]++
			b.f[i+k2][j+k2]++
		} else {
			w.f[i][j]++
			w.f[i+k2][j]++
			w.f[i][j+k2]++
			w.f[i+k2][j+k2]++
		}
	}
	//printSlice(b)
	for i := 0; i <= 2*k2; i++ {
		for j := 0; j <= 2*k2; j++ {
			if i > 0 {
				b.f[i][j] += b.f[i-1][j]
				w.f[i][j] += w.f[i-1][j]
			}
			if j > 0 {
				b.f[i][j] += b.f[i][j-1]
				w.f[i][j] += w.f[i][j-1]
			}
			if i > 0 && j > 0 {
				b.f[i][j] -= b.f[i-1][j-1]
				w.f[i][j] -= w.f[i-1][j-1]
			}
		}
	}

	var ans int
	for i := 0; i < k2; i++ {
		for j := 0; j < k2; j++ {
			nb := b.query(i, j, i+k-1, j+k-1) + b.query(i+k, j+k, i+k2-1, j+k2-1)
			nw := w.query(i, j+k, i+k-1, j+k2-1) + w.query(i+k, j, i+k2-1, j+k-1)
			ans = Max(ans, nb+nw)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
