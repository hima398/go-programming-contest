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

	ax, ay := nextInt(), nextInt()
	bx, by := nextInt(), nextInt()
	cx, cy := nextInt(), nextInt()
	dx, dy := nextInt(), nextInt()

	ok := solve(ax, ay, bx, by, cx, cy, dx, dy)

	if ok {
		Print("Yes")
	} else {
		Print("No")
	}
}

func solve(ax, ay, bx, by, cx, cy, dx, dy int) bool {
	var vs [][2]int
	vs = append(vs, [2]int{ax, ay})
	vs = append(vs, [2]int{bx, by})
	vs = append(vs, [2]int{cx, cy})
	vs = append(vs, [2]int{dx, dy})

	cross := func(a, b [2]int) int {
		return a[0]*b[1] - a[1]*b[0]
	}
	for i := 0; i < 4; i++ {
		v1, v2, v3 := vs[i], vs[(i+1)%4], vs[(i+2)%4]
		if cross([2]int{v2[0] - v1[0], v2[1] - v1[1]}, [2]int{v3[0] - v1[0], v3[1] - v1[1]}) <= 0 {
			return false
		}
	}
	return true
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
