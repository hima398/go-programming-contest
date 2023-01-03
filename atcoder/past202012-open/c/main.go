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
	m := make([]string, 36)
	for i := 0; i < 10; i++ {
		m[i] = string(i + '0')
	}
	for i := 0; i < 26; i++ {
		m[i+10] = string(i + 'A')
	}
	var t []string
	for n > 0 {
		x := n % 36
		n /= 36
		t = append(t, m[x])
	}
	var ans string
	if len(t) == 0 {
		ans += "0"
	} else {
		for i := len(t) - 1; i >= 0; i-- {
			ans += t[i]
		}
	}
	PrintString(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
