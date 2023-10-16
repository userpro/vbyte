package op

import (
	"unsafe"

	"github.com/klauspost/cpuid/v2"
)

func decodeVarintGroup(data *uint8, out *uint32, length uint64) (consumedBytes uint64) {
	dAtA := *(*[]byte)(unsafe.Pointer(&data))
	outRes := *(*[]uint32)(unsafe.Pointer(&out))
	for i := 0; i < int(length); i++ {
		var v uint32
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return
			}
			b := dAtA[consumedBytes]
			consumedBytes++
			v |= uint32(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		outRes[i] = v
	}
	return consumedBytes
}

// DecodeGroup 解码接口
var DecodeGroup = decodeVarintGroup

func init() {
	if cpuid.CPU.Supports(cpuid.SSE4) {
		DecodeGroup = __masked_vbyte_decode
	}
}
