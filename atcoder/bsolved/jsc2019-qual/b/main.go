package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func solve(n, k int, a []int) (ans int) {
	const p = int(1e9) + 7
	//t1:aの中での転倒数
	//t2:aの中でa[i]より大きい数
	t1, t2 := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if a[j] > a[i] {
				t1[i]++
			}
		}
		for j := 0; j < n; j++ {
			if i == j {
				continue
			} else if a[i] < a[j] {
				t2[i]++
			}
		}
	}
	var st1, st2 int
	for i := 0; i < n; i++ {
		st1 += t1[i]
		st2 += t2[i]
	}
	w := k * (k - 1) % p
	w *= Inv(2, p)
	w %= p
	ans = (k*st1 + w*st2) % p
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	a := nextIntSlice(n)
	ans := solve(n, k, a)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
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
