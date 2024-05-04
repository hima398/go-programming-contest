package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

const testCases = 100

/*
func TestMain(m *testing.M) {

	//scores = make([][7]int, testCases)

	code := m.Run()

	os.Exit(code)
}
*/

type args struct {
	n, m, k, turn int
	t, w, h, v    []int
}

type testCase struct {
	name string
	args args
	//ツールにのみ含まれる変数
	h, v, t, w, p []int
}

type File struct {
	buf [][]int
}

func (gate *File) PlayCard(c OutputCard) {
	gate.buf = append(gate.buf, []int{c.c, c.m})

	//ここから先はジャッジ内の状況を更新
	switch c.t {
	case RegularWork:
		judge.projects[c.m].remain -= c.w
		if judge.projects[c.m].remain <= 0 {
			judge.money += judge.projects[c.m].value
			//テスト出力
			//fmt.Println("Project ", judge.projects[c.m], " is finished.")
			//fmt.Println("money = ", judge.money)
			//
			judge.projects[c.m] = judge.remainedProjects[0]
			judge.remainedProjects = judge.remainedProjects[1:]
		}
	case HardWork:
		if c.m != 0 {
			panic(fmt.Sprintln("If HardWorkCard is played, m must be 0", c))
		}
		for i, v := range judge.projects {
			v.remain -= c.w
			if v.remain <= 0 {
				judge.money += v.value
				//テスト出力
				//fmt.Println("Project ", judge.projects[c.m], " is finished.")
				//fmt.Println("money = ", judge.money)
				//
				judge.projects[i] = judge.remainedProjects[0]
				judge.remainedProjects = judge.remainedProjects[1:]
			}
		}
	case Cancel:
		judge.projects[c.m] = judge.remainedProjects[0]
		judge.remainedProjects = judge.remainedProjects[1:]
	case Restructuring:
		if c.m != 0 {
			panic(fmt.Sprintln("If RestructuringCard is played, m must be 0", c))
		}

		for i := 0; i < judge.m; i++ {
			judge.projects[i] = judge.remainedProjects[0]
			judge.remainedProjects = judge.remainedProjects[1:]
		}
	case Investment:
		if c.m != 0 {
			panic(fmt.Sprintln("If InvestmentCard is played, m must be 0", c))
		}

		judge.usedInvestmentCards++
		for i := range judge.deck {
			judge.deck[i].w *= 2
			judge.deck[i].cost *= 2
		}
		for i := range judge.remainedProjects {
			judge.remainedProjects[i].remain *= 2
			judge.remainedProjects[i].value *= 2
		}
	}
}

// 最新のプロジェクトの情報を受け取る
func (gate *File) SubscribeProject(m int) []Project {
	return judge.projects
}

// 最新の所持金の情報を受け取る
func (gate *File) SubscribeMoney() int {
	return judge.money
}

// 方針カードの候補を受け取る
func (gate *File) DealCard(k int) []DealedCard {
	var res []DealedCard
	for i := 0; i < k; i++ {
		res = append(res, judge.deck[0])
		judge.deck = judge.deck[1:]
	}
	return res
}

// 候補から引く方針カードを宣言する
func (gate *File) DeclareCard(r int, c DealedCard) {
	gate.buf = append(gate.buf, []int{r})
	judge.money -= c.cost
	if judge.money < 0 {
		panic("")
	}
	fmt.Println("# judge.money = ", judge.money)
}

type Score struct {
	name  string
	score int
}

func writeScores(fileName string, s []Score) {
	fp, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	for _, v := range s {
		fmt.Println(v)
		fmt.Fprintf(fp, "%s, %d\n", v.name, v.score)
	}
}

// ツールで用いられる入出力ファイルの使用に従ってファイルを読み込む
func readInput(dir string) []testCase {
	//files, err := ioutil.ReadDir(dir)
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	var res []testCase

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		fp, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		sc := bufio.NewScanner(fp)
		sc.Split(bufio.ScanWords)

		readInt := func() int {
			sc.Scan()
			i, _ := strconv.Atoi(sc.Text())
			return int(i)
		}

		n, m, k, turn := readInt(), readInt(), readInt(), readInt()
		var h, v []int
		for i := 0; i < m; i++ {
			h = append(h, readInt())
			v = append(v, readInt())
		}
		var h2, v2 []int

		for i := 0; i < turn; i++ {
			for j := 0; j < m; j++ {
				h2 = append(h2, readInt())
				v2 = append(v2, readInt())
			}
		}
		var t, w []int
		for i := 0; i < n; i++ {
			t = append(t, readInt())
			w = append(w, readInt())
		}
		var t2, w2, p []int

		for i := 0; i < turn; i++ {
			for j := 0; j < k; j++ {
				t2 = append(t2, readInt())
				w2 = append(w2, readInt())
				p = append(p, readInt())
			}
		}

		res = append(res, testCase{file.Name(), args{n, m, k, turn, t, w, h, v}, h2, v2, t2, w2, p})
	}
	return res
}

