package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppendUnique(t *testing.T) {
	type args[T comparable] struct {
		slice1 []T
		slice2 []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "empty slices",
			args: args[string]{slice1: []string{}, slice2: []string{}},
			want: []string{},
		},
		{
			name: "empty slice1",
			args: args[string]{slice1: []string{}, slice2: []string{"a", "b"}},
			want: []string{"a", "b"},
		},
		{
			name: "empty slice2",
			args: args[string]{slice1: []string{"a", "b"}, slice2: []string{}},
			want: []string{"a", "b"},
		},
		{
			name: "no duplicates",
			args: args[string]{slice1: []string{"a", "b"}, slice2: []string{"c", "d"}},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "duplicates",
			args: args[string]{slice1: []string{"a", "b"}, slice2: []string{"b", "c"}},
			want: []string{"a", "b", "c"},
		},
		{
			name: "duplicates in both slices",
			args: args[string]{slice1: []string{"a", "b", "c"}, slice2: []string{"b", "c", "d"}},
			want: []string{"a", "b", "c", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, AppendSliceUnique(tt.args.slice1, tt.args.slice2), "AppendSliceUnique(%v, %v)", tt.args.slice1, tt.args.slice2)
		})
	}
}
