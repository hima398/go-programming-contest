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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b, c := nextInt(), nextInt(), nextInt()
	var memo [101][101][101]float64
	var dfs func(x, y, z int) float64
	dfs = func(x, y, z int) float64 {
		if memo[x][y][z] > 0 {
			return memo[x][y][z]
		}
		if x >= 100 || y >= 100 || z >= 100 {
			return 0
		}
		total := float64(x + y + z)
		var res float64
		res += dfs(x+1, y, z) * float64(x) / total
		res += dfs(x, y+1, z) * float64(y) / total
		res += dfs(x, y, z+1) * float64(z) / total
		res += 1.0
		memo[x][y][z] = res
		return memo[x][y][z]
	}
	ans := dfs(a, b, c)
	PrintFloat64(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
