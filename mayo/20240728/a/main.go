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

	const INF = math.MaxInt
	n, t := nextInt(), nextInt()
	var c, ts []int
	for i := 0; i < n; i++ {
		c = append(c, nextInt())
		ts = append(ts, nextInt())
	}

	ans := INF
	for i := 0; i < n; i++ {
		if ts[i] <= t {
			ans = Min(ans, c[i])
		}
	}
	if ans == INF {
		Print("TLE")
	} else {
		Print(ans)
	}
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
