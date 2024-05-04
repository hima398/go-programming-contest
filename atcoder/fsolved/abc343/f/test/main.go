package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var out = bufio.NewWriter(os.Stdout)

func main() {
	const maxA = int(1e9)
	n, q := 2*int(1e5), 2*int(1e5)

	defer out.Flush()
	fmt.Fprintln(out, n, q)
	var a []int
	for i := 0; i < n; i++ {
		a = append(a, rand.Intn(maxA)+1)
	}

	defer out.Flush()
	fmt.Fprintf(out, "%d", a[0])
	for i := 1; i < n; i++ {
		fmt.Fprintf(out, " %d", a[i])
	}
	fmt.Fprintln(out)

	t, p, x, l, r := make([]int, q), make([]int, q), make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = rand.Intn(2) + 1
		switch t[i] {
		case 1:
			p[i] = rand.Intn(n) + 1
			x[i] = rand.Intn(maxA) + 1
		case 2:
			for l[i] == 0 || r[i] == 0 || l[i] > r[i] {
				l[i] = rand.Intn(n) + 1
				r[i] = rand.Intn(n) + 1
			}
		}
	}
	out.Flush()
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			fmt.Fprintln(out, t[i], p[i], x[i])
		case 2:
			fmt.Fprintln(out, t[i], l[i], r[i])
		}
	}
}
