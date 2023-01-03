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

	h, m := nextInt(), nextInt()
	check := func(h, m int) bool {
		h2 := h/10*10 + m/10
		m2 := 10*(h%10) + m%10
		//h2 := 10*(m%10) + h%10
		//m2 := m/10*10 + h/10
		return 0 <= h2 && h2 <= 23 && 0 <= m2 && m2 <= 59
	}
	max := 48 * 60
	for i := h*60 + m; i <= max; i++ {
		if check(i/60, i%60) {
			fmt.Println((i/60)%24, i%60)
			return
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
