package main

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example1",
			args: args{
				l1: &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
				l2: &ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			},
			want: &ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
		{
			name: "example2",
			args: args{
				l1: &ListNode{0, nil},
				l2: &ListNode{0, nil},
			},
			want: &ListNode{0, nil},
		},
		{
			// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
			// Output: [8,9,9,9,0,0,0,1]
			name: "example3",
			args: args{
				l1: &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}},
				l2: &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
			},
			want: &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !equal(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", toString(got), toString(tt.want))
			}
		})
	}
}

func equal(l1 *ListNode, l2 *ListNode) bool {
	for {
		if l1 == nil && l2 == nil {
			return true
		}
		if l1 == nil {
			return false
		}
		if l2 == nil {
			return false
		}
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
}

func toString(l *ListNode) string {
	ret := "["
	for {
		if l == nil {
			break
		}
		ret += fmt.Sprint(l.Val)
		ret += ", "
		l = l.Next
	}
	ret += "]"
	return ret
}
