package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b, c := nextInt(), nextInt(), nextInt()
	type number struct {
		x    int
		name string
	}
	var ns []number
	ns = append(ns, number{a, "A"})
	ns = append(ns, number{b, "B"})
	ns = append(ns, number{c, "C"})
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].x < ns[j].x
	})
	PrintString(ns[1].name)
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
