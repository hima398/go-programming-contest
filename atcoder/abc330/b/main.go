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

	n, l, r := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n)

	var ans []int
	for _, ai := range a {
		if ai <= l {
			ans = append(ans, l)
		} else if ai >= r {
			ans = append(ans, r)
		} else {
			ans = append(ans, ai)
		}
	}
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintf(out, "%d ", v)
	}
	fmt.Fprintln(out)
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
