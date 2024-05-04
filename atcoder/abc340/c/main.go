package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/liyue201/gostl/ds/set"
	"github.com/liyue201/gostl/utils/comparator"
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
	Print(ans)
}

func solve(n int) int {
	bb := set.New[int](comparator.IntComparator)
	m := make(map[int]int)
	bb.Insert(n)
	m[n]++
	var ans int
	for it := bb.Last(); it.IsValid(); it.Prev() {
		x := it.Value()
		ans += x * m[x]
		y := Floor(x, 2)
		if y > 1 {
			bb.Insert(y)
			m[y] += m[x]
		}
		z := Ceil(x, 2)
		if z > 1 {
			bb.Insert(z)
			m[z] += m[x]
		}
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

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
