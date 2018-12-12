package main

import "testing"

func Test_convertToBin(t *testing.T) {
	type args struct {
		n int
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
			if got := convertToBin(tt.args.n); got != tt.want {
				t.Errorf("convertToBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
