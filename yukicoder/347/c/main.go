package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func divide(x int) []int {
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
	sort.Ints(res)
	return res
}

func solve(n, h int) ([][]int, error) {
	w := n / h

	var needRotate bool
	if w < h {
		h, w = w, h
		needRotate = true
	}
	div := divide(n)
	//fmt.Println(div)
	f := make([][]int, h)
	for i := 0; i < h; i++ {
		f[i] = make([]int, w)
	}
	cur := 0
	idx := len(div) - 2
	for div[idx] >= w && idx >= 0 {
		if div[idx]%w != 0 {
			return nil, errors.New("Can not construct.")
		}
		nc := cur + div[idx]/w
		for i := cur; i < nc; i++ {
			for j := 0; j < w; j++ {
				f[i][j] = div[idx]
			}
		}
		cur = nc
		idx--
	}
	//
	j := 0
	for idx >= 0 {
		v, cnt := div[idx], div[idx]
		for cnt > 0 {
			f[cur][j] = v
			cnt--
			j++
		}
		idx--
	}
	if !needRotate {
		return f, nil
	}
	ans := make([][]int, w)
	for i := 0; i < w; i++ {
		ans[i] = make([]int, h)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ans[j][i] = f[i][j]
		}
	}
	return ans, nil
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, h := nextInt(), nextInt()

	ans, err := solve(n, h)
	if err != nil {
		PrintInt(-1)
		return
	}
	PrintVertically(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
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

func PrintVertically(x [][]int) {
	defer out.Flush()
	for _, v := range x {
		PrintHorizonaly(v)
	}
}
