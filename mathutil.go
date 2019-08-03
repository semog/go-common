package gocommon

import (
	"math"
)

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

func getRangei16(minVal int16, maxVal int16) int16 {
	rng := (maxVal - minVal)
	if 0 == rng {
		return 1
	}
	return rng
}

func getRangeu16(minVal uint16, maxVal uint16) uint16 {
	rng := (maxVal - minVal)
	if 0 == rng {
		return 1
	}
	return rng
}

func getRangei32(minVal int32, maxVal int32) int32 {
	rng := (maxVal - minVal)
	if 0 == rng {
		return 1
	}
	return rng
}

func getRangeu32(minVal uint32, maxVal uint32) uint32 {
	rng := (maxVal - minVal)
	if 0 == rng {
		return 1
	}
	return rng
}

func getRangei64(minVal int64, maxVal int64) int64 {
	rng := (maxVal - minVal)
	if 0 == rng {
		return 1
	}
	return rng
}

func getRangeu64(minVal uint64, maxVal uint64) uint64 {
	rng := (maxVal - minVal)
	if 0 == rng {
		return 1
	}
	return rng
}

func getRangef32(minVal float32, maxVal float32) float32 {
	rng := (maxVal - minVal)
	if 0 == rng {
		return 1.0
	}
	return rng
}

func getRangef64(minVal float64, maxVal float64) float64 {
	rng := (maxVal - minVal)
	if 0 == rng {
		return 1.0
	}
	return rng
}

// Decompose a float32 value int two numbers. Given the following:
//    3.141592
// The return values will be:
//    base = 3
//    frac = 141592
func decompF32(val float32) (base int32, frac int32) {
	b, f := math.Modf(float64(val))
	return int32(b), int32(f)
}

// Decompose a float64 value int two numbers. Given the following:
//    3.141592
// The return values will be:
//    base = 3
//    frac = 141592
func decompF64(val float64) (base int64, frac int64) {
	b, f := math.Modf(val)
	return int64(b), int64(f)
}

// Make a float32 variable from two numbers. Given the following:
//    base = 3
//    frac = 141592
// The return float64 value will be:
//    3.14159
func makef32(base int32, frac int32) float32 {
	b := float32(base)
	f := float32(frac)
	for f > 1.0 {
		f /= 10.0
	}
	return b + f
}

// Make a float64 variable from two numbers. Given the following:
//    base = 3
//    frac = 141592
// The return float64 value will be:
//    3.14159
func makef64(base int64, frac int64) float64 {
	b := float64(base)
	f := float64(frac)
	for f > 1.0 {
		f /= 10.0
	}
	return b + f
}
