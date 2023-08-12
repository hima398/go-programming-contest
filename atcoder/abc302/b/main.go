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
	const t = "snuke"

	di := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dj := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] != 's' {
				continue
			}
			for k := 0; k < 8; k++ {
				var a []string
				var ans [][2]int
				ni, nj := i, j
				for l := 0; l < len(t); l++ {
					if ni < 0 || ni >= h || nj < 0 || nj >= w {
						break
					}
					a = append(a, string(s[ni][nj]))
					ans = append(ans, [2]int{ni, nj})
					ni, nj = ni+di[k], nj+dj[k]
				}
				if len(a) == 5 && strings.Join(a, "") == t {
					for _, v := range ans {
						fmt.Println(v[0]+1, v[1]+1)
					}
					return
				}
			}
		}
	}
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

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
