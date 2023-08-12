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

	inputSheet := func() (int, int, []string) {
		h, w := nextInt(), nextInt()
		var s []string
		for i := 0; i < h; i++ {
			s = append(s, nextString())
		}
		return h, w, s
	}
	ha, wa, a := inputSheet()
	hb, wb, b := inputSheet()
	hx, wx, x := inputSheet()

	ok := solve(ha, wa, a, hb, wb, b, hx, wx, x)

	if ok {
		PrintString("Yes")
	} else {
		PrintString("No")
	}
}

func solve(ha, wa int, a []string, hb, wb int, b []string, hx, wx int, x []string) bool {
	var nx int
	for i := 0; i < hx; i++ {
		for j := 0; j < wx; j++ {
			if x[i][j] == '#' {
				nx++
			}
		}
	}
	for i := 0; i < 50-hb; i++ {
		for j := 0; j < 50-wb; j++ {
			var s [50][50]bool
			for ai := 0; ai < ha; ai++ {
				for aj := 0; aj < wa; aj++ {
					if a[ai][aj] == '#' {
						s[20+ai][20+aj] = true
					}
				}
			}
			for bi := 0; bi < hb; bi++ {
				for bj := 0; bj < wb; bj++ {
					if b[bi][bj] == '#' {
						s[i+bi][j+bj] = true
					}
				}
			}
			var ns int
			for si := 0; si < 50; si++ {
				for sj := 0; sj < 50; sj++ {
					if s[si][sj] {
						ns++
					}
				}
			}
			if nx != ns {
				continue
			}
			for ii := 0; ii < 50-hx; ii++ {
				for jj := 0; jj < 50-wx; jj++ {
					ok := true
					for xi := 0; xi < hx; xi++ {
						for xj := 0; xj < wx; xj++ {
							ok = ok && s[ii+xi][jj+xj] == (x[xi][xj] == '#')
						}
					}
					if ok {
						return true
					}
				}
			}
		}
	}
	return false
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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
