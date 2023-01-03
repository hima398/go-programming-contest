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

	n, m := nextInt(), nextInt()
	a := make([][]bool, n)
	for i := 0; i < n; i++ {
		a[i] = make([]bool, n)
	}
	var x [][]int
	for i := 0; i < m; i++ {
		k := nextInt()
		x = append(x, nextIntSlice(k))
		for j := range x[i] {
			x[i][j]--
		}
	}
	//fmt.Println(x)
	for _, row := range x {
		k := len(row)
		for i := 0; i < k; i++ {
			for j := i + 1; j < k; j++ {
				a[row[i]][row[j]] = true
				a[row[j]][row[i]] = true
			}
		}
	}
	ok := true
	for i := 0; i < n; i++ {
		//fmt.Println(a[i])
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			ok = ok && a[i][j]
		}
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

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
