package day02

import (
	"aoc/2023/utils"
	"testing"
)

func Test_solvePart1(t *testing.T) {
	type args struct {
		fileContents []string
	}

	testInput, err := utils.ReadFile("test_input")
	if err != nil {
		t.Errorf(err.Error())
	}

	input, err := utils.ReadFile("input")
	if err != nil {
		t.Errorf(err.Error())
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "TestData",
			args:    args{testInput},
			want:    8,
			wantErr: false,
		},
		{
			name:    "RealData",
			args:    args{input},
			want:    2439,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solvePart1(tt.args.fileContents)
			if (err != nil) != tt.wantErr {
				t.Errorf("solvePart1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("solvePart1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvePart2(t *testing.T) {
	type args struct {
		fileContents []string
	}

	testInput, err := utils.ReadFile("test_input")
	if err != nil {
		t.Fail()
	}

	input, err := utils.ReadFile("input")
	if err != nil {
		t.Fail()
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "TestData",
			args:    args{testInput},
			want:    2286,
			wantErr: false,
		},
		{
			name:    "RealData",
			args:    args{input},
			want:    63711,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solvePart2(tt.args.fileContents)
			if (err != nil) != tt.wantErr {
				t.Errorf("solvePart2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("solvePart2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
