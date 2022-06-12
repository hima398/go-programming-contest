package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

var rToI = map[rune]int{
	'.': 0,
	'#': 1,
}
var m = map[int]string{
	0x7555: "0",
	0x2622: "1",
	0x7174: "2",
	0x7171: "3",
	0x5571: "4",
	0x7471: "5",
	0x7475: "6",
	0x7111: "7",
	0x7575: "8",
	0x7571: "9",
}

func ParseNumber(s string) string {
	var v = 0
	for _, c := range s {
		v |= rToI[c]
		v = v << 1
	}
	v = v >> 1
	return m[v]
}

func ReadNumber(lines []string, w int) string {
	var v string
	n := w - 1
	for i := 0; i < n; i += 4 {
		s := lines[0][i : i+4]
		s += lines[1][i : i+4]
		s += lines[2][i : i+4]
		s += lines[3][i : i+4]
		v += ParseNumber(s)
	}
	return v
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	w := 4*n + 1
	lines := make([]string, w)
	for i := 0; i < 5; i++ {
		lines[i] = nextString()
	}

	ans := ReadNumber(lines, w)

	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
