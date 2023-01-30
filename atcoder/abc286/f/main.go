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

	n := solve()

	PrintInt(n)
}

func solve() int {
	//ps := []int{4, 5}
	ps := []int{4, 5, 7, 9, 11, 13, 17, 19, 23}

	var a []int
	i := 0
	for _, p := range ps {
		for j := 0; j < p; j++ {
			a = append(a, i+(j+1)%p+1)
		}
		i += p
	}

	//psの各値の和と同じになるはず
	m := len(a)
	PrintInt(m)
	PrintHorizonaly(a)

	b := nextIntSlice(m)
	for i := range b {
		b[i]--
	}
	idx := 0
	rm := make(map[int]int)
	for _, p := range ps {
		rm[p] = (b[idx] - idx + p) % p
		idx += p
	}
	//fmt.Println(rm)
	r0, _ := Crt(rm)

	return r0
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func Gcd(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}

	return Gcd(y, x%y)
}

// a*p + b*q = Gcd(a, b)となるp, qを求める
func ExtGcd(a, b int) (int, int) {
	if b == 0 {
		return 1, 0
	}
	q, p := ExtGcd(b, a%b)
	q -= a / b * p
	return p, q
}

func Crt(mp map[int]int) (int, int) {
	r, m := 0, 1
	for m0, b0 := range mp {
		p, _ := ExtGcd(m, m0)
		//d := Gcd(p, q)
		//if (b0-r)%d != 0 {
		//	return 0, -1
		//}
		t := ((b0-r)%m0 + m0) % m0
		t = (t * p) % m0
		r += m * t
		m *= m0
	}
	return (r%m + m) % m, m
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
