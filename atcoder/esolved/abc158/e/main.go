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

	n, p := nextInt(), nextInt()
	s := nextString()

	ans := solve(n, p, s)

	PrintInt(ans)
}

func solve(n, p int, s string) int {
	if p == 2 || p == 5 {
		var ans int
		for i := 0; i < n; i++ {
			if int(s[i]-'0')%p == 0 {
				ans += i + 1
			}
		}
		return ans
	}

	m := make([]int, p)
	m[0]++
	w, v := 1, 0
	for i := n - 1; i >= 0; i-- {
		v += w * int(s[i]-'0')
		v %= p
		m[v]++

		w *= 10
		w %= p
	}
	//fmt.Println(m)
	var ans int
	for i := 0; i < p; i++ {
		ans += m[i] * (m[i] - 1) / 2
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
