package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{"babad"},
			want: "bab",
		},
		{
			name: "example2",
			args: args{"cbbd"},
			want: "bb",
		},
		{
			name: "example3",
			args: args{"a"},
			want: "a",
		},
		{
			name: "example4",
			args: args{"ac"},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
