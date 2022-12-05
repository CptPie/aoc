package day05

import (
	"2022/utils"
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

	testInput2, err := utils.ReadFile("test_input_2")
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
		want    string
		wantErr bool
	}{
		{
			name:    "TestData_simple",
			args:    args{testInput},
			want:    "CMZ",
			wantErr: false,
		},
		{
			name:    "TestData_complex",
			args:    args{testInput2},
			want:    "MPQFNHSTN",
			wantErr: false,
		},
		{
			name:    "RealData",
			args:    args{input},
			want:    "QPJPLMNNR",
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
		t.Errorf(err.Error())
	}

	input, err := utils.ReadFile("input")
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "TestData",
			args:    args{testInput},
			want:    "MCD",
			wantErr: false,
		},
		{
			name:    "RealData",
			args:    args{input},
			want:    "BQDNWJPVJ",
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
