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

	n, x := nextInt(), nextInt()
	var u, d []int
	for i := 0; i < n; i++ {
		u, d = append(u, nextInt()), append(d, nextInt())
	}

	ans := solve(n, x, u, d)

	Print(ans)
}

func solve(n, x int, u, d []int) int {
	var sum int
	for i := 0; i < n; i++ {
		sum += u[i] + d[i]
	}
	ok, ng := 0, 1<<60
	for (ng - ok) > 1 {
		mid := (ok + ng) / 2
		if check(mid, n, x, u, d) {
			ok = mid
		} else {
			ng = mid
		}
	}
	ans := sum - n*ok
	return ans
}

// 高さcandidateとなる噛み合わせを作ろうとする時、条件を満たすu_iが存在するかをチェックする
func check(candidate, n, x int, u, d []int) bool {
	low, high := Max(candidate-d[0], 1), Min(u[0], candidate-1)
	if low > high {
		return false
	}
	for i := 1; i < n; i++ {
		nextLow := Max(candidate-d[i], 1)
		nextHigh := Min(u[i], candidate-1)
		if nextLow > nextHigh {
			return false
		}
		low = Max(low-x, nextLow)
		high = Min(high+x, nextHigh)
		if low > high {
			return false
		}
	}
	return true
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
