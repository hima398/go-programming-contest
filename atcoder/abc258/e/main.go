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

func solve(n, q, x int, w, k []int) []int {
	s := make([]int, 2*n)
	s[0] = w[0]
	for i := 1; i < 2*n; i++ {
		s[i] = s[i-1] + w[i%n]
	}
	//PrintHorizonaly(s)

	nums := make([]int, n)
	idx := 0
	route := []int{0}
	visited := make([]int, n)
	for {
		visited[idx]++
		xx := x
		var w int
		if xx > s[2*n-1] {
			w = xx / s[2*n-1]
			xx %= s[2*n-1]
		}

		next := sort.Search(2*n, func(i int) bool {
			if idx == 0 {
				return xx <= s[i]
			} else {
				return xx <= s[i]-s[idx]
			}
		})
		if nums[idx] == 0 {
			nums[idx] = w*n + next - idx + 1
		}
		idx = (next + 1) % n
		//fmt.Println("idx = ", next)
		if visited[idx] == 0 {
			route = append(route, idx)
		}
		if visited[idx] >= 2 {
			break
		}
	}
	//fmt.Println("nums = ", nums)
	//fmt.Println("route = ", route)
	var root, l int
	for i := range visited {
		if visited[i] == 1 {
			root++
		} else if visited[i] == 2 {
			l++
		}
	}

	//PrintHorizonaly(visited)
	//fmt.Println(root, l)
	var ans []int
	for i := 0; i < q; i++ {
		k[i]--
		if k[i] < root {
			ans = append(ans, nums[route[k[i]]])
		} else {
			idx := root + (k[i]-root)%l
			ans = append(ans, nums[route[idx]])
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q, x := nextInt(), nextInt(), nextInt()
	w := nextIntSlice(n)
	var k []int
	for i := 0; i < q; i++ {
		k = append(k, nextInt())
	}
	ans := solve(n, q, x, w, k)
	PrintVertically(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
