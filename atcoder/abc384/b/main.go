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

	n, r := nextInt(), nextInt()
	var d, a []int
	for i := 0; i < n; i++ {
		d, a = append(d, nextInt()), append(a, nextInt())
	}

	ans := r
	for i := 0; i < n; i++ {
		switch d[i] {
		case 1:
			if 1600 <= ans && ans < 2800 {
				ans += a[i]
			}
		case 2:
			if 1200 <= ans && ans < 2400 {
				ans += a[i]
			}
		}
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
