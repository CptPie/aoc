package day08

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
			name: "Part1 - should return",
			args: args{[]string{
				"",
			}},
			want:    0,
			wantErr: false,
		},
		{
			name: "Part1 - should return error",
			args: args{[]string{
				"",
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
			name: "Part2 - should return ",
			args: args{[]string{
				"",
			}},
			want:    0,
			wantErr: false,
		},
		{
			name: "Part2 - should return error",
			args: args{[]string{
				"",
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
