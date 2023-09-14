package op

import "testing"

func Test_decodeVarintGroup(t *testing.T) {
	var out uint32
	data := encodeVarint(nil, 123)
	type args struct {
		data   *uint8
		out    *uint32
		length uint64
	}
	tests := []struct {
		name         string
		args         args
		wantConsumed uint64
		want         uint32
	}{
		{
			name: "test1",
			args: args{
				data:   &data[0],
				out:    &out,
				length: 1,
			},
			wantConsumed: 1,
			want:         123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotConsumed := decodeVarintGroup(tt.args.data, tt.args.out, tt.args.length); gotConsumed != tt.wantConsumed || *tt.args.out != tt.want {
				t.Errorf("decodeVarintGroup() = %v, want %v; out = %v, want %v", gotConsumed, tt.wantConsumed, *tt.args.out, tt.want)
			}
		})
	}
}
