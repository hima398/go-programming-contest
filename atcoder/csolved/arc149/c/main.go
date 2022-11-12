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

func solveHonestly(n int) [][]int {
	const eMax = 2 * int(1e6)
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	e := make([]bool, eMax+1)
	e[0] = true
	e[1] = true
	for i := 2; i <= eMax; i++ {
		if e[i] {
			continue
		}
		for j := i + i; j <= eMax; j += i {
			e[j] = true
		}
	}
	var p []int
	for i := 1; i <= n*n; i++ {
		p = append(p, i)
	}
	for {
		idx := 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				ans[i][j] = p[idx]
				idx++
			}
		}
		di := []int{-1, 0, 1, 0}
		dj := []int{0, -1, 0, 1}
		ok := true
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				for k := 0; k < 4; k++ {
					ni, nj := i+di[k], j+dj[k]
					if ni >= 0 && ni < n && nj >= 0 && nj < n {
						ok = ok && e[ans[i][j]+ans[ni][nj]]
					}
				}
			}
		}
		if ok {
			//PrintVertically(ans)
			//fmt.Println()
			return ans
		}

		if !NextPermutation(sort.IntSlice(p)) {
			break
		}
	}
	return nil
}

func solve(n int) [][]int {
	if n == 3 {
		return [][]int{
			{1, 3, 5},
			{8, 6, 9},
			{4, 2, 7}}
	} else if n == 4 {
		return [][]int{
			{15, 11, 16, 12},
			{13, 3, 6, 9},
			{14, 7, 8, 1},
			{4, 2, 10, 5}}
	} else if n == 5 {
		return [][]int{
			{1, 3, 5, 7, 9},
			{17, 19, 21, 25, 23},
			{11, 13, 15, 24, 2},
			{14, 12, 10, 4, 6},
			{8, 16, 18, 20, 22},
		}
	}
	// n>=6
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	var odd, even []int
	var odd3, even3 []int
	for i := 1; i <= n*n; i++ {
		if i%3 == 0 {
			if i%2 == 1 {
				odd3 = append(odd3, i)
			} else {
				even3 = append(even3, i)
			}
		} else {
			if i%2 == 1 {
				odd = append(odd, i)
			} else {
				even = append(even, i)
			}
		}
	}

	j := Ceil(n, 2) - 1
	for i := 0; i < n; i++ {
		ans[i][j] = odd3[0]
		ans[i][j+1] = even3[0]
		odd3 = odd3[1:]
		even3 = even3[1:]
		if n%2 == 1 && i == n/2 {
			j--
		}
	}
	for _, v := range odd3 {
		odd = append(odd, v)
	}
	for _, v := range even3 {
		even = append(even, v)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n && ans[i][j] == 0; j++ {
			ans[i][j] = odd[0]
			odd = odd[1:]
		}
	}

	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0 && ans[i][j] == 0; j-- {
			ans[i][j] = even[0]
			even = even[1:]
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	//ans := solveHonestly(n)
	ans := solve(n)
	PrintVertically(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func PrintVertically(x [][]int) {
	defer out.Flush()
	for _, v := range x {
		//fmt.Fprintln(out, v)
		PrintHorizonaly(v)
	}
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
