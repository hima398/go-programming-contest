package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var s []string
	var c []int
	for i := 0; i < n; i++ {
		s = append(s, nextString())
		c = append(c, nextInt())
	}
	type user struct {
		s string
		c int
	}
	var us []user
	var t int
	for i := 0; i < n; i++ {
		us = append(us, user{s[i], c[i]})
		t += c[i]
	}
	sort.Slice(us, func(i, j int) bool {
		return us[i].s < us[j].s
	})
	ans := us[t%n].s
	Print(ans)
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
