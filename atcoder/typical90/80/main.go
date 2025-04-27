package main

import (
	"bufio"
	"fmt"
	"math/bits"
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

	n, d := nextInt(), nextInt()
	a := nextIntSlice(n)

	//ans := solveHonestly(n, d, a)
	ans := solve(n, d, a)

	Print(ans)
}

func solveHonestly(n, d int, a []int) int {
	var ans int
	for x := 0; x < 1<<d; x++ {
		ok := true
		for _, v := range a {
			ok = ok && x&v > 0
		}
		if ok {
			ans++
		}
	}
	return ans
}

func solve(n, d int, a []int) int {
	ans := 0
	for pat := 0; pat < 1<<n; pat++ {
		var t int
		for i := 0; i < n; i++ {
			if (pat>>i)&1 == 0 {
				continue
			}
			t |= a[i]
		}
		var cnt int
		for i := 0; i < d; i++ {
			if (t>>i)&1 == 0 {
				cnt++
			}
		}
		if bits.OnesCount(uint(pat))%2 == 0 {
			ans += (1 << cnt)
		} else {
			ans -= (1 << cnt)
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
