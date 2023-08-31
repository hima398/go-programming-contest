package main

import (
	"bufio"
	"errors"
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

	s, t, m := nextInt(), nextInt(), nextInt()
	var u, v []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}

	ans, err := solve(s, t, m, u, v)

	if err != nil {
		PrintInt(-1)
	} else {
		PrintHorizonaly(ans)
	}
}

func solve(s, t, m int, u, v []int) ([]int, error) {
	e := make([][]int, s+t)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], v[i])
	}
	//v2[i][j]で集合V2内のs+i, s+jがv2[i][j]を経由して連結かどうかを示す
	v2 := make([][]int, t)
	for i := range v2 {
		v2[i] = make([]int, t)
		for j := range v2[i] {
			v2[i][j] = -1
		}
	}

	for k := 0; k < s; k++ {
		for i := 0; i < len(e[k]); i++ {
			for j := i + 1; j < len(e[k]); j++ {
				//V1内の点kから、V2内の点 e[k][i], e[k][j]に連結している
				ii, jj := e[k][i]-s, e[k][j]-s
				if v2[ii][jj] >= 0 {
					return []int{k + 1, v2[ii][jj] + 1, e[k][i] + 1, e[k][j] + 1}, nil
				} else {
					v2[ii][jj] = k
					v2[jj][ii] = k
				}
			}
		}
	}
	return nil, errors.New("Impossible")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
