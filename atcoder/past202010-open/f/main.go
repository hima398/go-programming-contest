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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}
	ans := solve(n, k, s)
	PrintString(ans)
}

func solve(n, k int, s []string) string {
	m := make(map[string]int)
	for _, si := range s {
		m[si]++
	}
	//i回出現する単語がいくつあるかを数える
	ns := make([]int, int(1e5)+1)
	type word struct {
		x int    //出現回数
		s string //単語
	}
	var ws []word
	for k, v := range m {
		ws = append(ws, word{v, k})
		ns[v]++
	}
	sort.Slice(ws, func(i, j int) bool {
		return ws[i].x > ws[j].x
	})
	candidate := ws[k-1]
	if ns[candidate.x] > 1 {
		return "AMBIGUOUS"
	} else {
		return candidate.s
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
