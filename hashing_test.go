package benchmark

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"hash"
	"testing"

	"github.com/cespare/xxhash"
	"github.com/dchest/siphash"
	blake2bsimd "github.com/minio/blake2b-simd"
	"github.com/spaolacci/murmur3"
	xxhash32 "github.com/vova616/xxhash"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

const hashBufferSize = 8

func benchmarkHash(b *testing.B, hash func() hash.Hash, length int64) {
	data := make([]byte, length)
	b.SetBytes(length)

	for i := 0; i < b.N; i++ {
		h := hash()
		_, err := h.Write(data[:])
		if err != nil {
			panic(err)
		}
		h.Sum(nil)
	}
}

func benchmarkHash64(b *testing.B, hash func() hash.Hash64, length int64) {
	data := make([]byte, length)
	b.SetBytes(length)

	for i := 0; i < b.N; i++ {
		h := hash()
		_, err := h.Write(data[:])
		if err != nil {
			panic(err)
		}
		h.Sum(nil)
	}
}
func benchmarkHash32(b *testing.B, hash func(uint32) hash.Hash32, length int64) {
	data := make([]byte, length)
	b.SetBytes(length)

	for i := 0; i < b.N; i++ {
		h := hash(1471)
		_, err := h.Write(data[:])
		if err != nil {
			panic(err)
		}
		h.Sum(nil)
	}
}

func benchmarkHashKeyError(b *testing.B, hash func([]byte) (hash.Hash, error), length int64) {
	data := make([]byte, length)
	b.SetBytes(length)
	key := make([]byte, 16)

	for i := 0; i < b.N; i++ {
		h, _ := hash(key)
		_, err := h.Write(data[:])
		if err != nil {
			panic(err)
		}
		h.Sum(nil)
	}
}

func benchmarkHashKey64(b *testing.B, hash func([]byte) hash.Hash64, length int64) {
	data := make([]byte, length)
	b.SetBytes(length)
	key := make([]byte, 16)

	for i := 0; i < b.N; i++ {
		h := hash(key)
		_, err := h.Write(data[:])
		if err != nil {
			panic(err)
		}
		h.Sum(nil)
	}
}
func benchmarkHashKey32(b *testing.B, hash func([]byte) hash.Hash32, length int64) {
	data := make([]byte, length)
	b.SetBytes(length)
	key := make([]byte, 8)

	for i := 0; i < b.N; i++ {
		h := hash(key)
		_, err := h.Write(data[:])
		if err != nil {
			panic(err)
		}
		h.Sum(nil)
	}
}

func BenchmarkComparisonMD5(b *testing.B) {
	benchmarkHash(b, md5.New, hashBufferSize)
}

func BenchmarkComparisonSHA1(b *testing.B) {
	benchmarkHash(b, sha1.New, hashBufferSize)
}

func BenchmarkComparisonSHA256(b *testing.B) {
	benchmarkHash(b, sha256.New, hashBufferSize)
}

func BenchmarkComparisonSHA3B224(b *testing.B) {
	benchmarkHash(b, sha3.New224, hashBufferSize)
}

func BenchmarkComparisonSHA3B256(b *testing.B) {
	benchmarkHash(b, sha3.New256, hashBufferSize)
}

func BenchmarkComparisonRIPEMD160(b *testing.B) {
	benchmarkHash(b, ripemd160.New, hashBufferSize)
}

func BenchmarkComparisonBlake2B(b *testing.B) {
	benchmarkHashKeyError(b, blake2b.New256, hashBufferSize)
}

func BenchmarkComparisonBlake2BSimd(b *testing.B) {
	benchmarkHash(b, blake2bsimd.New256, hashBufferSize)
}

func BenchmarkComparisonMurmur3(b *testing.B) {
	benchmarkHash64(b, murmur3.New64, hashBufferSize)
}
func BenchmarkComparisonSipHash(b *testing.B) {
	benchmarkHashKey64(b, siphash.New, hashBufferSize)
}
func BenchmarkComparisonXXHash64(b *testing.B) {
	benchmarkHash64(b, xxhash.New, hashBufferSize)
}
func BenchmarkComparisonXXHash32(b *testing.B) {
	benchmarkHash32(b, xxhash32.New, hashBufferSize)
}
