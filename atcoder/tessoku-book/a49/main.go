package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

const n = 20

type state struct {
	score         int
	x             [n]int
	prevOperation string
	prevState     int
}

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := nextInt()
	var p, q, r []int
	for i := 0; i < t; i++ {
		p = append(p, nextInt()-1)
		q = append(q, nextInt()-1)
		r = append(r, nextInt()-1)
	}
	//ans := solveRandom(t, p, q, r)
	//ans := solveGreedily(t, p, q, r)
	ans := solveBeamSearch(t, p, q, r)
	for _, v := range ans {
		PrintString(v)
	}
}

func solveBeamSearch(t int, p, q, r []int) []string {
	const beamWidth = 5000

	//nums := make([]int, t+1)
	var nums [101]int
	nums[0] = 1
	var beam [101][beamWidth]state
	//for i := 0; i <= t; i++ {
	//	beam[i] = make([]state, beamWidth)
	//}
	var v state
	beam[0][0] = v

	newState := func(s state) state {
		var res state
		res.score = s.score
		res.x = s.x
		res.prevOperation = s.prevOperation
		res.prevState = s.prevState
		return res
	}
	for i := 1; i <= t; i++ {
		var candidate []state
		for j := 0; j < nums[i-1]; j++ {
			a := newState(beam[i-1][j])
			a.x[p[i-1]]++
			a.x[q[i-1]]++
			a.x[r[i-1]]++
			for k := 0; k < n; k++ {
				if a.x[k] == 0 {
					a.score++
				}
			}
			a.prevOperation = "A"
			a.prevState = j
			candidate = append(candidate, a)
			b := newState(beam[i-1][j])
			b.x[p[i-1]]--
			b.x[q[i-1]]--
			b.x[r[i-1]]--
			for k := 0; k < n; k++ {
				if b.x[k] == 0 {
					b.score++
				}
			}
			b.prevOperation = "B"
			b.prevState = j
			candidate = append(candidate, b)
		}
		//fmt.Println(candidate)
		sort.Slice(candidate, func(i, j int) bool {
			return candidate[i].score > candidate[j].score
		})
		nums[i] = Min(len(candidate), beamWidth)
		//fmt.Println("nums[i]=", nums[i])
		for j := 0; j < nums[i]; j++ {
			//beam[i][j] = candidate[j]
			beam[i][j] = candidate[j]
		}
		//fmt.Println(beam)
	}
	var buf []string
	var cur int
	for i := t; i >= 1; i-- {
		buf = append(buf, beam[i][cur].prevOperation)
		cur = beam[i][cur].prevState
	}

	ans := make([]string, t)
	for i := 0; i < t; i++ {
		ans[i] = buf[t-i-1]
	}
	return ans
}

func solveGreedily(t int, p, q, r []int) []string {
	var ans []string
	const size = 20
	x := make([]int, size)
	for k := 0; k < t; k++ {
		xa := make([]int, size)
		copy(xa, x)
		xa[p[k]]++
		xa[q[k]]++
		xa[r[k]]++
		xb := make([]int, size)
		copy(xb, x)
		xb[p[k]]--
		xb[q[k]]--
		xb[r[k]]--
		var sa, sb int
		for i := 0; i < size; i++ {
			if xa[i] == 0 {
				sa++
			}
			if xb[i] == 0 {
				sb++
			}
		}
		if sa >= sb {
			ans = append(ans, "A")
			x = xa
		} else {
			ans = append(ans, "B")
			x = xb
		}
	}
	return ans
}

func solveRandom(t int, p, q, r []int) []string {
	var ans []string
	for i := 0; i < t; i++ {
		v := rand.Intn(2)
		if v == 0 {
			ans = append(ans, "A")
		} else {
			ans = append(ans, "B")
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
