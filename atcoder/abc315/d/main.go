package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/liyue201/gostl/ds/queue"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	var c []string
	for i := 0; i < h; i++ {
		c = append(c, nextString())
	}
	ans := solve(h, w, c)
	Print(ans)
}

func solve(h, w int, c []string) int {
	type item struct {
		c map[byte]int
	}
	type node struct {
		k byte
		v int
	}

	row := queue.New[item]()
	for i := 0; i < h; i++ {
		var itm item
		itm.c = make(map[byte]int)
		for j := 0; j < w; j++ {
			itm.c[c[i][j]]++
		}
		row.Push(itm)
	}

	col := queue.New[item]()
	for j := 0; j < w; j++ {
		var itm item
		itm.c = make(map[byte]int)
		for i := 0; i < h; i++ {
			itm.c[c[i][j]]++
		}
		col.Push(itm)
	}

	deleteRow := queue.New[node]()
	deleteCol := queue.New[node]()

	//fmt.Println("row = ", row)
	//fmt.Println("col = ", col)
	//fmt.Println("delete(row) = ", deleteRow)
	//fmt.Println("delete(col) = ", deleteCol)

	for {
		//前のループで操作1、2後の操作3
		for !deleteRow.Empty() {
			nextCol := queue.New[item]()
			cur1 := deleteRow.Pop()
			for !col.Empty() {
				cur2 := col.Pop()
				if _, found := cur2.c[cur1.k]; found {
					cur2.c[cur1.k] -= cur1.v
					if cur2.c[cur1.k] <= 0 {
						delete(cur2.c, cur1.k)
					}
				}
				nextCol.Push(cur2)
			}
			col = nextCol
		}
		for !deleteCol.Empty() {
			nextRow := queue.New[item]()
			cur1 := deleteCol.Pop()
			for !row.Empty() {
				cur2 := row.Pop()
				if _, found := cur2.c[cur1.k]; found {
					cur2.c[cur1.k] -= cur1.v
					if cur2.c[cur1.k] <= 0 {
						delete(cur2.c, cur1.k)
					}
				}
				nextRow.Push(cur2)
			}
			row = nextRow
		}

		//操作1
		nextRow := queue.New[item]()
		for !row.Empty() {
			cur := row.Pop()
			if len(cur.c) == 1 {
				for k, v := range cur.c {
					if v >= 2 {
						deleteRow.Push(node{k, v})
					} else {
						nextRow.Push(cur)
					}
				}
			} else {
				nextRow.Push(cur)
			}
		}
		row = nextRow

		//操作2
		nextCol := queue.New[item]()
		for !col.Empty() {
			cur := col.Pop()
			if len(cur.c) == 1 {
				for k, v := range cur.c {
					if v >= 2 {
						deleteCol.Push(node{k, v})
					} else {
						nextCol.Push(cur)
					}
				}
			} else {
				nextCol.Push(cur)
			}
		}
		col = nextCol
		//fmt.Println("row = ", row)
		//fmt.Println("col = ", col)
		//fmt.Println("delete(row) = ", deleteRow)
		//fmt.Println("delete(col) = ", deleteCol)

		//ループを抜ける判定
		if deleteRow.Empty() && deleteCol.Empty() {
			break
		}
	}
	var ans int
	for !row.Empty() {
		cur := row.Pop()
		for _, v := range cur.c {
			ans += v
		}
	}

	return ans
}

func fistsolve(h, w int, c []string) int {
	//row := make([]map[byte]int, h)
	row := make([][]int, h)
	//col := make([]map[byte]int, w)
	col := make([][]int, w)
	for i := 0; i < h; i++ {
		row[i] = make([]int, 26)
		for j := 0; j < w; j++ {
			row[i][int(c[i][j]-'a')]++
		}
	}
	for j := 0; j < w; j++ {
		//col[j] = make(map[byte]int)
		col[j] = make([]int, 26)
		for i := 0; i < h; i++ {
			col[j][int(c[i][j]-'a')]++
		}
	}
	type node struct {
		k, v int
	}
	deleted := true
	for deleted {
		//for l := 0; l < Max(h, w); l++ {
		deleted = false
		//操作1
		del1 := make([]int, 26)
		for i := 0; i < h; i++ {
			//消せる行がある
			var cnt []node
			for j := 0; j < 26; j++ {
				if row[i][j] > 0 {
					cnt = append(cnt, node{j, row[i][j]})
				}
			}
			if len(cnt) == 1 {
				//for k, v := range row[i] {
				if cnt[0].v >= 2 {
					del1[cnt[0].k]++
					row[i][cnt[0].k] = 0
					//delete(row[i], k)
					deleted = true
				}
				//}
			}
		}
		//操作2
		del2 := make([]int, 26)
		for j := 0; j < w; j++ {
			//消せる列がある
			var cnt []node
			for k := 0; k < 26; k++ {
				if col[j][k] > 0 {
					cnt = append(cnt, node{k, col[j][k]})
				}
			}
			if len(cnt) == 1 {
				//for k, v := range col[j] {
				if cnt[0].v >= 2 {
					del2[cnt[0].k]++
					col[j][cnt[0].k] = 0
					//delete(col[j], k)
					deleted = true
				}
				//}
			}
		}
		//操作3
		for k, v := range del1 {
			for j := 0; j < w; j++ {
				if col[j][k] > 0 {
					col[j][k] -= v
				}
			}
		}
		for k, v := range del2 {
			for i := 0; i < h; i++ {
				if row[i][k] > 0 {
					row[i][k] -= v
				}
			}
		}
	}
	//for _, v := range row {
	//	fmt.Println(v)
	//}
	//fmt.Println(col)
	var ans int
	for i := range row {
		for _, v := range row[i] {
			if v > 0 {
				ans += v
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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
