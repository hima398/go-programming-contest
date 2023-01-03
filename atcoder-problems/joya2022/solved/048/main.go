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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var s, t []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
		t = append(t, nextString())
	}
	for i := 0; i < n; i++ {
		possible := false
		for _, v := range []string{s[i], t[i]} {
			ok := true
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				if v == s[j] || v == t[j] {
					ok = false
				}
			}
			possible = possible || ok
		}
		if !possible {
			PrintString("No")
			return
		}
	}
	PrintString("Yes")
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
