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

	k, g, m := nextInt(), nextInt(), nextInt()
	var glass, mug int
	for i := 0; i < k; i++ {
		if glass == g {
			glass = 0
		} else if mug == 0 {
			mug = m
		} else {
			//マグカップからグラスに移す水の量
			w := Min(g-glass, mug)
			glass += w
			mug -= w
		}
	}
	fmt.Println(glass, mug)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
