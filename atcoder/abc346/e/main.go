package main

import (
	"bufio"
	"fmt"
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

	h, w, m := nextInt(), nextInt(), nextInt()
	var t, a, x []int
	for i := 0; i < m; i++ {
		t = append(t, nextInt())
		a = append(a, nextInt()-1)
		x = append(x, nextInt())
	}
	ans := solve(h, w, m, t, a, x)
	Print(len(ans))
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}
}

func solve(h, w, m int, t, a, x []int) [][2]int {
	colors := make(map[int]int)
	rows, cols := make(map[int]struct{}), make(map[int]struct{})

	for i := m - 1; i >= 0; i-- {
		switch t[i] {
		case 1: //ai行目を全て色xiに塗り替える
			//後から同じ行が塗り替えられるので集計不要
			if _, found := rows[a[i]]; found {
				continue
			}
			colors[x[i]] += w - len(cols)
			rows[a[i]] = struct{}{}
		case 2: //ai列目を全て色xiに塗り替える
			//後から同じ列が塗り替えられるので集計不要
			if _, found := cols[a[i]]; found {
				continue
			}
			colors[x[i]] += h - len(rows)
			cols[a[i]] = struct{}{}
		}
	}
	nz := h * w
	for _, v := range colors {
		nz -= v
	}
	colors[0] += nz
	//fmt.Println(colors)
	var ans [][2]int
	for k, v := range colors {
		if v == 0 {
			continue
		}
		ans = append(ans, [2]int{k, v})
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
