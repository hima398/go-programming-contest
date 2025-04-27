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

	r := nextInt()
	ans := solve(r)
	Print(ans)
}

func solve(r int) int {
	var ans int
	for x := 1; x < 2*r; x += 2 {
		y := sort.Search(2*r, func(i int) bool {
			return i*i > 4*r*r-x*x
		})
		y = y/2 - 1
		//fmt.Printf("(x, y) = (%d, %d)\n", x, y)
		ans += y
	}
	ans *= 4
	ans++
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
