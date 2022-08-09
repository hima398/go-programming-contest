package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solveHonestly(n, c int, t, a []int) (ans []int) {
	cc := c
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			switch t[j] {
			case 1:
				cc &= a[j]
			case 2:
				cc |= a[j]
			case 3:
				cc ^= a[j]
			}
		}
		ans = append(ans, cc)
	}
	return ans
}

func solve(n, c int, t, a []int) []int {
	const maxBits = 30

	var andBits, orBits, xorBits [maxBits]int

	for j := 0; j < maxBits; j++ {
		andBits[j] = -1
		orBits[j] = -1
	}

	x := []int{c}
	for i := 0; i < n; i++ {
		//Aiの各桁を見て影響する値の更新
		for j := 0; j < maxBits; j++ {
			switch t[i] {
			case 1:
				if a[i]>>j&1 == 0 {
					andBits[j] = i
					xorBits[j] = 0
				}
			case 2:
				if a[i]>>j&1 > 0 {
					orBits[j] = i
					xorBits[j] = 0
				}
			case 3:
				if a[i]>>j&1 > 0 {
					xorBits[j]++
				}
			}
		}

		//影響する値とxorを見て次のXを計算
		var ncs [maxBits]int
		nextX := 0
		for j := 0; j < maxBits; j++ {
			if andBits[j] < 0 && orBits[j] < 0 {
				ncs[j] = x[i] >> j & 1
			} else if andBits[j] < orBits[j] {
				ncs[j] = 1
			} else if andBits[j] > orBits[j] {
				ncs[j] = 0
			}
			ncs[j] = (ncs[j] + xorBits[j]) % 2
			if ncs[j] == 1 {
				nextX |= 1 << j
			}
		}

		x = append(x, nextX)
	}

	// X0を除いて、X1~Xnまでを解答として返す
	return x[1:]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, c := nextInt(), nextInt()
	var t, a []int
	for i := 0; i < n; i++ {
		t = append(t, nextInt())
		a = append(a, nextInt())
	}
	ans := solve(n, c, t, a)
	PrintVertically(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
