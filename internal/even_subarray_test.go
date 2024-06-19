package internal

import "testing"

func TestEvenSubArray(t *testing.T) {
	type args struct {
		nums []int32
		k    int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "[1,2,3,4], k:8",
			args: args{
				nums: []int32{1, 2, 3, 4},
				k:    1,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvenSubArray(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("EvenSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
