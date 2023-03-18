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

	n, q := nextInt(), nextInt()
	p := make([][2]int, n)
	for i := 0; i < q; i++ {
		t, x := nextInt(), nextInt()-1
		switch t {
		case 1:
			p[x][0]++
		case 2:
			p[x][1]++
		case 3:
			if p[x][0] >= 2 || p[x][1] >= 1 {
				PrintString("Yes")
			} else {
				PrintString("No")
			}
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
