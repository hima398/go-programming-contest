package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func convert(s string) string {
	switch s {
	case "d":
		return "p"
	case "p":
		return "d"
	}
	return ""
}

func solve(n int, s string) string {
	l := -1
	for i, r := range s {
		if r == 'p' {
			l = i
			break
		}
	}
	//全てdなので何も操作しなくて良い
	if l < 0 {
		return s
	}
	ans := []string{s}
	ss := strings.Split(s, "")
	for r := l; r < n; r++ {
		ts := make([]string, n)
		copy(ts, ss)
		for i := l; i <= r; i++ {
			ts[i] = convert(ts[i])
		}
		for i := 0; i <= (r-l)/2; i++ {
			j := r - l - i
			ts[i+l], ts[j+l] = ts[j+l], ts[i+l]
		}
		ans = append(ans, strings.Join(ts, ""))
	}

	sort.Strings(ans)
	//fmt.Println(ans)
	return ans[0]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := nextString()
	ans := solve(n, s)
	PrintString(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
