package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

func TestAnySliceToTypedSlice(t *testing.T) {
	type args struct {
		input any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Convert []any with int values to []int",
			args: args{input: []any{1, 2, 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "Convert []any with float64 values to []float64",
			args: args{input: []any{1.1, 2.2, 3.3}},
			want: []float64{1.1, 2.2, 3.3},
		},
		{
			name: "Convert []any with string values to []string",
			args: args{input: []any{"apple", "banana", "cherry"}},
			want: []string{"apple", "banana", "cherry"},
		},
		{
			name: "Convert []any with time.Time values to []time.Time",
			args: args{
				input: []any{
					time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2025, time.February, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			want: []time.Time{
				time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2025, time.February, 2, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Convert empty []any to empty []string (default behavior)",
			args: args{input: []any{}},
			want: []string{}, // Assuming default empty case returns []string{}
		},
		{
			name: "Convert single-element []any to correct type",
			args: args{input: []any{42}},
			want: []int{42}, // Should infer as []int
		},
		{
			name: "Convert mixed-type []any should remain unchanged",
			args: args{input: []any{"apple", 1, 2.3}},
			want: []any{"apple", 1, 2.3}, // If types are mixed, no conversion should happen
		},
		{
			name: "Convert non-slice input should remain unchanged",
			args: args{input: "not a slice"},
			want: "not a slice", // If not a slice, should return as-is
		},
		{
			name: "Convert nil input should return nil",
			args: args{input: nil},
			want: nil,
		},
		{
			name: "Convert []any with empty values should return empty slice",
			args: args{input: []any{}},
			want: []string{}, // Should return an empty slice, not nil
		},
		{
			name: "Convert []any with bool values to []bool",
			args: args{input: []any{true, false, true}},
			want: []bool{true, false, true},
		},
		{
			name: "Convert []any with complex numbers to []complex64",
			args: args{input: []any{complex(1, 2), complex(3, 4)}},
			want: []complex128{complex(1, 2), complex(3, 4)},
		},
		{
			name: "Convert []any with byte values to []byte",
			args: args{input: []any{byte(65), byte(66), byte(67)}},
			want: []byte{65, 66, 67}, // Should recognize byte values
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, AnySliceToTypedSlice(tt.args.input), "AnySliceToTypedSlice(%v)", tt.args.input)
		})
	}
}
