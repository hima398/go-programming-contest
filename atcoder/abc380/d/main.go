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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	q := nextInt()
	k := nextIntSlice(q)

	ans := solve(s, q, k)

	Print(ans)
}

func solve(s string, q int, k []int) string {
	var f func(x, l, cnt int) (int, int)
	f = func(x, l, cnt int) (int, int) {
		if x < len(s) {
			return x, cnt
		}
		if x < l/2 {
			return f(x, l/2, cnt)
		} else {
			// l/2 <= x
			return f(x-l/2, l/2, cnt+1)
		}
	}

	var ans []string
	for _, ki := range k {
		ki--
		//ki以上のPow(2, x)を計算しておく
		l := len(s)
		for l <= ki {
			l <<= 1
		}
		idx, cnt := f(ki, l, 0)
		//fmt.Println(ki, l, idx, cnt)

		v := string(s[idx])
		if cnt%2 == 1 {
			if "A" <= v && v <= "Z" {
				v = strings.ToLower(v)
			} else if "a" <= v && v <= "z" {
				v = strings.ToUpper(v)
			}
		}
		ans = append(ans, v)
	}
	return strings.Join(ans, " ")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
