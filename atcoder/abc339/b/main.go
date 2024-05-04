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

	h, w, n := nextInt(), nextInt(), nextInt()
	var ans [][]string
	for i := 0; i < h; i++ {
		ans = append(ans, strings.Split(strings.Repeat(".", w), ""))
	}
	di := []int{-1, 0, 1, 0}
	dj := []int{0, 1, 0, -1}
	ci, cj := 0, 0
	dir := 0
	for i := 0; i < n; i++ {
		switch ans[ci][cj] {
		case ".":
			ans[ci][cj] = "#"
			dir = (dir + 1) % 4
		case "#":
			ans[ci][cj] = "."
			dir = (dir + 3) % 4
		}
		ci, cj = (ci+di[dir]+h)%h, (cj+dj[dir]+w)%w
	}

	for _, v := range ans {
		Print(strings.Join(v, ""))
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
