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

func computeDist(x1, y1, x2, y2 int) float64 {
	xx, yy := x2-x1, y2-y1
	return math.Sqrt(float64(xx*xx + yy*yy))
}
func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	x, y := []int{0}, []int{0}
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	x, y = append(x, 0), append(y, 0)
	var ans float64
	for i := 0; i < len(x)-1; i++ {
		ans += computeDist(x[i], y[i], x[i+1], y[i+1])
	}

	Print(ans)
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
