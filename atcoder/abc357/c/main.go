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

	n := nextInt()
	ans := solve(n)
	for _, v := range ans {
		Print(v)
	}
}

func makeCarpet(carpet [][]string, part [][]string, si, sj int) {
	n := len(part)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			carpet[si+i][sj+j] = part[i][j]
		}
	}
}

func solve(n int) []string {
	carpet := [][]string{[]string{"#"}}
	for k := 0; k < n; k++ {
		next := make([][]string, Pow(3, k+1))
		for i := range next {
			next[i] = make([]string, Pow(3, k+1))
			for j := range next[i] {
				next[i][j] = "."
			}
		}
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if i == 1 && j == 1 {
					continue
				}
				makeCarpet(next, carpet, i*Pow(3, k), j*Pow(3, k))
			}
		}
		carpet = next
	}
	var ans []string
	for _, v := range carpet {
		ans = append(ans, strings.Join(v, ""))
	}
	return ans
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

func Pow(x, y int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x
		}
		y >>= 1
		x = x * x
	}
	return ret
}
