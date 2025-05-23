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

	q := nextInt()
	var deck []int
	for i := 0; i < q; i++ {
		t := nextInt()
		switch t {
		case 1:
			x := nextInt()
			deck = append(deck, x)
		case 2:
			n := len(deck)
			var ans int
			if n > 0 {
				ans = deck[n-1]
				deck = deck[:n-1]
			}
			Print(ans)
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
