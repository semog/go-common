// Generate random numbers with a very high random repeat period.
// Based on the article "Fast, High-Quality, Parallel Random Number Generators"
// by Mark A. Overton, published in Dr. Dobb's Journal on 24 May, 2011.
// http://www.drdobbs.com/tools/fast-high-quality-parallel-random-number/229625477

package gocommon

import "math"

// IRandom - random number generator interface
type IRandom interface {
	Next() uint64
}

// Nexti16 - Get the next random int16 number
func Nexti16(rnd IRandom) int16 {
	return int16(rnd.Next())
}

// Nexti16n - Get the next random int16 number within the specified range.
func Nexti16n(rnd IRandom, minVal int16, maxVal int16) int16 {
	return minVal + (Nexti16(rnd) % getRangei16(minVal, maxVal))
}

// Nextu16 - Get the next random uint16 number
func Nextu16(rnd IRandom) uint16 {
	return uint16(rnd.Next())
}

// Nextu16n - Get the next random uint16 number within the specified range
func Nextu16n(rnd IRandom, minVal uint16, maxVal uint16) uint16 {
	return (minVal + (Nextu16(rnd) % getRangeu16(minVal, maxVal)))
}

// Nexti32 - Get the next random int32 number
func Nexti32(rnd IRandom) int32 {
	return int32(rnd.Next())
}

// Nexti32n - Get the next random int32 number within the specified range
func Nexti32n(rnd IRandom, minVal int32, maxVal int32) int32 {
	return minVal + (Nexti32(rnd) % getRangei32(minVal, maxVal))
}

// Nextu32 - Get the next random unt16 number
func Nextu32(rnd IRandom) uint32 {
	return uint32(rnd.Next())
}

// Nextu32n - Get the next random unt16 number within the specified range
func Nextu32n(rnd IRandom, minVal uint32, maxVal uint32) uint32 {
	return minVal + (Nextu32(rnd) % getRangeu32(minVal, maxVal))
}

// Nexti64 - Get the next random int64 number
func Nexti64(rnd IRandom) int64 {
	return int64(rnd.Next())
}

// Nexti64n - Get the next random int64 number within the specified range
func Nexti64n(rnd IRandom, minVal int64, maxVal int64) int64 {
	return minVal + (Nexti64(rnd) % getRangei64(minVal, maxVal))
}

// Nextu64 - Get the next random uint64 number
func Nextu64(rnd IRandom) uint64 {
	return uint64(rnd.Next())
}

// Nextu64n - Get the next random uint64 number within the specified range
func Nextu64n(rnd IRandom, minVal uint64, maxVal uint64) uint64 {
	return minVal + (Nextu64(rnd) % getRangeu64(minVal, maxVal))
}

// Nextf32 - Get the next random float32 number
func Nextf32(rnd IRandom) float32 {
	return float32(Makef32(int32(rnd.Next()), int32(rnd.Next())))
}

// Nextf32n - Get the next random float32 number within the specified range
func Nextf32n(rnd IRandom, minVal float32, maxVal float32) float32 {
	return float32(minVal + float32((math.Mod(float64(Nextf32(rnd)), float64(getRangef32(minVal, maxVal))))))
}

// Nextf64 - Get the next random float64 number
func Nextf64(rnd IRandom) float64 {
	return float64(Makef64(int64(rnd.Next()), int64(rnd.Next())))
}

// Nextf64n - Get the next random float64 number within the specified range
func Nextf64n(rnd IRandom, minVal float64, maxVal float64) float64 {
	return minVal + (math.Mod(Nextf64(rnd), getRangef64(minVal, maxVal)))
}
