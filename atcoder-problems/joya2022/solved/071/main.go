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
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}
	ans := solve(n, s)
	PrintString(ans)
}

func solve(n int, s []string) string {
	const size = 6
	for k := 0; k < 2; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n-size+1; j++ {
				black := 0
				for k := 0; k < size; k++ {
					if s[i][j+k] == '#' {
						black++
					}
				}
				if black >= size-2 {
					return "Yes"
				}
			}
		}
		for i := 0; i < n-size+1; i++ {
			for j := 0; j < n-size+1; j++ {
				black := 0
				for k := 0; k < size; k++ {
					if s[i+k][j+k] == '#' {
						black++
					}
				}
				if black >= size-2 {
					return "Yes"
				}
			}
		}

		t := make([][]rune, n)
		for i := range t {
			t[i] = make([]rune, n)
		}
		for i := 0; i < n; i++ {
			for j := n - 1; j >= 0; j-- {
				t[i][n-1-j] = rune(s[j][i])
			}
		}
		for i := 0; i < n; i++ {
			s[i] = string(t[i])
		}
	}
	return "No"
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
