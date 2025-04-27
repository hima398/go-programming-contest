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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var r, c, x []int
	for i := 0; i < n; i++ {
		r = append(r, nextInt())
		c = append(c, nextInt())
		x = append(x, nextInt())
	}

	ans := solve(n, r, c, x)

	Print(ans)
}

func solve(n int, r, c, x []int) int {
	field := make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		if field[r[i]] == nil {
			field[r[i]] = make(map[int]int)
		}
		field[r[i]][c[i]] = x[i]
	}

	//r[i]行目, c[i]列目の総和
	rows := make(map[int]int)
	cs := make(map[int]int)
	for i := 0; i < n; i++ {
		rows[r[i]] += x[i]
		cs[c[i]] += x[i]
	}

	//c[i]列目の総和を別のデータ構造に持ち帰る
	type colmn struct {
		j, sum int
	}
	var colmns []colmn
	for k, v := range cs {
		colmns = append(colmns, colmn{k, v})
	}
	sort.Slice(colmns, func(i, j int) bool {
		return colmns[i].sum > colmns[j].sum
	})

	var ans int
	for row, sum := range rows {
		for j := range colmns {
			if _, found := field[row][colmns[j].j]; !found {
				ans = Max(ans, sum+colmns[j].sum)
				break
			} else {
				ans = Max(ans, sum+colmns[j].sum-field[row][colmns[j].j])
			}
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
