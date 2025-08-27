package ohsnap

import (
	"math"
	"math/rand/v2"
	"unicode"
)

type Builder struct {
	rnd *rand.Rand

	minByte byte
	maxByte byte

	minRune rune
	maxRune rune

	minInt int
	maxInt int

	minInt8 int8
	maxInt8 int8

	minInt16 int16
	maxInt16 int16

	minInt32 int32
	maxInt32 int32

	minInt64 int64
	maxInt64 int64

	minUint uint
	maxUint uint

	minUint8 uint8
	maxUint8 uint8

	minUint16 uint16
	maxUint16 uint16

	minUint32 uint32
	maxUint32 uint32

	minUint64 uint64
	maxUint64 uint64

	minFloat32 float32
	maxFloat32 float32

	minFloat64 float64
	maxFloat64 float64

	minSliceLen int
	maxSliceLen int
}

func NewBuilder(rnd *rand.Rand) *Builder {
	return &Builder{
		rnd: rnd,

		minByte: 0,
		maxByte: math.MaxUint8,

		minRune: 0,
		maxRune: unicode.MaxRune,

		minInt: 0,
		maxInt: math.MaxInt,

		minInt8: 0,
		maxInt8: math.MaxInt8,

		minInt16: 0,
		maxInt16: math.MaxInt16,

		minInt32: 0,
		maxInt32: math.MaxInt32,

		minInt64: 0,
		maxInt64: math.MaxInt64,

		minUint: 0,
		maxUint: math.MaxUint,

		minUint8: 0,
		maxUint8: math.MaxUint8,

		minUint16: 0,
		maxUint16: math.MaxUint16,

		minUint32: 0,
		maxUint32: math.MaxUint32,

		minUint64: 0,
		maxUint64: math.MaxUint64,

		minFloat32: 0,
		maxFloat32: math.MaxFloat32,

		minFloat64: 0,
		maxFloat64: math.MaxFloat64,

		minSliceLen: 0,
		maxSliceLen: 10,
	}
}

func (b *Builder) copy() *Builder {
	bb := *b
	return &bb
}

// methods to change config

func (b *Builder) Rand(rnd *rand.Rand) *Builder {
	bb := b.copy()
	bb.rnd = rnd
	return bb
}

func (b *Builder) MinByte(x byte) *Builder {
	bb := b.copy()
	bb.minByte = x
	return bb
}

func (b *Builder) MaxByte(x byte) *Builder {
	bb := b.copy()
	bb.maxByte = x
	return bb
}

func (b *Builder) MinRune(x rune) *Builder {
	bb := b.copy()
	bb.minRune = x
	return bb
}

func (b *Builder) MaxRune(x rune) *Builder {
	bb := b.copy()
	bb.maxRune = x
	return bb
}

func (b *Builder) MinInt(x int) *Builder {
	bb := b.copy()
	bb.minInt = x
	return bb
}

func (b *Builder) MaxInt(x int) *Builder {
	bb := b.copy()
	bb.maxInt = x
	return bb
}

func (b *Builder) MinInt8(x int8) *Builder {
	bb := b.copy()
	bb.minInt8 = x
	return bb
}

func (b *Builder) MaxInt8(x int8) *Builder {
	bb := b.copy()
	bb.maxInt8 = x
	return bb
}

func (b *Builder) MinInt16(x int16) *Builder {
	bb := b.copy()
	bb.minInt16 = x
	return bb
}

func (b *Builder) MaxInt16(x int16) *Builder {
	bb := b.copy()
	bb.maxInt16 = x
	return bb
}

func (b *Builder) MinInt32(x int32) *Builder {
	bb := b.copy()
	bb.minInt32 = x
	return bb
}

func (b *Builder) MaxInt32(x int32) *Builder {
	bb := b.copy()
	bb.maxInt32 = x
	return bb
}

func (b *Builder) MinInt64(x int64) *Builder {
	bb := b.copy()
	bb.minInt64 = x
	return bb
}

func (b *Builder) MaxInt64(x int64) *Builder {
	bb := b.copy()
	bb.maxInt64 = x
	return bb
}

func (b *Builder) MinUint(x uint) *Builder {
	bb := b.copy()
	bb.minUint = x
	return bb
}

func (b *Builder) MaxUint(x uint) *Builder {
	bb := b.copy()
	bb.maxUint = x
	return bb
}

func (b *Builder) MinUint8(x uint8) *Builder {
	bb := b.copy()
	bb.minUint8 = x
	return bb
}

func (b *Builder) MaxUint8(x uint8) *Builder {
	bb := b.copy()
	bb.maxUint8 = x
	return bb
}

func (b *Builder) MinUint16(x uint16) *Builder {
	bb := b.copy()
	bb.minUint16 = x
	return bb
}

func (b *Builder) MaxUint16(x uint16) *Builder {
	bb := b.copy()
	bb.maxUint16 = x
	return bb
}

func (b *Builder) MinUint32(x uint32) *Builder {
	bb := b.copy()
	bb.minUint32 = x
	return bb
}

func (b *Builder) MaxUint32(x uint32) *Builder {
	bb := b.copy()
	bb.maxUint32 = x
	return bb
}

func (b *Builder) MinUint64(x uint64) *Builder {
	bb := b.copy()
	bb.minUint64 = x
	return bb
}

func (b *Builder) MaxUint64(x uint64) *Builder {
	bb := b.copy()
	bb.maxUint64 = x
	return bb
}

func (b *Builder) MinFloat32(x float32) *Builder {
	bb := b.copy()
	bb.minFloat32 = x
	return bb
}

