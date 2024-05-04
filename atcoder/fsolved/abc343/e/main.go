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

	v := nextIntSlice(3)

	ans, err := solve(v)

	if err != nil {
		Print("No")
	} else {
		Print("Yes")
		PrintHorizonaly(ans)
	}
}

type Point struct {
	x, y, z int
}

type Cube struct {
	min Point
	max Point
}

func solve(v []int) ([]int, error) {
	const n = 7

	computeOverlapedVolume2 := func(c1, c2 Cube) int {
		minX := Max(c1.min.x, c2.min.x)
		minY := Max(c1.min.y, c2.min.y)
		minZ := Max(c1.min.z, c2.min.z)

		maxX := Min(c1.max.x, c2.max.x)
		maxY := Min(c1.max.y, c2.max.y)
		maxZ := Min(c1.max.z, c2.max.z)

		lengthX := Max(0, maxX-minX)
		lengthY := Max(0, maxY-minY)
		lengthZ := Max(0, maxZ-minZ)

		return lengthX * lengthY * lengthZ
	}

	computeOverlapedVolume3 := func(c1, c2, c3 Cube) int {
		minX := Max(c1.min.x, Max(c2.min.x, c3.min.x))
		minY := Max(c1.min.y, Max(c2.min.y, c3.min.y))
		minZ := Max(c1.min.z, Max(c2.min.z, c3.min.z))

		maxX := Min(c1.max.x, Min(c2.max.x, c3.max.x))
		maxY := Min(c1.max.y, Min(c2.max.y, c3.max.y))
		maxZ := Min(c1.max.z, Min(c2.max.z, c3.max.z))

		lengthX := Max(0, maxX-minX)
		lengthY := Max(0, maxY-minY)
		lengthZ := Max(0, maxZ-minZ)

		return lengthX * lengthY * lengthZ
	}

	computeVolumes := func(c1, c2, c3 Cube) []int {
		v3 := computeOverlapedVolume3(c1, c2, c3)
		v12, v13, v23 := computeOverlapedVolume2(c1, c2), computeOverlapedVolume2(c1, c3), computeOverlapedVolume2(c2, c3)
		v2 := v12 + v13 + v23 - 3*v3
		v1 := 3*n*n*n - 2*v2 - 3*v3
		return []int{v1, v2, v3}
	}

	c1 := Cube{Point{n, n, n}, Point{2 * n, 2 * n, 2 * n}}
	for i := 0; i <= 2*n; i++ {
		for j := 0; j <= 2*n; j++ {
			for k := 0; k <= 2*n; k++ {
				c2 := Cube{Point{i, j, k}, Point{i + n, j + n, k + n}}
				for ii := 0; ii <= 2*n; ii++ {
					for jj := 0; jj <= 2*n; jj++ {
						for kk := 0; kk <= 2*n; kk++ {
							c3 := Cube{Point{ii, jj, kk}, Point{ii + n, jj + n, kk + n}}
							u := computeVolumes(c1, c2, c3)
							ok := true
							for idx := 0; idx < 3; idx++ {
								ok = ok && u[idx] == v[idx]
							}
							if ok {
								return []int{n, n, n, i, j, k, ii, jj, kk}, nil
							}
						}
					}
				}
			}
		}
	}
	return nil, errors.New("Impossible")
}

func solveHonestly(v []int) ([]int, error) {
	const n = 7
	var t [3 * n][3 * n][3 * n]int
	for i := n; i < 2*n; i++ {
		for j := n; j < 2*n; j++ {
			for k := n; k < 2*n; k++ {
				t[i][j][k]++
			}
		}
	}

	fillField := func(a, b, c []int) []int {
		var f [3 * n][3 * n][3 * n]int
		for idx := 0; idx < 3; idx++ {
			for i := a[idx]; i < a[idx]+n; i++ {
				for j := b[idx]; j < b[idx]+n; j++ {
					for k := c[idx]; k < c[idx]+n; k++ {
						f[i][j][k]++
					}
				}
			}
		}
		ret := make([]int, 3)
		for i := 0; i < 3*n; i++ {
			for j := 0; j < 3*n; j++ {
				for k := 0; k < 3*n; k++ {
					if f[i][j][k] == 0 {
						continue
					}
					ret[f[i][j][k]-1]++
				}
			}
		}
		return ret
	}
	// 7529536 ~ 1e6
	for i := 0; i <= 2*n; i++ {
		for j := 0; j <= 2*n; j++ {
			for k := 0; k <= 2*n; k++ {
				for ii := 0; ii <= 2*n; ii++ {
					for jj := 0; jj <= 2*n; jj++ {
						for kk := 0; kk <= 2*n; kk++ {
							u := fillField([]int{n, i, ii}, []int{n, j, jj}, []int{n, k, kk})
							ok := true
							for idx := 0; idx < 3; idx++ {
								ok = ok && u[idx] == v[idx]
							}
							if ok {
								return []int{n, n, n, i, j, k, ii, jj, kk}, nil
							}
						}
					}
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

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Print(x any) {
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
