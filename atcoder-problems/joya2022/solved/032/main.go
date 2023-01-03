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

	r, c := nextInt()-1, nextInt()-1
	var f [15][15]int
	for k := 0; k <= 15/2; k++ {
		for i := k; i < 15-k; i++ {
			for j := k; j < 15-k; j++ {
				f[i][j] = k % 2
			}
		}
	}
	//for _, row := range f {
	//	fmt.Println(row)
	//}
	if f[r][c] == 0 {
		PrintString("black")
	} else {
		PrintString("white")
	}
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
