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

	ln, lt := nextInt(), nextInt()
	var c, t []int
	for i := 0; i < ln; i++ {
		c = append(c, nextInt())
		t = append(t, nextInt())
	}
	ans := 1001
	for i := 0; i < ln; i++ {
		if t[i] <= lt {
			ans = Min(ans, c[i])
		}
	}
	if ans == 1001 {
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
