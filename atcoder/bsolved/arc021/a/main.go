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

	const n = 4
	var a [n][n]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = nextInt()
		}
	}
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := range di {
				ni, nj := i+di[k], j+dj[k]
				if ni < 0 || ni >= n || nj < 0 || nj >= n {
					continue
				}
				if a[i][j] == a[ni][nj] {
					PrintString("CONTINUE")
					return
				}
			}
		}
	}
	PrintString("GAMEOVER")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
