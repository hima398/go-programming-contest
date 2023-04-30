package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()

	ans := solve(s)

	PrintInt(ans)
}

func solve(s string) int {
	const p = 2019
	n := len(s)
	//m := make(map[int]int)
	m := make([]int, p)
	m[0] = 1
	sum := 0
	w := 1
	for i := n - 1; i >= 0; i-- {
		v := int(s[i] - '0')
		sum = w*v + sum
		sum %= p
		m[sum]++
		//次の計算のために重みを更新
		w *= 10
		w %= p
	}
	//fmt.Println(m)
	var ans int
	for _, v := range m {
		ans += v * (v - 1) / 2
	}
	return ans
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
