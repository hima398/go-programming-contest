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
	a := nextIntSlice(n)
	ans := solve(n, a)
	PrintInt(ans)
}

func dividePrime(x int) map[int]int {
	res := make(map[int]int)
	for i := 2; i*i <= x; i++ {
		for x%i == 0 {
			res[i]++
			x /= i
		}
		if x == 1 {
			break
		}
	}
	if x != 1 {
		res[x]++
	}
	return res
}

func solve(n int, a []int) int {
	const p = int(1e9) + 7
	//最小公倍数を求めるがオーバーフローするので対策する
	//最小公倍数をmapで、kのv乗で管理する
	pm := make(map[int]int)
	for _, ai := range a {
		for k, v := range dividePrime(ai) {
			pm[k] = Max(pm[k], v)
		}
	}
	//Ai*Bj=Sとする
	s := 1
	for k, v := range pm {
		s *= Pow(k, v, p)
		s %= p
	}
	//fmt.Println("s = ", s)
	var ans int
	for _, ai := range a {
		ans += s * Inv(ai, p)
		ans %= p
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
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
