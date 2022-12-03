package day02

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
		// TODO: Add test cases.
		{
			name:    "Part1 - should return 15",
			args:    args{[]string{"A Y", "B X", "C Z"}},
			want:    15,
			wantErr: false,
		},
		{
			name:    "Part1 - should return error",
			args:    args{[]string{"A Y", "B s", "C Z"}},
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
		// TODO: Add test cases.
		{
			name:    "Part2 - should return 12",
			args:    args{[]string{"A Y", "B X", "C Z"}},
			want:    12,
			wantErr: false,
		},
		{
			name:    "Part2 - should return error",
			args:    args{[]string{"A Y", "B s", "C Z"}},
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
