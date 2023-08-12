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

	h, w := nextInt(), nextInt()
	var s []string
	for i := 0; i < h; i++ {
		s = append(s, nextString())
	}
	ok := solve(h, w, s)
	if ok {
		PrintString("Yes")
	} else {
		PrintString("No")
	}
}

func solve(h, w int, s []string) bool {
	if s[0][0] != 's' {
		return false
	}
	m := map[byte]int{'s': 0, 'n': 1, 'u': 2, 'k': 3, 'e': 4}
	type cell struct {
		i, j int
		k    int
	}
	visited := make([][][5]bool, h)
	for i := range visited {
		visited[i] = make([][5]bool, w)
	}
	var q []cell
	q = append(q, cell{0, 0, 0})
	visited[0][0][0] = true
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for l := 0; l < 4; l++ {
			ni, nj := cur.i+di[l], cur.j+dj[l]
			if ni < 0 || ni >= h || nj < 0 || nj >= w {
				continue
			}
			if m[s[ni][nj]] != (cur.k+1)%5 {
				continue
			}
			nk := m[s[ni][nj]]
			if visited[ni][nj][nk] {
				continue
			}
			q = append(q, cell{ni, nj, nk})
			visited[ni][nj][nk] = true
		}
	}
	var ok bool

	for k := 0; k < 5; k++ {
		ok = ok || visited[h-1][w-1][k]
	}
	return ok
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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
