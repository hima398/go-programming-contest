package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func solve(n int) {
	const INF = 1 << 60
	d := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		d[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == j {
				continue
			}
			d[i][j] = INF
		}
	}
	for u := 1; u <= 2; u++ {
		for v := 3; v <= n; v++ {
			//fmt.Fprintln(out, "?", u, v)
			fmt.Println("?", u, v)
			duv := nextInt()
			if duv < 0 {
				fmt.Println(-1)
				return
			}
			d[u][v] = Min(d[u][v], duv)
			d[v][u] = Min(d[v][u], duv)
		}
	}
	//PrintVertically(d)

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if k == i || k == j {
					continue
				}
				d[i][j] = Min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}
	if d[1][2] != 3 {
		fmt.Println("!", d[1][2])
		return
	}

	var cnt []int
	for i := 3; i <= n; i++ {
		dd := d[1][i] + d[i][2]
		if dd == 3 {
			cnt = append(cnt, i)
		}
	}
	if len(cnt) != 2 {
		fmt.Println("!", 1)
		return
	} else {
		// len(cnt)==2
		fmt.Println("?", cnt[0], cnt[1])
		duv := nextInt()
		if duv == 1 {
			fmt.Println("!", 3)
			return
		} else {
			//duv == 2 || duv == 3
			fmt.Println("!", 1)
			return
		}
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	solve(n)
	//defer out.Flush()
	//fmt.Fprintln(out, "!", ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func PrintHorizonaly(x []int) {
	fmt.Printf("%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Printf(" %d", x[i])
	}
	fmt.Println()
}

func PrintVertically(x [][]int) {
	for _, v := range x {
		PrintHorizonaly(v)
	}
}
