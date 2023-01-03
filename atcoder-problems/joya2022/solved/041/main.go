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

	n, m := nextInt(), nextInt()
	var s []string
	var t []int
	for i := 0; i < n; i++ {
		s = append(s, nextString())
		v := 0
		for j := 0; j < m; j++ {
			if s[i][j] == 'o' {
				v |= 1 << j
			}
		}
		t = append(t, v)
	}

	var ans int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if t[i]|t[j] == 1<<m-1 {
				ans++
			}
		}
	}
	PrintInt(ans)
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
