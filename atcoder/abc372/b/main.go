package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func Pow(x, y int) int {
	/*
		res := 1
		for i := 0; i < y; i++ {
			res *= x
		}
		return res
	*/
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x
		}
		y >>= 1
		x = x * x
	}
	return ret
}

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	m := nextInt()

	pows := make([]int, 11)
	pows[0] = 1
	for i := 1; i <= 10; i++ {
		pows[i] = 3 * pows[i-1]
	}
	var ans []int
	for i := 10; i >= 0; i-- {
		for m >= pows[i] {
			m -= pows[i]
			ans = append(ans, i)
		}
	}
	Print(len(ans))
	PrintHorizonaly(ans)
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
