package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func solve(n, m, k int) int {
	const p = 998244353
	var ans int
	if n == 1 {
		return Pow(k, m, p)
	}
	if m == 1 {
		return Pow(k, n, p)
	}
	for i := 1; i <= k; i++ {
		h := Pow(i, n, p) - Pow((i-1), n, p) + p
		h %= p
		w := Pow((k - i + 1), m, p)
		//fmt.Println(h, w)
		ans += h * w % p
		ans %= p
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, k := nextInt(), nextInt(), nextInt()
	ans := solve(n, m, k)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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
