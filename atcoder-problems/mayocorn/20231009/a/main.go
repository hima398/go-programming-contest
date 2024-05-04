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

	v := nextInt()
	a, b, c := nextInt(), nextInt(), nextInt()
	u := []int{a, b, c}
	ans := []string{"F", "M", "T"}
	idx := 0
	for {
		if v < u[idx] {
			Print(ans[idx])
			return
		}
		v -= u[idx]
		idx = (idx + 1) % 3
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
