package op

import (
	"reflect"
	"testing"
	"unsafe"
)

func Test_decodeVarintGroup(t *testing.T) {
	var out uint32
	data := encodeVarint(nil, 123)
	data = append(data, encodeVarint(nil, 12345)...)
	type args struct {
		data   *uint8
		out    *uint32
		length uint64
	}
	tests := []struct {
		name         string
		args         args
		wantConsumed uint64
		want         []uint32
	}{
		{
			name: "test1",
			args: args{
				data:   &data[0],
				out:    &out,
				length: 2,
			},
			wantConsumed: 3,
			want:         []uint32{123, 12345},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotConsumed := decodeVarintGroup(tt.args.data, tt.args.out, tt.args.length); gotConsumed != tt.wantConsumed || !reflect.DeepEqual(*(*[]uint32)(unsafe.Pointer(&tt.args.out)), tt.want) {
				t.Errorf("decodeVarintGroup() = %v, want %v; out = %v, want %v", gotConsumed, tt.wantConsumed, *tt.args.out, tt.want)
			}
		})
	}
}
