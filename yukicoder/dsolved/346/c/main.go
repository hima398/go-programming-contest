package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func divide(x int) []int {
	m := make(map[int]struct{})
	for i := 1; i*i <= x; i++ {
		if x%i == 0 {
			m[i] = struct{}{}
			m[x/i] = struct{}{}
		}
	}
	var res []int
	for k := range m {
		res = append(res, k)
	}
	sort.Ints(res)
	return res
}

func solveHonestly(t int, n, m []int) (ans []int) {
	const p = int(1e9) + 7

	nm := Pow(n[0], m[0], p)
	var s int
	for i := 1; i <= nm; i++ {
		d := divide(i)
		if len(d)%2 == 0 {
			idx := len(d)/2 - 1
			mid := d[idx] + d[idx+1]
			if mid*mid > i*4 {
				s += i
			} else {
				fmt.Println(len(d), i)
			}
		} else {
			mid := d[(len(d)+1)/2-1]
			if mid*mid > i {
				s += i
			} else {
				fmt.Println(len(d), i)
			}
		}
	}
	ans = append(ans, s)
	return ans
}

func solve(t int, n, m []int) (ans []int) {
	const p = int(1e9) + 7

	inv2 := Inv(2, p)
	inv6 := Inv(6, p)
	for i := 0; i < t; i++ {

		nm := Pow(n[i], m[i], p)
		v := nm * (nm + 1) % p
		v = v * inv2 % p

		nm2 := Pow(n[i], m[i]/2, p)
		mi := nm2 * (nm2 + 1) % p
		mi = mi * (2*nm2 + 1) % p
		mi = mi * inv6 % p
		v = (v - mi + p) % p
		ans = append(ans, v)
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := nextInt()
	var n, m []int
	for i := 0; i < t; i++ {
		n = append(n, nextInt())
		m = append(m, nextInt())
	}
	ans := solve(t, n, m)
	//ans := solveHonestly(t, n, m)
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
