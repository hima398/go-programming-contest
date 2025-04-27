package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	q := nextInt()
	var c, d []string
	for i := 0; i < q; i++ {
		c = append(c, nextString())
		d = append(d, nextString())
	}
	ans := solve(n, s, q, c, d)
	Print(ans)
}

func solve(n int, s string, q int, c, d []string) string {
	m := make([]byte, 256)
	for j := 'a'; j <= 'z'; j++ {
		m[j] = byte(j)
	}
	for i := 0; i < q; i++ {
		for j := 'a'; j <= 'z'; j++ {
			if m[j] == c[i][0] {
				m[j] = d[i][0]
			}
		}
	}
	var ans []string
	for i := 0; i < n; i++ {
		ans = append(ans, string(m[s[i]]))
	}
	return strings.Join(ans, "")
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
