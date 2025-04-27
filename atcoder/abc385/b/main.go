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

	h, w, x, y := nextInt(), nextInt(), nextInt()-1, nextInt()-1
	var s []string
	for i := 0; i < h; i++ {
		s = append(s, nextString())
	}
	t := nextString()

	dirs := map[rune][2]int{'U': {-1, 0}, 'D': {1, 0}, 'L': {0, -1}, 'R': {0, 1}}
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}
	var ans int
	if s[x][y] == '@' {
		ans++
	}
	visited[x][y] = true
	for _, ti := range t {
		nx, ny := x+dirs[ti][0], y+dirs[ti][1]
		if nx < 0 || nx >= h || ny < 0 || ny >= w {
			continue
		}
		if s[nx][ny] == '#' {
			continue
		}
		x, y = nx, ny
		if s[x][y] == '@' && !visited[x][y] {
			ans++
		}
		visited[x][y] = true
	}

	fmt.Println(x+1, y+1, ans)
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
