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

	n, m := nextInt(), nextInt()
	var a []string
	for i := 0; i < 2*n; i++ {
		a = append(a, nextString())
	}
	type player struct {
		i, w int
		s    string
	}
	ps := make([]player, 2*n)
	for i := 0; i < 2*n; i++ {
		ps[i].i = i + 1
		ps[i].s = a[i]
	}
	eval := func(a, b byte) int {
		if a == b {
			return 0
		}
		switch a {
		case 'G':
			if b == 'P' {
				return -1
			} else {
				return 1
			}
		case 'C':
			if b == 'G' {
				return -1
			} else {
				return 1
			}
		case 'P':
			if b == 'C' {
				return -1
			} else {
				return 1
			}
		}
		return -1
	}
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			v := eval(ps[2*i].s[j], ps[2*i+1].s[j])
			if v == 1 {
				ps[2*i].w++
			} else if v == -1 {
				ps[2*i+1].w++
			}
		}
		sort.Slice(ps, func(i, j int) bool {
			if ps[i].w == ps[j].w {
				return ps[i].i < ps[j].i
			}
			return ps[i].w > ps[j].w
		})
	}
	for _, p := range ps {
		Print(p.i)
	}

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
