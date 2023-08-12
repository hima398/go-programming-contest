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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b := big.NewInt(0), big.NewInt(0)
	a.SetString(nextString(), 10)
	b.SetString(nextString(), 10)
	if a.Cmp(b) < 0 {
		PrintString("LESS")
	} else if a.Cmp(b) > 0 {
		PrintString("GREATER")
	} else {
		PrintString("EQUAL")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
