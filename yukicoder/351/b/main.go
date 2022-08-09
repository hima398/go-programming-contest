package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solveHonestly(l, r, a, b, c int) int {
	const p = 998244353
	var ans int
	for x := l; x <= r; x++ {
		for y := l; y <= r; y++ {
			for z := l; z <= r; z++ {
				if x+a <= y && y+b <= z && x+c <= z {
					fmt.Println(x, y, z)
					ans++
					ans %= p
				}
			}
		}
	}
	return ans
}

func solve(l, r, a, b, c int) string {
	const p = 998244353
	bp := big.NewInt(p)

	cd := Max(c-b-a, 0)
	if cd < 0 {
		return "0"
	}
	bcd := big.NewInt(int64(cd))

	bl, br := big.NewInt(int64(l)), big.NewInt(int64(r))
	ba, bb := big.NewInt(int64(a)), big.NewInt(int64(b))
	//ba, bb, bc := big.NewInt(int64(a)), big.NewInt(int64(b)), big.NewInt(int64(c))

	one := big.NewInt(1)
	two := big.NewInt(2)

	// d = r+1 - (l+a+b+c')
	d := big.NewInt(0).Add(br, one)

	d = d.Sub(d, big.NewInt(0).Add(bcd, big.NewInt(0).Add(bb, big.NewInt(0).Add(bl, ba))))
	d = d.Mod(d, bp)

	// d1 = d + 1
	d1 := big.NewInt(0).Add(d, one)
	// d2 = 2*d + 1
	d2 := big.NewInt(0).Add(big.NewInt(0).Mul(two, d), one)

	i2 := big.NewInt(2)
	i2 = i2.ModInverse(i2, bp)
	i6 := big.NewInt(6)
	i6 = i6.ModInverse(i6, bp)

	//fmt.Println("i2, i6 = ", i2, i6)

	a0 := big.NewInt(0).Mul(big.NewInt(0).Mul(d1, bcd), d)
	a1 := big.NewInt(0).Mul(big.NewInt(0).Sub(d1, bcd), big.NewInt(0).Mul(d1, d))
	a2 := big.NewInt(0).Mul(d2, big.NewInt(0).Mul(d, d1))
	//fmt.Println("a2, a1, a0 = ", a2, a1, a0)
	a1 = big.NewInt(0).Mul(a1, i2)
	a1 = a1.Mod(a1, bp)
	a2 = big.NewInt(0).Mul(a2, i6)
	a2 = a2.Mod(a2, bp)

	ans := big.NewInt(0).Sub(big.NewInt(0).Add(a0, a1), a2)
	ans = ans.Mod(ans, bp)

	//fmt.Println("cd, d = ", cd, d)
	return ans.String()
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	l, r := nextInt(), nextInt()
	a, b, c := nextInt(), nextInt(), nextInt()

	//ans := solveHonestly(l, r, a, b, c)
	ans := solve(l, r, a, b, c)
	PrintString(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintString(x string) {
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
