package internal

import (
	"reflect"
	"testing"
)

func TestCountSubString(t *testing.T) {
	type args struct {
		inputs []string
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "abaaba",
			args: args{
				inputs: []string{"abaaba"},
			},
			want: []int32{11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSubString(tt.args.inputs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountSubString() = %v, want %v", got, tt.want)
			}
		})
	}
}
