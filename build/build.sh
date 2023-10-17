# 需要 macOS 上编译
# clang -mno-red-zone -fno-asynchronous-unwind-tables -fno-builtin -fno-exceptions \
# -fno-rtti -fno-stack-protector -nostdlib -O3 -msse4 -mavx -mno-avx2 -DUSE_AVX=1 \
#  -DUSE_AVX2=0 -S *.c

# python ../asm2asm/asm2asm.py ../op.s ./op.s
python asm2asm/asm2asm.py ../simd.s ./simd.s
