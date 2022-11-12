package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type sushi struct {
	i, x, v int
}

func compute(n int, ss, rs []sushi) int {
	//時計回りにi個目の寿司にたどり着くまでの累積和
	d := make([]int, n+1)
	//時計回りにi個目の寿司を食べて得られるカロリーの累積和
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		d[i] = d[i-1] + ss[i-1].x
		s[i] = s[i-1] + ss[i-1].v - ss[i-1].x
	}
	//反時計回りにi個目の寿司を食べて得られるカロリーの累積和
	r := make([]int, n+1)
	//反時計回りにi個目までの寿司までの間に得られるカロリーの最大値
	m := make([]int, n+1)
	for i := 1; i <= n; i++ {
		r[i] = r[i-1] + rs[i-1].v - rs[i-1].x
		m[i] = Max(m[i-1], r[i-1]+(rs[i-1].v-rs[i-1].x))
	}

	res := 0
	for i := 1; i <= n; i++ {
		//時計回りにi個の寿司を食べるカロリー
		res = Max(res, s[i])
		//開始地点に戻って反時計回りに残りn-i個から得られる最大カロリー
		res2 := s[i] - d[i] + m[n-i]
		res = Max(res, res2)
	}
	return res
}

func solve(n, c int, x, v []int) int {
	x = append([]int{0}, x...)
	var ss, rs []sushi
	for i := 1; i <= n; i++ {
		ss = append(ss, sushi{i - 1, x[i] - x[i-1], v[i-1]})
	}
	x = append(x[1:], c)
	for i := n; i >= 1; i-- {
		rs = append(rs, sushi{n - i, x[i] - x[i-1], v[i-1]})
	}
	//fmt.Println(ss)
	//fmt.Println(rs)
	ans := compute(n, ss, rs)
	ans = Max(ans, compute(n, rs, ss))
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, c := nextInt(), nextInt()
	var x, v []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		v = append(v, nextInt())
	}
	ans := solve(n, c, x, v)
	PrintInt(ans)
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
