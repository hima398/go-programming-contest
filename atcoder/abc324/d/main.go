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
	s := nextString()

	ans := solve(n, s)

	Print(ans)
}

func solve(n int, s string) int {
	var m [10]int
	for _, si := range s {
		m[si-'0']++
	}
	var ans int
	for i := 0; i < int(1e7); i++ {
		v := i * i
		var m1 [10]int
		for v > 0 {
			m1[v%10]++
			v /= 10
		}
		ok := m[0] >= m1[0]
		for j := 1; j <= 9; j++ {
			ok = ok && m[j] == m1[j]
		}
		if ok {
			ans++
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
