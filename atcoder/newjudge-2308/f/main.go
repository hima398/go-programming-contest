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
		Print(-1)
	} else {
		for _, v := range ans {
			Print(v)
		}
	}
}

func solve(s, t, m int, u, v []int) ([]int, error) {
	e := make([][]int, s+t)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], v[i])
	}
	//集合V2
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
				ii, jj := e[k][i]-s, e[k][j]-s
				if v2[ii][jj] >= 0 {
					return []int{k + 1, v2[ii][jj] + 1, ii + s + 1, jj + s + 1}, nil
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
