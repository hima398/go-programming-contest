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

	n, q := nextInt(), nextInt()
	var f, t, x []int
	for i := 0; i < q; i++ {
		f = append(f, nextInt()-1)
		t = append(t, nextInt()-1)
		x = append(x, nextInt()-1)
	}
	ans := solve(n, q, f, t, x)
	PrintVertically(ans)
}

func solve(n, q int, f, t, x []int) []int {
	type container struct {
		i int
		b int
	}
	cs := make([]container, n)
	for i := range cs {
		cs[i].i = i
		cs[i].b = -1
	}
	//テーブルiの一番上にあるコンテナ
	tops := make([]int, n)
	for i := range tops {
		tops[i] = i
	}
	ans := make([]int, n)
	for k := 0; k < q; k++ {
		//コンテナxiはfiにあることは保証されているので、tops[f[k]]<0のチェックは省略しておく
		nt := tops[f[k]]
		//テーブルfからコンテナを取り出す
		if cs[x[k]].b < 0 {
			tops[f[k]] = -1
		} else {
			tops[f[k]] = cs[x[k]].b
			cs[x[k]].b = -1
		}
		//テーブルfにコンテナを置く
		if tops[t[k]] >= 0 {
			cs[x[k]].b = tops[t[k]]
		}
		tops[t[k]] = nt
	}
	for i := 0; i < n; i++ {
		//机にコンテナが乗っていないケース
		if tops[i] < 0 {
			continue
		}
		cur := tops[i]
		for cur >= 0 {
			//fmt.Println(i, cur)
			ans[cur] = i + 1
			cur = cs[cur].b
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
