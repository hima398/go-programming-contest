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

	n, k := nextInt(), nextInt()
	ans := solve(n, k)
	PrintInt(ans)
}

func solve(n, k int) int {
	const p = int(1e9) + 7

	var ans int
	for i := k; i <= n+1; i++ {
		//0〜(i-1)までの和
		l := i * (i - 1) / 2
		//n〜(n-i)までの和
		r := (2*n - i + 1) * i / 2
		ans += r - l + 1
		//fmt.Println(i, r-l+1)
		ans %= p
	}
	return ans
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
