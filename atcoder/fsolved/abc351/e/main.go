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

	n := nextInt()
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	ans := solve(n, x, y)
	Print(ans)
}

func solve(n int, x, y []int) int {
	us := make([][]int, 2)
	vs := make([][]int, 2)
	for i := 0; i < n; i++ {
		u, v := x[i]-y[i], x[i]+y[i]
		k := (x[i] + y[i]) % 2
		us[k] = append(us[k], u)
		vs[k] = append(vs[k], v)
	}
	var ans int
	for k := 0; k < 2; k++ {
		sort.Ints(us[k])
		m := len(us[k])
		for i, u := range us[k] {
			w := 2*i - m + 1
			ans += u * w
		}

		sort.Ints(vs[k])
		l := len(vs[k])
		for i, v := range vs[k] {
			w := 2*i - l + 1
			ans += v * w
		}
	}
	ans /= 2
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
