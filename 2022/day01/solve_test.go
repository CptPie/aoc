package day01

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
			name:    "First Elf Sum",
			args:    args{[]string{"1000", "2000", "3000", ""}},
			want:    6000,
			wantErr: false,
		},
		{
			name: "Find maximal Sum",
			args: args{[]string{
				"1000",
				"2000",
				"3000",
				"",
				"4000",
				"",
				"5000",
				"6000",
				"",
				"7000",
				"8000",
				"9000",
				"",
				"10000",
				""}},
			want:    24000,
			wantErr: false,
		},
		{
			name: "Fail with input error",
			args: args{[]string{
				"1000",
				"2000",
				"3000",
				"",
				"4000j",
				"",
				"5000",
				"6000",
				"",
				"7000",
				"8000",
				"9000",
				"",
				"10000",
				""}},
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
			name: "Sum of 3 max sums",
			args: args{[]string{
				"1000",
				"2000",
				"3000",
				"",
				"4000",
				"",
				"5000",
				"6000",
				"",
				"7000",
				"8000",
				"9000",
				"",
				"10000",
				""}},
			want:    45000,
			wantErr: false,
		},
		{
			name: "Fail with input error",
			args: args{[]string{
				"1000",
				"2000",
				"3000",
				"",
				"4000j",
				"",
				"5000",
				"6000",
				"",
				"7000",
				"8000",
				"9000",
				"",
				"10000",
				""}},
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
