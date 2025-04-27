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

	decodeTable := map[string]int{"-----": 0, ".----": 1, "..---": 2, "...--": 3, "....-": 4, ".....": 5, "-....": 6, "--...": 7, "---..": 8, "----.": 9}

	n := nextInt()
	var s []string

	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}

	var ans string
	for _, si := range s {
		ans += strconv.Itoa(decodeTable[si])
	}

	Print(ans)
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
