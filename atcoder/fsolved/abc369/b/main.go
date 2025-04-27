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

	n := nextInt()
	var a []int
	var s []string
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		s = append(s, nextString())
	}

	const NumKeys = 100
	ans := math.MaxInt
	for l := 1; l <= NumKeys; l++ {
		for r := 1; r <= NumKeys; r++ {
			cl, cr := l, r
			var fatigue int
			for i := 0; i < n; i++ {
				switch s[i] {
				case "L":
					fatigue += Abs(a[i] - cl)
					cl = a[i]
				case "R":
					fatigue += Abs(a[i] - cr)
					cr = a[i]
				}
			}
			ans = Min(ans, fatigue)
		}
	}

	Print(ans)
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
