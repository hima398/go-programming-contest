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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	var s []string
	var t []string
	for i := 0; i < h; i++ {
		s = append(s, nextString())
	}
	for i := 0; i < h; i++ {
		t = append(t, nextString())
	}
	ans := solve(h, w, s, t)
	PrintString(ans)
}

func solve(h, w int, s, t []string) string {

	ms := make(map[string]int)
	mt := make(map[string]int)

	for j := 0; j < w; j++ {
		var col []string
		for i := 0; i < h; i++ {
			col = append(col, string(s[i][j]))
		}
		ms[strings.Join(col, "")]++
	}
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
	for k, v := range mt {
		ok = ok && ms[k] == v
	}
	if ok {
		return "Yes"
	} else {
		return "No"
	}
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
