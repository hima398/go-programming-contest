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

	n := nextInt()
	s := nextString()

	//ans := solve01(n, s)
	ans := solve(n, s)

	PrintString(ans)
}

func solve(n int, s string) string {
	var t, a int
	for i := 0; i < n; i++ {
		if s[i] == 'T' {
			t++
		} else {
			a++
		}
	}
	if t > a {
		return "T"
	} else if t < a {
		return "A"
	} else {
		if s[n-1] == 'T' {
			return "A"
		} else {
			return "T"
		}
	}
}

func solve01(n int, s string) string {
	var t, a int
	var flag int
	for i := 0; i < n; i++ {
		if s[i] == 'T' {
			t++
			if t >= n/2 && flag == 0 {
				flag |= 1
			}
		} else {
			a++
			if a >= n/2 && flag == 0 {
				flag |= 2
			}
		}
	}
	if t > a {
		return "T"
	} else if t < a {
		return "A"
	} else {
		if flag == 1 {
			return "T"
		} else {
			return "A"
		}
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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