func writeOutput(dir, name string, ans [][]int) {
	path := filepath.Join(dir, name)
	fp, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	for i, o := range ans {
		//fmt.Fprintf(fp, "%d %d %d\n", o.y, o.x, o.p)
		fmt.Fprintln(fp, "## turn = ", i)
		//if len(o) == 1 {
		fmt.Fprintf(fp, "%d %d\n", o[0], o[1])
		fmt.Fprintf(fp, "%d\n", o[2])
		//} else if len(o) == 2 {
		//}
	}
}

type Judge struct {
	m                   int // プロジェクトの個数
	money               int
	usedInvestmentCards int
	projects            []Project

	deck             []DealedCard
	remainedProjects []Project
}

var judge *Judge

func Test_solve02(t *testing.T) {
	tests := readInput("./in")

	gateway = new(File)

	var scores []Score
	for _, tt := range tests {
		judge = new(Judge)
		judge.m = tt.args.m
		judge.money = 0
		judge.usedInvestmentCards = 0
		judge.projects = createProjects(tt.args.h, tt.args.v)
		judge.remainedProjects = createProjects(tt.h, tt.v)
		judge.deck = createDeck(tt.t, tt.w, tt.p)
		var ans [][]int
		var score int
		//fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			ans, score = solve02(tt.args.n, tt.args.m, tt.args.k, tt.args.turn, tt.args.t, tt.args.w, tt.args.h, tt.args.v)
			writeOutput("./out", tt.name, ans)
		})
		scores = append(scores, Score{tt.name, score})
	}
	//fmt.Println(scores)
	sort.Slice(scores, func(i, j int) bool {
		if scores[i].score == scores[j].score {
			return scores[i].name < scores[j].name
		}
		return scores[i].score > scores[j].score
	})
	writeScores("./scores.txt", scores)

}

func Test_solve01(t *testing.T) {
	t.Skip()
	tests := readInput("./in")

	gateway = new(File)

	var scores []Score
	for _, tt := range tests {
		judge = new(Judge)
		judge.m = tt.args.m
		judge.money = 0
		judge.usedInvestmentCards = 0
		judge.projects = createProjects(tt.args.h, tt.args.v)
		judge.remainedProjects = createProjects(tt.h, tt.v)
		judge.deck = createDeck(tt.t, tt.w, tt.p)
		var ans [][]int
		var score int
		//fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			ans, score = solve01(tt.args.n, tt.args.m, tt.args.k, tt.args.turn, tt.args.t, tt.args.w, tt.args.h, tt.args.v)
			writeOutput("./out", tt.name, ans)
		})
		scores = append(scores, Score{tt.name, score})
	}
	//fmt.Println(scores)
	sort.Slice(scores, func(i, j int) bool {
		if scores[i].score == scores[j].score {
			return scores[i].name < scores[j].name
		}
		return scores[i].score > scores[j].score
	})
	writeScores("./scores.txt", scores)

}

func TestDivideSlice(t *testing.T) {
	type args struct {
		A []PolicyCard
		K int
	}
	tests := []struct {
		name    string
		args    args
		want    []PolicyCard
		want1   []PolicyCard
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{[]PolicyCard{{0, 0, false}, {1, 1, false}, {2, 2, false}}, 0}, []PolicyCard{}, []PolicyCard{{1, 1, false}, {2, 2, false}}, false},
		{"", args{[]PolicyCard{{0, 0, false}, {1, 1, false}, {2, 2, false}}, 1}, []PolicyCard{{0, 0, false}}, []PolicyCard{{2, 2, false}}, false},
		{"", args{[]PolicyCard{{0, 0, false}, {1, 1, false}, {2, 2, false}}, 2}, []PolicyCard{{0, 0, false}, {1, 1, false}}, []PolicyCard{}, false},
		{"", args{[]PolicyCard{{0, 0, false}, {1, 1, false}, {2, 2, false}}, -1}, nil, nil, true},
		{"", args{[]PolicyCard{{0, 0, false}, {1, 1, false}, {2, 2, false}}, 3}, nil, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := DividePolicyCard(tt.args.A, tt.args.K)
			if (err != nil) != tt.wantErr {
				t.Errorf("DivideSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DivideSlice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DivideSlice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHands_hasRestructuring(t *testing.T) {
	type fields struct {
		n     int
		cs    []PolicyCard
		kinds []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"1枚", fields{0, nil, []int{0, 0, 0, 1, 0}}, true},
		{"1枚以上", fields{0, nil, []int{0, 0, 0, 2, 0}}, true},
		{"0枚", fields{0, nil, []int{0, 0, 0, 0, 0}}, false},
		{"違うカードを数えていないか確認", fields{0, nil, []int{1, 1, 1, 0, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hands{
				n:     tt.fields.n,
				cs:    tt.fields.cs,
				kinds: tt.fields.kinds,
			}
			if got := h.hasRestructuring(); got != tt.want {
				t.Errorf("Hands.hasRestructuring() = %v, want %v", got, tt.want)
			}
		})
	}
}
