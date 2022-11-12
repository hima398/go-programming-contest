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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var w []string
	for i := 0; i < n; i++ {
		w = append(w, nextString())
	}
	m := make(map[string]struct{})
	m[w[0]] = struct{}{}
	violate := func(prev, cur string) bool {
		//すでに使用している単語を使った
		if _, found := m[cur]; found {
			return true
		}
		if prev[len(prev)-1] != cur[0] {
			return true
		}
		return false
	}
	for i := 1; i < n; i++ {
		if i%2 == 1 {
			//高橋クンのターン
			if violate(w[i-1], w[i]) {
				PrintString("WIN")
				return
			}
		} else {
			//高橋くんのターン
			if violate(w[i-1], w[i]) {
				PrintString("LOSE")
				return
			}
		}
		m[w[i]] = struct{}{}
	}
	PrintString("DRAW")
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
