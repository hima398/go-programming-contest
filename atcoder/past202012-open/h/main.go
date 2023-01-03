package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	r, c := nextInt()-1, nextInt()-1
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = nextString()
	}
	ans := solve(h, w, r, c, s)
	for _, v := range ans {
		PrintString(v)
	}
}

func solve(h, w, r, c int, s []string) []string {
	type cell struct {
		i, j int
	}
	f := make([][]string, h)
	for i := 0; i < h; i++ {
		f[i] = make([]string, w)
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				f[i][j] = "#"
			} else {
				f[i][j] = "-"
			}
		}
	}
	var q []cell
	q = append(q, cell{r, c})
	f[r][c] = "o"
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	dk := []byte{'v', '>', '^', '<'}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for dir := 0; dir < 4; dir++ {
			ni, nj := cur.i+di[dir], cur.j+dj[dir]
			//範囲外
			if ni < 0 || ni >= h || nj < 0 || nj >= w {
				continue
			}
			//壁かすでに到達できたところ
			if f[ni][nj] == "#" || f[ni][nj] == "o" {
				continue
			}
			//一方通行により進めない
			if s[ni][nj] != '.' && s[ni][nj] != dk[dir] {
				continue
			}
			q = append(q, cell{ni, nj})
			f[ni][nj] = "o"
		}
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if f[i][j] == "-" {
				f[i][j] = "x"
			}
		}
	}
	ans := make([]string, h)
	for i := 0; i < h; i++ {
		ans[i] = strings.Join(f[i], "")
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
