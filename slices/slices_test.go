package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type E comparable

func TestContains_string(t *testing.T) {
	type args struct {
		slice []string
		item  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{
				slice: []string{"a", "b", "c"},
				item:  "b",
			},
			want: true,
		},
		{
			name: "false",
			args: args{
				slice: []string{"a", "b", "c"},
				item:  "d",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.want, Contains(tt.args.slice, tt.args.item))
		})
	}
}

func TestContains_int(t *testing.T) {
	type args struct {
		slice []int
		item  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{
				slice: []int{1, 2, 3},
				item:  2,
			},
			want: true,
		},
		{
			name: "false",
			args: args{
				slice: []int{1, 2, 3},
				item:  4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.want, Contains(tt.args.slice, tt.args.item))
		})
	}
}
