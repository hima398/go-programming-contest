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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	mark := "HDCS"
	number := "A23456789TJQK"
	n := nextInt()
	ok := true
	m := make(map[string]bool)
	for i := 0; i < n; i++ {
		s := nextString()
		ok = ok && strings.Contains(mark, string(s[0])) && strings.Contains(number, string(s[1])) && !m[s]
		m[s] = true
	}
	if ok {
		PrintString("Yes")
	} else {
		PrintString("No")
	}
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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
