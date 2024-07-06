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

	n, s, k := nextInt(), nextInt(), nextInt()
	var p, q []int
	for i := 0; i < n; i++ {
		p = append(p, nextInt())
		q = append(q, nextInt())
	}
	var ans int
	for i := 0; i < n; i++ {
		ans += p[i] * q[i]
	}
	if ans < s {
		ans += k
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
