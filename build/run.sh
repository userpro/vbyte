clang -fPIC -msse4 -std=c99 -O3 -Wall -Wextra -pedantic -Wshadow -I../include simd.c ../src/varintencode.c && ./a.out