package protocol

import (
	"reflect"
	"testing"
)

func TestUintToBytes(t *testing.T) {
	tests := []struct {
		name string
		i    uint64
		want []byte
	}{
		{
			name: "0",
			i:    0,
			want: []byte{0},
		},
		{
			name: "255",
			i:    255,
			want: []byte{255},
		},
		{
			name: "65535",
			i:    65535,
			want: []byte{255, 255},
		},
		{
			name: "65536",
			i:    65536,
			want: []byte{0, 1, 0, 0},
		},
		{
			name: "0xFFFFFFFFFFFFFFFe",
			i:    0xFFFFFFFFFFFFFFFe,
			want: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := uintToBytes(tt.i)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(b, tt.want) {
				t.Fatalf("got\n%+v\nwant\n%+v", b, tt.want)
			}
		})
	}
}

func TestLenForOpPushdata(t *testing.T) {
	tests := []struct {
		name   string
		opcode byte
		want   int
	}{
		{
			name:   "OP_PUSHDATA1",
			opcode: 0x4c,
			want:   1,
		},
		{
			name:   "OP_PUSHDATA2",
			opcode: 0x4d,
			want:   2,
		},
		{
			name:   "OP_PUSHDATA4",
			opcode: 0x4e,
			want:   4,
		},
		{
			name: "Fail low",
		},
		{
			name:   "Fail 0x4f",
			opcode: 0x4f,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lenForOpPushdata(tt.opcode)

			if got != tt.want {
				t.Fatalf("got\n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}
