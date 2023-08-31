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

	n := nextInt()

	ans := solve(n)

	PrintInt(ans)
}

func solve(n int) int {
	q := []int{3, 5, 7}
	is753 := func(x int) bool {
		var m [8]bool
		for x > 0 {
			m[x%10] = true
			x /= 10
		}
		return m[3] && m[5] && m[7]
	}
	var ans int
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur > n {
			break
		}
		if is753(cur) {
			ans++
		}
		q = append(q, cur*10+3)
		q = append(q, cur*10+5)
		q = append(q, cur*10+7)
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
