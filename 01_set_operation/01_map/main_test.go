package main

import (
	"reflect"
	"testing"
)

func TestGetIntersectionOfSets(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Get Intersection A and B",
			args: args{
				l1: []int{1, 2, 3, 4, 5},
				l2: []int{3, 4, 5, 6, 7},
			},
			want: []int{3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionOfSets(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionOfSets = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUnionOfSets(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Get Union A and B",
			args: args{
				l1: []int{1, 2, 3, 4, 5},
				l2: []int{3, 4, 5, 6, 7},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getUnionOfSets(tt.args.l1, tt.args.l2)
			if result := equalSlice(got, tt.want); !result {
				t.Errorf("getUnionOfSets = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSymmetricDifferenceOfSets(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Get Difference A and B",
			args: args{
				l1: []int{1, 2, 3, 4, 5},
				l2: []int{3, 4, 5, 6, 7},
			},
			want: []int{1, 2, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSymmetricDifferenceOfSets(tt.args.l1, tt.args.l2)
			if result := equalSlice(got, tt.want); !result {
				t.Errorf("getSymmetricDifferenceOfSets = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDifferenceOfSets(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Get Difference A and B",
			args: args{
				l1: []int{1, 2, 3, 4, 5},
				l2: []int{3, 4, 5, 6, 7},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getDifferenceOfSets(tt.args.l1, tt.args.l2)
			if result := equalSlice(got, tt.want); !result {
				t.Errorf("getDifferenceOfSets = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalSlice(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	m1 := make(map[int]struct{}, len(s1))
	m2 := make(map[int]struct{}, len(s2))

	for _, data := range s1 {
		m1[data] = struct{}{}
	}
	for _, data := range s2 {
		m2[data] = struct{}{}
	}

	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v2 != v1 {
			return false
		}
	}
	return true
}
