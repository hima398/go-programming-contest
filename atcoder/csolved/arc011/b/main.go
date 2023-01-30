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
	var w []string
	for i := 0; i < n; i++ {
		w = append(w, nextString())
	}
	ans := solve(n, w)
	PrintHorizonaly(ans)
}

func solve(n int, w []string) []string {
	m := map[rune]string{'b': "1", 'c': "1", 'd': "2", 'w': "2", 't': "3", 'j': "3", 'f': "4", 'q': "4", 'l': "5", 'v': "5", 's': "6", 'x': "6", 'p': "7", 'm': "7", 'h': "8", 'k': "8", 'n': "9", 'g': "9", 'z': "0", 'r': "0"}
	convert := func(s string) string {
		res := ""
		for _, r := range s {
			if _, found := m[r]; found {
				res += m[r]
			}
		}

		return res

	}
	var ans []string
	for _, wi := range w {
		v := convert(strings.ToLower(wi))
		if v != "" {
			ans = append(ans, v)
		}
	}
	//fmt.Println(ans)
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

func PrintHorizonaly(x []string) {
	defer out.Flush()
	if len(x) == 0 {
		fmt.Fprintln(out)
		return
	}
	fmt.Fprintf(out, "%s", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %s", x[i])
	}
	fmt.Fprintln(out)
}
