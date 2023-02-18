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

	n, m := nextInt(), nextInt()
	c := make([]int, m)
	a := make([][]int, m)
	for i := 0; i < m; i++ {
		c[i] = nextInt()
		for j := 0; j < c[i]; j++ {
			a[i] = append(a[i], nextInt())
		}
	}
	pattern := 1<<m - 1
	var ans int
	for pat := 1; pat <= pattern; pat++ {
		t := make(map[int]struct{})
		for i := 0; i < m; i++ {
			if pat>>i&1 == 0 {
				continue
			}
			for _, v := range a[i] {
				t[v] = struct{}{}
			}
		}
		ok := true
		for i := 1; i <= n; i++ {
			_, found := t[i]
			ok = ok && found
		}
		if ok {
			ans++
		}
	}
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
