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

	q := nextInt()
	t, x := make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = nextInt()
		if t[i] == 1 {
			x[i] = nextInt()
		}
	}
	ans := solve(q, t, x)
	PrintVertically(ans)
}

func solve(q int, t, x []int) (ans []int) {
	const p = 998244353
	s := []int{1}
	v := 1
	w := make([]int, q+1)
	w[0] = 1
	for k := 1; k <= q; k++ {
		w[k] = (10 * w[k-1]) % p
	}
	for k := 0; k < q; k++ {
		switch t[k] {
		case 1:
			s = append(s, x[k])
			v = (10*v + x[k]) % p
		case 2:
			idx := len(s) - 1
			v = (v - s[0]*w[idx]%p + p) % p
			s = s[1:]
		case 3:
			ans = append(ans, v)
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}

func Inv(x, p int) int {
	return Pow(x, p-2, p)
}
