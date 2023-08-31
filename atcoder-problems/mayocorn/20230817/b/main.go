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

	n := nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	var t float64
	for i := 0; i < n; i++ {
		t += float64(a[i]) / float64(b[i])
	}
	t /= 2.0
	var ans float64
	for i := 0; i < n; i++ {
		c := float64(a[i]) / float64(b[i])
		if t > c {
			ans += float64(a[i])
			t -= c
		} else {
			ans += float64(b[i]) * t
			break
		}
	}
	PrintFloat64(ans)
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
