package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	_, _ = nextInt(), nextInt()
	s := nextString()
	t := nextString()

	hasPrefix := strings.HasPrefix(t, s)
	hasSuffix := strings.HasSuffix(t, s)

	if hasPrefix && hasSuffix {
		Print(0)
	} else if hasPrefix {
		Print(1)
	} else if hasSuffix {
		Print(2)
	} else {
		Print(3)
	}
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
