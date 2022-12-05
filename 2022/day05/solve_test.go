package day05

import (
	"2022/utils"
	"testing"
)

func Test_solvePart1(t *testing.T) {
	type args struct {
		fileContents []string
	}
	input, err := utils.ReadFile("test_input")
	if err != nil {
		t.Fail()
	}
	input2, err := utils.ReadFile("test_input_2")
	if err != nil {
		t.Fail()
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Simple",
			args:    args{input},
			want:    "CMZ",
			wantErr: false,
		},
		{
			name:    "Complex",
			args:    args{input2},
			want:    "MPQFNHSTN",
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
	input, err := utils.ReadFile("test_input")
	if err != nil {
		t.Fail()
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Part2 - should return ",
			args:    args{input},
			want:    "MCD",
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
