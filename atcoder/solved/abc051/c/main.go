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

func solve(sx, sy, tx, ty int) string {
	var ans string
	dx, dy := tx-sx, ty-sy
	ans += strings.Repeat("R", dx)
	ans += strings.Repeat("U", dy)
	ans += strings.Repeat("L", dx)
	ans += strings.Repeat("D", dy)
	ans += "D"
	ans += strings.Repeat("R", dx+1)
	ans += strings.Repeat("U", dy+1)
	ans += "L"
	ans += "U"
	ans += strings.Repeat("L", dx+1)
	ans += strings.Repeat("D", dy+1)
	ans += "R"
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	sx, sy, tx, ty := nextInt(), nextInt(), nextInt(), nextInt()
	ans := solve(sx, sy, tx, ty)
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
