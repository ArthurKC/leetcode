package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{"abcabcbb"},
			want: 3,
		},
		{
			name: "example2",
			args: args{"bbbbb"},
			want: 1,
		},
		{
			name: "example3",
			args: args{"pwwkew"},
			want: 3,
		},
		{
			name: "example4",
			args: args{""},
			want: 0,
		},
		{
			name: "wrong answer1",
			args: args{"aab"},
			want: 2,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
