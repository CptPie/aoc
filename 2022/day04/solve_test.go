package day04

import "testing"

func Test_solvePart1(t *testing.T) {
	type args struct {
		fileContents []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Part1 - should return 2",
			args: args{[]string{
				"2-4,6-8",
				"2-3,4-5",
				"5-7,7-9",
				"2-8,3-7",
				"6-6,4-6",
				"2-6,4-8",
			}},
			want:    2,
			wantErr: false,
		},
		{
			name: "Part1 - should return error",
			args: args{[]string{
				"2-4,6-8",
				"2-3,4-5",
				"5-7,7-9",
				"2-8,37",
				"6-6,4-6",
				"2-6,4-8",
			}},
			want:    0,
			wantErr: true,
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
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Part2 - should return 4",
			args: args{[]string{
				"2-4,6-8",
				"2-3,4-5",
				"5-7,7-9",
				"2-8,3-7",
				"6-6,4-6",
				"2-6,4-8",
			}},
			want:    4,
			wantErr: false,
		},
		{
			name: "Part2 - should return error",
			args: args{[]string{
				"2-4,6-8",
				"2-3,4-5",
				"5-7,7-9",
				"2-8,37",
				"6-6,4-6",
				"2-6,4-8",
			}},
			want:    0,
			wantErr: true,
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
