package gocommon

import (
	"math"
	"strconv"
)

// Rotli - Rotate left int
func Rotli(value int, count uint) int {
	return (value << count) | (value >> (strconv.IntSize - count))
}

// Rotlu - Rotate left uint
func Rotlu(value uint, count uint) uint {
	return (value << count) | (value >> (strconv.IntSize - count))
}

// Rotli32 - Rotate left int32
func Rotli32(value int32, count uint32) int32 {
	return (value << count) | (value >> (32 - count))
}

// Rotlu32 - Rotate left uint32
func Rotlu32(value uint32, count uint32) uint32 {
	return (value << count) | (value >> (32 - count))
}

// Rotli64 - Rotate left int64
func Rotli64(value int64, count uint32) int64 {
	return (value << count) | (value >> (64 - count))
}

// Rotlu64 - Rotate left uint64
func Rotlu64(value uint64, count uint32) uint64 {
	return (value << count) | (value >> (64 - count))
}

// Rotri - Rotate right int
func Rotri(value int, count uint) int {
	return (value >> count) | (value << (strconv.IntSize - count))
}

// Rotru - Rotate right uint
func Rotru(value uint, count uint) uint {
	return (value >> count) | (value << (strconv.IntSize - count))
}

// Rotri32 - Rotate right int32
func Rotri32(value int32, count uint32) int32 {
	return (value >> count) | (value << (32 - count))
}

// Rotru32 - Rotate right uint32
func Rotru32(value uint32, count uint32) uint32 {
	return (value >> count) | (value << (32 - count))
}

// Rotri64 - Rotate right int64
func Rotri64(value int64, count uint32) int64 {
	return (value >> count) | (value << (64 - count))
}

// Rotru64 - Rotate right uint64
func Rotru64(value uint64, count uint32) uint64 {
	return (value >> count) | (value << (64 - count))
}

// Decompf32 - Decompose a float32 value int two numbers. Given the following:
//
//	3.141592
//
// The return values will be:
//
//	base = 3
//	frac = 141592
func Decompf32(val float32) (base, frac int32) {
	b, f := math.Modf(float64(val))
	return int32(b), int32(f)
}

// Decompf64 - Decompose a float64 value int two numbers. Given the following:
//
//	3.141592
//
// The return values will be:
//
//	base = 3
//	frac = 141592
func Decompf64(val float64) (base, frac int64) {
	b, f := math.Modf(val)
	return int64(b), int64(f)
}

// Makef32 - Make a float32 variable from two numbers. Given the following:
//
//	base = 3
//	frac = 141592
//
// The return float64 value will be:
//
//	3.14159
func Makef32(base, frac int32) float32 {
	b := float32(base)
	f := float32(frac)
	for f > 1.0 {
		f /= 10.0
	}
	return b + f
}

// Makef64 - Make a float64 variable from two numbers. Given the following:
//
//	base = 3
//	frac = 141592
//
// The return float64 value will be:
//
//	3.14159
func Makef64(base, frac int64) float64 {
	b := float64(base)
	f := float64(frac)
	for f > 1.0 {
		f /= 10.0
	}
	return b + f
}

func getRangei16(minVal, maxVal int16) int16 {
	rng := (maxVal - minVal)
	if rng == 0 {
		return 1
	}
	return rng
}

func getRangeu16(minVal, maxVal uint16) uint16 {
	rng := (maxVal - minVal)
	if rng == 0 {
		return 1
	}
	return rng
}

func getRangei32(minVal, maxVal int32) int32 {
	rng := (maxVal - minVal)
	if rng == 0 {
		return 1
	}
	return rng
}

func getRangeu32(minVal, maxVal uint32) uint32 {
	rng := (maxVal - minVal)
	if rng == 0 {
		return 1
	}
	return rng
}

func getRangei64(minVal, maxVal int64) int64 {
	rng := (maxVal - minVal)
	if rng == 0 {
		return 1
	}
	return rng
}

func getRangeu64(minVal, maxVal uint64) uint64 {
	rng := (maxVal - minVal)
	if rng == 0 {
		return 1
	}
	return rng
}

func getRangef32(minVal, maxVal float32) float32 {
	rng := (maxVal - minVal)
	if rng == 0 {
		return 1.0
	}
	return rng
}

func getRangef64(minVal, maxVal float64) float64 {
	rng := (maxVal - minVal)
	if rng == 0 {
		return 1.0
	}
	return rng
}

// Mini - Return the smaller int of x or y
func Mini(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Minu - Return the smaller uint of x or y
func Minu(x, y uint) uint {
	if x < y {
		return x
	}
	return y
}

// Mini32 - Return the smaller int32 of x or y
func Mini32(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

// Minu32 - Return the smaller uint32 of x or y
func Minu32(x, y uint32) uint32 {
	if x < y {
		return x
	}
	return y
}

// Mini64 - Return the smaller int64 of x or y
func Mini64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

// Minu64 - Return the smaller uint64 of x or y
func Minu64(x, y uint64) uint64 {
	if x < y {
		return x
	}
	return y
}

// Maxi - Return the larger int of x or y
func Maxi(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Maxu - Return the larger uint of x or y
func Maxu(x, y uint) uint {
	if x > y {
		return x
	}
	return y
}

// Maxi32 - Return the larger int32 of x or y
func Maxi32(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

// Maxu32 - Return the larger uint32 of x or y
func Maxu32(x, y uint32) uint32 {
	if x > y {
		return x
	}
	return y
}

// Maxi64 - Return the larger int64 of x or y
func Maxi64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

// Maxu64 - Return the larger uint64 of x or y
func Maxu64(x, y uint64) uint64 {
	if x > y {
		return x
	}
	return y
}
