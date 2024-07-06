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

	n, t := nextInt(), nextString()
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}

	ans := solve(n, t, s)

	Print(len(ans))
	for _, v := range ans {
		Print(v)
	}
}

func countDifferent(dest, src string) int {
	var res int
	for i := 0; i < len(dest); i++ {
		if dest[i] != src[i] {
			res++
		}
	}
	return res
}

func isAdded(dest, src string) bool {
	var i, j int
	for i < len(src) {
		if src[i] != dest[j] {
			if j-i >= 1 {
				return false
			} else {
				j++
				continue
			}
		}
		i++
		j++
	}
	return true
}

func solve(n int, t string, s []string) []int {
	var ans []int
	for i := 0; i < n; i++ {
		if len(t) == len(s[i]) {
			diff := countDifferent(s[i], t)
			if diff <= 1 {
				ans = append(ans, i+1)
			}
		} else if len(t)+1 == len(s[i]) {
			if isAdded(s[i], t) {
				ans = append(ans, i+1)
			}
		} else if len(t)-1 == len(s[i]) {
			if isAdded(t, s[i]) {
				ans = append(ans, i+1)
			}
		}
	}

	return ans
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
