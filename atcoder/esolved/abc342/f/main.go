package main

import (
	"bufio"
	"fmt"
	"math"
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

	n, l, d := nextInt(), nextInt(), nextInt()

	ans := solveHonestly(n, l, d)

	Print(ans)
}

func solveHonestly(n, l, d int) float64 {
	//ディーラーの得点yがiになる確率
	dpy := make([]float64, n+1)
	dpy[0] = 1.0
	for i := 1; i <= n; i++ {
		for j := Max(i-d, 0); j < Min(l, i); j++ {
			dpy[i] += dpy[j] / float64(d)
		}
	}
	fmt.Println(dpy)

	//プレイヤーの得点xがiのとき勝てる確率
	dpx := make([]float64, n+1)
	for i := n; i >= 0; i-- {
		var s1 float64
		for j := Max(l, i); j <= n; j++ {
			s1 += dpy[j]
		}
		p1 := 1.0 - s1

		var s2 float64
		for j := Min(n, i+d); j >= i; j-- {
			s2 += dpx[j]
		}
		p2 := s2 / float64(d)

		fmt.Println(i, p1, p2)
		dpx[i] = math.Max(p1, p2)
	}
	fmt.Println(dpx)

	return dpx[0]
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
