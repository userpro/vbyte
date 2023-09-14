package op

import (
	"unsafe"

	"github.com/klauspost/cpuid/v2"
)

func decodeVarintGroup(data *uint8, out *uint32, length uint64) (consumedBytes uint64) {
	dAtA := *(*[]byte)(unsafe.Pointer(&data))
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
	*out = v
	return consumedBytes
}

// VbyteDecodeGroup 解码接口
var VbyteDecodeGroup = decodeVarintGroup

func init() {
	// 检测CPU是否支持AVX512
	if cpuid.CPU.Supports(cpuid.AVX2) {
		VbyteDecodeGroup = __masked_vbyte_decode
	}
}
