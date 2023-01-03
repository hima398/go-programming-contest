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

	n, a, b := nextInt(), nextInt(), nextInt()
	ans := make([][]rune, n*a)
	for i := 0; i < n*a; i++ {
		ans[i] = make([]rune, n*b)
		for j := 0; j < n*b; j++ {
			ans[i][j] = '.'
		}
	}
	hIsWhite := 1
	for i := 0; i < n*a; i += a {
		wIsWhite := hIsWhite
		for j := 0; j < n*b; j += b {
			if wIsWhite == 0 {
				for ii := 0; ii < a; ii++ {
					for jj := 0; jj < b; jj++ {
						ans[i+ii][j+jj] = '#'
					}
				}
			}
			wIsWhite ^= 1
		}
		hIsWhite ^= 1
	}
	for i := 0; i < n*a; i++ {
		PrintString(string(ans[i]))
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
