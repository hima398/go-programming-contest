package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()

	ans := solve(s)

	Print(ans)
}

func solve(s string) int {
	/*
		var m [26][]int
		for i, si := range s {
			m[int(si-'A')] = append(m[int(si-'A')], i)
		}

		fmt.Println(m)
	*/
	n := len(s)

	m := make([][26]int, n)
	for i := 0; i < n; i++ {
		m[i][int(s[i]-'A')]++
	}
	//iより左にある文字の総和、iより右にある文字の総和
	l, r := make([][26]int, n), make([][26]int, n)
	for i := 0; i < n-1; i++ {
		for j := 0; j < 26; j++ {
			l[i+1][j] = l[i][j] + m[i][j]
		}
	}
	for i := n - 1; i > 0; i-- {
		for j := 0; j < 26; j++ {
			r[i-1][j] = r[i][j] + m[i][j]
		}
	}
	var ans int
	for i := 1; i < n-1; i++ {
		for j := 0; j < 26; j++ {
			ans += l[i][j] * r[i][j]
		}
	}
	return ans
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
