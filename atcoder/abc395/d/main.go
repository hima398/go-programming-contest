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

	n, q := nextInt(), nextInt()
	t, a, b := make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = nextInt()
		if t[i] == 3 {
			a[i] = nextInt() - 1
		} else {
			a[i], b[i] = nextInt()-1, nextInt()-1
		}
	}

	ans := solve(n, q, t, a, b)

	for _, v := range ans {
		Print(v)
	}
}

func solve(n, q int, t, a, b []int) []int {
	//i番目の鳩が入っている巣の番号
	pigeons := make([]int, n)
	//i番目の巣に入っているラベルの番号
	nests := make([]int, n)
	//i番目のラベルが入っている巣の番号
	labels := make([]int, n)
	for i := 0; i < n; i++ {
		pigeons[i] = i
		nests[i] = i
		labels[i] = i
	}

	var ans []int
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			pigeons[a[i]] = labels[b[i]]
		case 2:
			labels[a[i]], labels[b[i]] = labels[b[i]], labels[a[i]]
			nests[labels[a[i]]], nests[labels[b[i]]] = nests[labels[b[i]]], nests[labels[a[i]]]
		case 3:
			ans = append(ans, nests[pigeons[a[i]]]+1)
		}
	}
	return ans
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
