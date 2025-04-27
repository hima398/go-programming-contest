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

	h, w := nextInt(), nextInt()
	var s []string
	for i := 0; i < h; i++ {
		s = append(s, nextString())
	}

	ans := solve(h, w, s)

	for _, v := range ans {
		fmt.Println(v[0], v[1])
	}

}

func solve(h, w int, s []string) [][2]int {
	const t = "snuke"

	di := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dj := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for k := 0; k < 8; k++ {
				var a []string
				var ans [][2]int
				ni, nj := i, j
				for l := 0; l < len(t); l++ {
					if ni < 0 || ni >= h || nj < 0 || nj >= w {
						break
					}
					a = append(a, string(s[ni][nj]))
					ans = append(ans, [2]int{ni + 1, nj + 1})
					ni, nj = ni+di[k], nj+dj[k]
				}
				if len(a) == len(t) && strings.Join(a, "") == t {
					return ans
				}
			}
		}
	}
	return nil
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
