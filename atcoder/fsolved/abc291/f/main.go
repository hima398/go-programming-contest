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

	n, m := nextInt(), nextInt()
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}
	ans := solve(n, m, s)
	PrintHorizonaly(ans)
}

func solve(n, m int, s []string) []int {
	const INF = 1 << 60

	//0(0-indexed)からの最短距離を求める
	dist1 := make([]int, n)
	for i := range dist1 {
		dist1[i] = INF
	}
	var q1 []int
	q1 = append(q1, 0)
	dist1[0] = 0
	for len(q1) > 0 {
		cur := q1[0]
		q1 = q1[1:]
		for i, si := range s[cur] {
			if si == '1' {
				next := cur + i + 1
				if dist1[next] < INF {
					continue
				}
				q1 = append(q1, next)
				dist1[next] = Min(dist1[next], dist1[cur]+1)
			}
		}
	}

	//iからn-1(0-indexed)の最短距離を求める
	distN := make([]int, n)
	for i := range distN {
		distN[i] = INF
	}
	var q2 []int
	q2 = append(q2, n-1)
	distN[n-1] = 0
	for len(q2) > 0 {
		cur := q2[0]
		q2 = q2[1:]
		for i := 0; i < m; i++ {
			next := cur - i - 1
			if next < 0 {
				break
			}
			if s[next][i] == '1' {
				if distN[next] < INF {
					continue
				}
				q2 = append(q2, next)
				distN[next] = Min(distN[next], distN[cur]+1)
			}
		}
	}
	//fmt.Println(dist1)
	//fmt.Println(distN)
	var ans []int
	for i := 1; i < n-1; i++ {
		v := INF
		for j := Max(i-m, 0); j < i; j++ {
			for k := i + 1; k < Min(i+m, n); k++ {
				d := k - j
				if d > m {
					continue
				}
				if s[j][d-1] == '0' {
					continue
				}
				//j->kへの道がある
				v = Min(v, dist1[j]+distN[k]+1)
			}
		}
		ans = append(ans, v)
	}
	for i := range ans {
		if ans[i] == INF {
			ans[i] = -1
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
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
