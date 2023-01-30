package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

//初項a、交差d、項数nの等差数列の和(mod p)
func arithmeticSum(a, d, n, p int) int {
	l := (a + (n-1+p)*d%p) % p
	s := (a + l) * n % p
	s = s * Inv(2, p) % p
	return s
}

func solve(n, m, q int, a, b, c, d []int) []int {
	const p = 998244353
	computeRowSum := func(x, l, r, p int) (int, int) {
		//行の左端、右端が奇数であれば補正して必ず0以上になるようにする
		if (x+l)%2 == 1 {
			l++
		}
		if (x+r)%2 == 1 {
			r--
		}
		if l > r {
			return 0, 0
		}

		a := ((x-1)*m + l) % p
		n := Ceil(r-l+1, 2)
		sum := arithmeticSum(a, 2, n, p)

		return sum, 2 * m * n % p
	}
	var ans []int
	for i := 0; i < q; i++ {
		a1, d1 := computeRowSum(a[i], c[i], d[i], p)
		a2, d2 := computeRowSum(a[i]+1, c[i], d[i], p)
		n1 := (b[i]-a[i]+1)/2 + (b[i]-a[i]+1)%2
		n2 := (b[i] - a[i] + 1) / 2
		v := arithmeticSum(a1, d1, n1, p) + arithmeticSum(a2, d2, n2, p)
		v %= p
		ans = append(ans, v)
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	q := nextInt()
	var a, b, c, d []int
	for i := 0; i < q; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
		c = append(c, nextInt())
		d = append(d, nextInt())
	}
	ans := solve(n, m, q, a, b, c, d)
	PrintVertically(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
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
