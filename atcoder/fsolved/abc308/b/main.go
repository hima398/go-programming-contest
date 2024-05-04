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

	n, m := nextInt(), nextInt()
	var c []string
	for i := 0; i < n; i++ {
		c = append(c, nextString())
	}
	var d []string
	for i := 0; i < m; i++ {
		d = append(d, nextString())
	}
	p := nextIntSlice(m + 1)

	price := make(map[string]int)
	for i := 0; i < m; i++ {
		price[d[i]] = p[i+1]
	}
	//fmt.Println(fee)
	var ans int
	for _, ci := range c {
		if v, found := price[ci]; found {
			ans += v
		} else {
			ans += p[0]
		}
	}
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
