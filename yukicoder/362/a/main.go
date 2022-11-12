package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b := nextString(), nextString()
	ba, bb := big.NewInt(0), big.NewInt(0)
	ba.SetString(a, 10)
	bb.SetString(b, 10)
	ans := big.NewInt(0)
	ans = ans.Add(ba, bb)
	PrintString(ans.String())
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
