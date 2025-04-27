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

	n, k := nextInt(), nextInt()
	s := nextString()

	ans := solve(n, k, s)

	Print(ans)
}

type node struct {
	n int
	c rune
}

func RunLengthEncoding(s string) []node {
	var res []node
	for _, v := range s {
		if len(res) == 0 || res[len(res)-1].c != v {
			res = append(res, node{1, v})
		} else {
			res[len(res)-1].n++
		}
	}
	return res
}

func solve(n, k int, s string) string {
	//ss := strings.Split(s, "0")
	ss := RunLengthEncoding(s)
	//fmt.Println(ss)
	var ans []string
	var cnt int //解答につなげた1の塊の個数
	var t []string

	for _, v := range ss {
		if cnt < k-1 {
			ans = append(ans, strings.Repeat(string(v.c), v.n))
		} else if cnt == k-1 {
			if v.c == '0' {
				t = append(t, strings.Repeat(string(v.c), v.n))
			} else {
				ans = append(ans, strings.Repeat(string(v.c), v.n))
			}
		} else {
			// cnt>k-1
			t = append(t, strings.Repeat(string(v.c), v.n))
		}
		if v.c == '1' {
			cnt++
		}
	}
	ans = append(ans, t...)

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
