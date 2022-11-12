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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b, c, d, e, f := nextInt64(), nextInt64(), nextInt64(), nextInt64(), nextInt64(), nextInt64()
	//ans := solveInt(a, b, c, d, e, f)
	ans := solveBigInt(a, b, c, d, e, f)
	PrintInt(ans)
}

func solveInt(a, b, c, d, e, f int64) int64 {
	const p = 998244353

	a %= p
	b %= p
	c %= p
	d %= p
	e %= p
	f %= p

	g := a * b % p
	g = g * c % p
	h := d * e % p
	h = h * f % p
	ans := (g - h + p) % p

	return ans
}

func solveBigInt(a, b, c, d, e, f int64) int64 {
	p := big.NewInt(998244353)
	ba, bb, bc, bd, be, bf := big.NewInt(a), big.NewInt(b), big.NewInt(c), big.NewInt(d), big.NewInt(e), big.NewInt(f)

	bg := ba.Mul(ba, bb)
	bg = bg.Mul(bg, bc)
	bh := bd.Mul(bd, be)
	bh = bh.Mul(bh, bf)
	ans := bg.Sub(bg, bh)
	ans = ans.Mod(ans, p)
	return ans.Int64()
}

func nextInt64() int64 {
	sc.Scan()
	i, _ := strconv.ParseInt(sc.Text(), 10, 64)
	return i
}

func PrintInt(x int64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
