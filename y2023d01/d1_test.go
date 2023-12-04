package y2023d01

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		inputPath   string
		logFilePath string
	}
	tests := []struct {
		name string
		args args
		want d1ResultStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args.inputPath, tt.args.logFilePath); got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCalibrationValues(t *testing.T) {
	type args struct {
		lines                []string
		includeNumbersAsText bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{lines: []string{
				"xt36five77",
				"two8five6zfrtjj",
				"eightthree8fiveqjgsdzgnnineeight",
				"7chmvlhnpfive",
				"1tcrgthmeight5mssseight",
				"eightoneqxspfzjk4zpfour",
				"fdbtmkhdfzrck9kxckbnft",
				"9six9",
				"goneightczdzjk18589",
				"41two3eightfscdmqjhdhtvsixld",
			}, includeNumbersAsText: true},
			want: []int{
				37,
				26,
				88,
				75,
				18,
				84,
				99,
				99,
				19,
				46,
			},
		}, {
			name: "test2",
			args: args{lines: []string{
				"t8929",
				"fourtwoxsxqqmqf3sixfoursixmmjhdlx",
				"bcbsfd14cjg",
				"95three6threendpqpjmbpcblone",
				"tdmvthreeonefive8574",
				"5eight82sixtwonev",
				"ninemg2shhmsqh",
				"thmlz4",
				"xtxjmm2tbbntrmdqxpkdjgh1vzjvtvg7nine",
				"vf19fourddfsvmzeight9",
			}, includeNumbersAsText: true},
			want: []int{
				89,
				46,
				14,
				91,
				34,
				51,
				92,
				44,
				29,
				19,
			},
		}, {
			name: "test3",
			args: args{lines: []string{"1four6ncdvzqjqhx1"}, includeNumbersAsText: true},
			want: []int{11},
		}, {
			name: "test4",
			args: args{lines: []string{"1bgqspl958lrj"}},
			want: []int{18},
		}, {
			name: "test5",
			args: args{lines: []string{"7nvmqrnthreejbzgnzvzpgkr69"}},
			want: []int{79},
		}, {
			name: "test6",
			args: args{lines: []string{"7576threesix"}},
			want: []int{76},
		}, {
			name: "test7",
			args: args{lines: []string{"twoc83pt"}},
			want: []int{23},
		}, {
			name: "test8",
			args: args{lines: []string{"fourkdnsvcq9sevendmhsdgt54threej"}},
			want: []int{43},
		}, {
			name: "test9",
			args: args{lines: []string{"zrjts8sixsix237flm"}},
			want: []int{87},
		}, {
			name: "test10",
			args: args{lines: []string{"8eightrndfour"}},
			want: []int{84},
		}, {
			name: "test11",
			args: args{lines: []string{"two9jsix5gcxf"}},
			want: []int{25},
		}, {
			name: "test12",
			args: args{lines: []string{"fivefour7nineseven1qtcdqbp1four"}},
			want: []int{54},
		}, {
			name: "test13",
			args: args{lines: []string{"fourzvkqhdninetwoftscrmsd64nxsgx"}},
			want: []int{44},
		}, {
			name: "test14",
			args: args{lines: []string{"q1tdsskthree"}},
			want: []int{13},
		}, {
			name: "test15",
			args: args{lines: []string{"mkhttggvjh9ctzffdqdjnheightninegmxqxhqrfqgbgzt"}},
			want: []int{99},
		}, {
			name: "test16",
			args: args{lines: []string{"ninep2fourf"}},
			want: []int{94},
		}, {
			name: "test17",
			args: args{lines: []string{"fiveeight2zxjpzffvdsevenjhjvjfiveone"}},
			want: []int{51},
		}, {
			name: "test18",
			args: args{lines: []string{"15737seven"}},
			want: []int{17},
		}, {
			name: "test19",
			args: args{lines: []string{"pdrss6oneone4fournine"}},
			want: []int{69},
		}, {
			name: "test20",
			args: args{lines: []string{"7b"}},
			want: []int{77},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCalibrationValues(tt.args.lines, tt.args.includeNumbersAsText); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCalibrationValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberTextToNumberAsString(t *testing.T) {
	type args struct {
		numberText string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberTextToNumberAsString(tt.args.numberText); got != tt.want {
				t.Errorf("numberTextToNumberAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getInput(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getInput(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
