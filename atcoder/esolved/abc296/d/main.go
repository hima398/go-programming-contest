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

	n, m := nextInt(), nextInt()
	ans := solve(n, m)
	PrintInt(ans)
}

func solve(n, m int) int {
	const intSize = 32 << (^uint(0) >> 63) // 32 or 64
	const INF = 1<<(intSize-1) - 1

	ans := INF
	for a := 1; a <= n; a++ {
		b := Ceil(m, a)
		if b <= n {
			ans = Min(ans, a*b)
		}
		if a > b {
			break
		}
	}
	if ans == INF {
		return -1
	} else {
		return ans
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
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

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
