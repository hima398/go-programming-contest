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

	n, l := nextInt(), nextInt()
	t := n*l + n*(n-1)/2

	diff := 1 << 60
	var ans int
	for i := 0; i < n; i++ {
		var t2 int
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			t2 += l + j
		}
		if diff > Abs(t-t2) {
			diff = Abs(t - t2)
			ans = t2
		}
	}
	PrintInt(ans)
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
