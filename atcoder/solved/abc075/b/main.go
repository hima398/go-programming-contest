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
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = nextString()
	}
	ans := make([][]string, h)
	for i := 0; i < h; i++ {
		ans[i] = make([]string, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				ans[i][j] = "#"
				continue
			}
			v := 0
			for ii := i - 1; ii <= i+1; ii++ {
				if ii < 0 || ii >= h {
					continue
				}
				for jj := j - 1; jj <= j+1; jj++ {
					if jj < 0 || jj >= w {
						continue
					}
					if s[ii][jj] == '#' {
						v++
					}
				}
			}
			ans[i][j] = strconv.Itoa(v)
		}
	}
	for i := 0; i < h; i++ {
		PrintString(strings.Join(ans[i], ""))
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
