package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solveHonestly(n, m int) int {
	const p = int(1e9) + 7
	v := make(map[int]int)
	var f func(ai, s, i int)
	f = func(ai, s, i int) {
		ns := s ^ ai
		if i == n {
			v[ns]++
			return
		}
		for j := 0; j < 1<<m; j++ {
			f(j, ns, i+1)
		}
	}
	for i := 0; i < 1<<m; i++ {
		f(i, 0, 0)
	}
	fmt.Println(v)
	return 0
}

func solve(n, m int) int {
	const mod = int(1e9) + 7

	ans := Pow(2, m, mod)
	ans = (ans - 1 + mod) % mod
	ans = ans * Inv(2, mod) % mod
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	//ans := solveHonestly(n, m)
	ans := solve(n, m)
	PrintInt(ans)
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
