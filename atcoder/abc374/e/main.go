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

	n, x := nextInt(), nextInt()
	var a, p, b, q []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		p = append(p, nextInt())
		b = append(b, nextInt())
		q = append(q, nextInt())
	}

	ans := solve(n, x, a, p, b, q)

	Print(ans)
}

// 製造能力wを達成できる最小コストがx以下かどうかを判定する
func check(n, x, w int, a, p, b, q []int) bool {
	var cost int
	for i := 0; i < n; i++ {
		c := int(1e9) + 1
		for j := 0; j <= b[i]; j++ {
			rem := Ceil(Max(w-j*a[i], 0), b[i])
			c = Min(c, j*p[i]+rem*q[i])
		}
		for j := 0; j <= a[i]; j++ {
			rem := Ceil(Max(w-j*b[i], 0), a[i])
			c = Min(c, rem*p[i]+j*q[i])
		}

		cost += c
	}
	//fmt.Println(w, cost, x)
	return cost <= x
}

func solve(n, x int, a, p, b, q []int) int {
	ok, ng := 0, int(1e9)+1
	for ng-ok > 1 {
		mid := (ng + ok) / 2
		if check(n, x, mid, a, p, b, q) {
			//fmt.Printf("mid = %d: ok\n", mid)
			ok = mid
		} else {
			//fmt.Printf("mid = %d: ng\n", mid)
			ng = mid
		}
	}
	return ok
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Floor(x, y int) int {
	/*
		return x / y
	*/
	r := (x%y + y) % y
	return (x - r) / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
