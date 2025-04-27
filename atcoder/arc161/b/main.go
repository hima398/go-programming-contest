package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := nextInt()
	n := nextIntSlice(t)

	ans := solve(t, n)

	for _, v := range ans {
		Print(v)
	}
}

func solve(t int, n []int) []int {
	var f []int
	for i := 0; i <= 61; i++ {
		for j := i + 1; j <= 62; j++ {
			for k := j + 1; k <= 63; k++ {
				v := (1 << i) | (1 << j) | (1 << k)
				if v >= 0 && v <= int(1e18) {
					f = append(f, v)
				}
			}
		}
	}
	//fmt.Println("len(f) = ", len(f))
	sort.Ints(f)
	var ans []int
	for _, x := range n {
		idx := sort.Search(len(f), func(i int) bool {
			return x < f[i]
		})
		if idx == 0 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, f[idx-1])
		}
	}
	return ans
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
