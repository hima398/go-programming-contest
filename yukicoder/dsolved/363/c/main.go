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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	s := make([]string, 2*n)
	for i := 0; i < 2*n; i++ {
		s[i] = nextString()
	}
	c := make([][]int, 2*n)
	for i := 0; i < 2*n; i++ {
		c[i] = nextIntSlice(2 * n)
	}
	ans := solve(n, k, s, c)
	PrintInt(ans)
}

func solve(n, k int, s []string, c [][]int) int {
	//Ci, jが大きい順に黒くする
	nb := 0
	var cs []int
	for i := 0; i < 2*n; i++ {
		for j := 0; j < 2*n; j++ {
			if s[i][j] == '#' {
				nb++
			}
			cs = append(cs, c[i][j])
		}
	}
	sort.Slice(cs, func(i, j int) bool {
		return cs[i] > cs[j]
	})
	ans := 0
	for i := 0; i < nb; i++ {
		ans += cs[i]
	}
	//シンメトリーが作れないのでここで終了
	if nb%2 == 1 {
		return ans
	}
	//シンメトリーのスコアと比較する
	var hcs []int
	for i := 0; i < 2*n; i++ {
		for j := 0; j < n; j++ {
			hcs = append(hcs, c[i][j]+c[i][2*n-1-j])
		}
	}
	ans2 := k
	sort.Slice(hcs, func(i, j int) bool {
		return hcs[i] > hcs[j]
	})
	for i := 0; i < nb/2; i++ {
		ans2 += hcs[i]
	}
	return Max(ans, ans2)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
