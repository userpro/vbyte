package op

import (
	"math/rand"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func encodeVarint(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	return append(dAtA, uint8(v))
}

func decodeVarint(data []byte) (int64, int) {
	iNdEx := 0
	var v int64
	for shift := uint(0); ; shift += 7 {
		if shift >= 64 {
			return 0, 0
		}
		b := data[iNdEx]
		iNdEx++
		v |= int64(b&0x7F) << shift
		if b < 0x80 {
			break
		}
	}
	return v, iNdEx
}

func Test__decode(t *testing.T) {
	num := 256
	data := setupDecodeData(num)

	// simd 加速
	out1 := make([]uint32, num)
	cnt1 := __masked_vbyte_decode((*uint8)(unsafe.Pointer(&data[0])), &out1[0], uint64(num))

	// 传统方法
	out2 := make([]uint32, 0, num)
	cnt2 := 0
	for cnt2 < len(data) {
		res, idx := decodeVarint(data[cnt2:])
		cnt2 += idx
		out2 = append(out2, uint32(res))
	}

	for i := 0; i < num; i++ {
		assert.EqualValues(t, out1[i], out2[i])
	}
	t.Log(len(data), cnt1, cnt2)
	assert.EqualValues(t, cnt1, cnt2)
}

func setupDecodeData(num int) []byte {
	a := make([]int32, num, num)
	for i := range a {
		a[i] = rand.Int31()
	}

	data := make([]byte, 0, num*2)
	for _, v := range a {
		data = encodeVarint(data, uint64(v))
	}
	return data
}

const vbyteNum = 256

// BenchmarkDecodeRaw ...
func BenchmarkDecodeRaw(b *testing.B) {
	d := setupDecodeData(vbyteNum)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		idx := 0
		for j := 0; j < vbyteNum; j++ {
			_, idx = decodeVarint(d[idx:])
		}
	}
}

// BenchmarkDecodeSIMD ...
// func BenchmarkDecodeSIMD(b *testing.B) {
// 	d := setupDecodeData(vbyteNum)

// 	out := make([]uint64, vbyteNum)
// 	var cnt int32

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		idx := 0
// 		for j := 0; j < vbyteNum; j++ {
// 			__decode((*uint8)(unsafe.Pointer(&d[idx])), &out[0], &cnt)
// 			idx += int(cnt)
// 		}
// 	}
// }

// BenchmarkDecodeMaskedVbyteSIMD ...
func BenchmarkDecodeMaskedVbyteSIMD(b *testing.B) {
	d := setupDecodeData(vbyteNum)

	out := make([]uint32, vbyteNum)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = __masked_vbyte_decode((*uint8)(unsafe.Pointer(&d[0])), &out[0], vbyteNum)
	}
}

// BenchmarkDecodeCombineAPI ...
func BenchmarkDecodeCombineAPI(b *testing.B) {
	d := setupDecodeData(vbyteNum)

	out := make([]uint32, vbyteNum)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DecodeGroup((*uint8)(unsafe.Pointer(&d[0])), &out[0], vbyteNum)
	}
}
