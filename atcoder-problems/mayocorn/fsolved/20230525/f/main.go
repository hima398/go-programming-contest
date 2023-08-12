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

	n, m, k := nextInt(), nextInt(), nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	var c, d []int
	for i := 0; i < m; i++ {
		c = append(c, nextInt())
		d = append(d, nextInt())
	}
	ans := solve(n, m, k, a, b, c, d)
	PrintFloat64(ans)
}

func solve(n, m, k int, a, b, c, d []int) float64 {
	ng, ok := 0.0, 1.0
	//砂糖水の濃度xは濃度の高い方からk番目未満か？
	check := func(x float64) bool {
		//水の重さにこれをかけることで、濃度xの砂糖水を作るのに
		//必要な砂糖のグラムになる
		w := x / (1 - x)

		//水d[i]を使って濃度xの砂糖水を作るとき、
		//c[i]と比べてdiff[i]差がある
		var diff []float64
		for i := 0; i < m; i++ {
			diff = append(diff, float64(c[i])-w*float64(d[i]))
		}
		sort.Float64s(diff)

		//n*m個の砂糖水のうち、濃度x未満になるものの個数
		var k2 int
		for i := 0; i < n; i++ {
			v := float64(a[i]) - w*float64(b[i])
			//高橋君が持っている砂糖水ごとに、昇順にソートした
			//青木君の砂糖水の中から濃度x以上になるインデックスを探す
			idx := sort.Search(m, func(j int) bool {
				return v+diff[j] >= 0
			})
			k2 += m - idx
		}
		return k2 < k
	}

	//fmt.Println(math.Log(1e12))
	//for ok-ng > 1e-12 {
	for i := 0; i < 42; i++ {
		mid := (ok + ng) / 2.0
		if check(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	//fmt.Println("cnt = ", cnt)
	return 100.0 * ok
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
