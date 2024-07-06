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

	n, k := nextInt(), nextInt()
	a := nextIntSlice(k)

	ans := solve(n, k, a)

	Print(ans)
}

func solve(n, k int, a []int) int {
	if k%2 == 0 {
		var ans int
		for i := 0; i < k; i += 2 {
			ans += a[i+1] - a[i]
		}
		return ans
	} else { // k%2==1
		l := []int{0}
		for i := 0; i+1 < k; i += 2 {
			m := len(l)
			l = append(l, l[m-1]+a[i+1]-a[i])
		}
		r := []int{0}
		for i := k - 1; i-1 >= 0; i -= 2 {
			m := len(r)
			r = append(r, r[m-1]+a[i]-a[i-1])
		}
		ans := 1 << 60
		for i := 0; i < len(l); i++ {
			ans = Min(ans, l[i]+r[len(r)-i-1])
		}
		return ans
	}
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
