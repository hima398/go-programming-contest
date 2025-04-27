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

	h, w := nextInt(), nextInt()
	var s []string
	for i := 0; i < h; i++ {
		s = append(s, nextString())
	}
	var t []string
	for i := 0; i < h; i++ {
		t = append(t, nextString())
	}

	ok := solve(h, w, s, t)

	if ok {
		Print("Yes")
	} else {
		Print("No")
	}
}

func solve(h, w int, s, t []string) bool {
	ms := make(map[string]int)

	for j := 0; j < w; j++ {
		var col []string
		for i := 0; i < h; i++ {
			col = append(col, string(s[i][j]))
		}
		ms[strings.Join(col, "")]++
	}

	mt := make(map[string]int)
	for j := 0; j < w; j++ {
		var col []string
		for i := 0; i < h; i++ {
			col = append(col, string(t[i][j]))
		}
		mt[strings.Join(col, "")]++
	}
	ok := true
	for k, v := range ms {
		ok = ok && mt[k] == v
	}
	return ok
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
