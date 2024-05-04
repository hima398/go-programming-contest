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

	_, q := nextInt(), nextInt()
	var t, a, b []int
	for i := 0; i < q; i++ {
		t = append(t, nextInt())
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	var ans []bool
	m := make(map[int]map[int]struct{})
	for i := range t {
		switch t[i] {
		case 1:
			if m[a[i]] == nil {
				m[a[i]] = make(map[int]struct{})
			}
			m[a[i]][b[i]] = struct{}{}
		case 2:
			if m[a[i]] == nil {
				m[a[i]] = make(map[int]struct{})
			}
			delete(m[a[i]], b[i])
		case 3:
			if _, found1 := m[a[i]][b[i]]; found1 {
				_, found2 := m[b[i]][a[i]]
				ans = append(ans, found2)
			} else {
				ans = append(ans, false)
			}
		}
	}

	for _, ok := range ans {
		if ok {
			Print("Yes")
		} else {
			Print("No")
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