func (b *Builder) MaxFloat32(x float32) *Builder {
	bb := b.copy()
	bb.maxFloat32 = x
	return bb
}

func (b *Builder) MinFloat64(x float64) *Builder {
	bb := b.copy()
	bb.minFloat64 = x
	return bb
}

func (b *Builder) MaxFloat64(x float64) *Builder {
	bb := b.copy()
	bb.maxFloat64 = x
	return bb
}

func (b *Builder) MinSliceLen(x int) *Builder {
	bb := b.copy()
	bb.minSliceLen = x
	return bb
}

func (b *Builder) MaxSliceLen(x int) *Builder {
	bb := b.copy()
	bb.maxSliceLen = x
	return bb
}

// methods to to generate arbitrary primitives

func (b *Builder) Bool() Arbitrary[bool] {
	return ArbitraryBool(b.rnd)
}

func (b *Builder) Byte() Arbitrary[byte] {
	return ArbitraryByte(b.rnd, b.minByte, b.maxByte)
}

func (b *Builder) Rune() Arbitrary[rune] {
	return ArbitraryRune(b.rnd, b.minRune, b.maxRune)
}

func (b *Builder) Int() Arbitrary[int] {
	return ArbitraryInt(b.rnd, b.minInt, b.maxInt)
}

func (b *Builder) Int8() Arbitrary[int8] {
	return ArbitraryInt8(b.rnd, b.minInt8, b.maxInt8)
}

func (b *Builder) Int16() Arbitrary[int16] {
	return ArbitraryInt16(b.rnd, b.minInt16, b.maxInt16)
}

func (b *Builder) Int32() Arbitrary[int32] {
	return ArbitraryInt32(b.rnd, b.minInt32, b.maxInt32)
}

func (b *Builder) Int64() Arbitrary[int64] {
	return ArbitraryInt64(b.rnd, b.minInt64, b.maxInt64)
}

func (b *Builder) Uint() Arbitrary[uint] {
	return ArbitraryUint(b.rnd, b.minUint, b.maxUint)
}

func (b *Builder) Uint8() Arbitrary[uint8] {
	return ArbitraryUint8(b.rnd, b.minUint8, b.maxUint8)
}

func (b *Builder) Uint16() Arbitrary[uint16] {
	return ArbitraryUint16(b.rnd, b.minUint16, b.maxUint16)
}

func (b *Builder) Uint32() Arbitrary[uint32] {
	return ArbitraryUint32(b.rnd, b.minUint32, b.maxUint32)
}

func (b *Builder) Uint64() Arbitrary[uint64] {
	return ArbitraryUint64(b.rnd, b.minUint64, b.maxUint64)
}

func (b *Builder) Float32() Arbitrary[float32] {
	return ArbitraryFloat32(b.rnd, b.minFloat32, b.maxFloat32)
}

func (b *Builder) Float64() Arbitrary[float64] {
	return ArbitraryFloat64(b.rnd, b.minFloat64, b.maxFloat64)
}

// methods to to generate arbitrary slices of primitives

func (b *Builder) BoolSlice() Arbitrary[[]bool] {
	return ArbitrarySlice(b.rnd, b.Bool(), b.minSliceLen, b.maxSliceLen)
}

func (b *Builder) ByteSlice() Arbitrary[[]byte] {
	return ArbitrarySlice(b.rnd, b.Byte(), b.minSliceLen, b.maxSliceLen)
}

func (b *Builder) RuneSlice() Arbitrary[[]rune] {
	return ArbitrarySlice(b.rnd, b.Rune(), b.minSliceLen, b.maxSliceLen)
}

func (b *Builder) IntSlice() Arbitrary[[]int] {
	return ArbitrarySlice(b.rnd, b.Int(), b.minSliceLen, b.maxSliceLen)
}

func (b *Builder) Int8Slice() Arbitrary[[]int8] {
	return ArbitrarySlice(b.rnd, b.Int8(), b.minSliceLen, b.maxSliceLen)
}

func (b *Builder) Int16Slice() Arbitrary[[]int16] {
	return ArbitrarySlice(b.rnd, b.Int16(), b.minSliceLen, b.maxSliceLen)
}

func (b *Builder) Int32Slice() Arbitrary[[]int32] {
	return ArbitrarySlice(b.rnd, b.Int32(), b.minSliceLen, b.maxSliceLen)
}

func (b *Builder) Int64Slice() Arbitrary[[]int64] {
	return ArbitrarySlice(b.rnd, b.Int64(), b.minSliceLen, b.maxSliceLen)
}

func (b *Builder) UintSlice() Arbitrary[[]uint] {
	return ArbitrarySlice(b.rnd, b.Uint(), b.minSliceLen, b.maxSliceLen)
}

func (b *Builder) Uint8Slice() Arbitrary[[]uint8] {
	return ArbitrarySlice(b.rnd, b.Uint8(), b.minSliceLen, b.maxSliceLen)
}

func (b *Builder) Uint16Slice() Arbitrary[[]uint16] {
	return ArbitrarySlice(b.rnd, b.Uint16(), b.minSliceLen, b.maxSliceLen)
}

func (b *Builder) Uint32Slice() Arbitrary[[]uint32] {
	return ArbitrarySlice(b.rnd, b.Uint32(), b.minSliceLen, b.maxSliceLen)
}

func (b *Builder) Uint64Slice() Arbitrary[[]uint64] {
	return ArbitrarySlice(b.rnd, b.Uint64(), b.minSliceLen, b.maxSliceLen)
}

func (b *Builder) RuneFromTable(tbl *unicode.RangeTable) Arbitrary[rune] {
	return RuneFromTable(b.rnd, tbl)
}
