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

	n := nextInt()
	var s, t []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
		t = append(t, nextString())
	}
	for i := 0; i < n; i++ {
		canUseS, canUseT := true, true
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			canUseS = canUseS && s[i] != s[j] && s[i] != t[j]
			canUseT = canUseT && t[i] != s[j] && t[i] != t[j]
		}
		if !canUseS && !canUseT {
			Print("No")
			return
		}
	}
	Print("Yes")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
