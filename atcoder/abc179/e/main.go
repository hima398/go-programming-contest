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

	n, x, m := nextInt(), nextInt(), nextInt()

	ans := solve(n, x, m)

	PrintInt(ans)
}

func solveHonestly(n, x, m int) int {
	ai := x
	var ans int
	for i := 0; i < n; i++ {
		ans += ai
		ai = ai * ai % m
	}
	return ans
}

func solve(n, x, m int) int {
	//if n <= int(1e6) {
	//	return solveHonestly(n, x, m)
	//}

	// n > int(1e6)
	ms := make([]int, m+1)
	ai := x
	var loop []int
	for i := 0; ms[ai] < 3; i++ {
		if ms[ai] == 1 {
			loop = append(loop, ai)
		}
		ms[ai]++
		ai = ai * ai % m
	}

	//循環に入るまでの値を足す
	var ans int
	var f int
	for i, v := range ms {
		if v == 1 {
			f++
			ans += i
		}
	}

	//循環部分を足す
	var s int
	for _, v := range loop {
		s += v
	}
	rem := n - f
	w := rem / len(loop)
	ans += s * w

	//残りを足す
	rem = rem % len(loop)
	for i := 0; i < rem; i++ {
		ans += loop[i]
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
