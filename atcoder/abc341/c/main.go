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

	h, w, n := nextInt(), nextInt(), nextInt()
	t := nextString()
	var s []string
	for i := 0; i < h; i++ {
		s = append(s, nextString())
	}

	ans := solve(h, w, n, t, s)

	Print(ans)
}

func simulateSpaceship(i, j int, t string, s []string) bool {
	if s[i][j] == '#' {
		return false
	}

	dir := map[rune][]int{'L': {0, -1}, 'R': {0, 1}, 'U': {-1, 0}, 'D': {1, 0}}
	ci, cj := i, j
	for _, ti := range t {
		ci, cj = ci+dir[ti][0], cj+dir[ti][1]
		if s[ci][cj] == '#' {
			return false
		}
	}
	return true
}

func solve(h, w, n int, t string, s []string) int {
	var ans int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if simulateSpaceship(i, j, t, s) {
				ans++
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
