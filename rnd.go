// Generate random numbers with a very high random repeat period.
// Based on the article "Fast, High-Quality, Parallel Random Number Generators"
// by Mark A. Overton, published in Dr. Dobb's Journal on 24 May, 2011.
// http://www.drdobbs.com/tools/fast-high-quality-parallel-random-number/229625477

package gocommon

import "math"

// IRandomGenerator - random number generator interface
type IRandomGenerator interface {
	Next() uint64
}

// Random - Generate random numbers using the given random number generator.
type Random struct {
	rg IRandomGenerator
}

// MakeRandom - Instantiate a random number producer, given the random number generator.
func MakeRandom(rg IRandomGenerator) Random {
	return Random{rg}
}

// Next - Get the next random number
func (rnd Random) Next() uint64 {
	return rnd.rg.Next()
}

// Nexti16 - Get the next random int16 number
func (rnd Random) Nexti16() int16 {
	return int16(rnd.Next())
}

// Nexti16n - Get the next random int16 number within the specified range.
func (rnd Random) Nexti16n(minVal int16, maxVal int16) int16 {
	return minVal + (rnd.Nexti16() % getRangei16(minVal, maxVal))
}

// Nextu16 - Get the next random uint16 number
func (rnd Random) Nextu16() uint16 {
	return uint16(rnd.Next())
}

// Nextu16n - Get the next random uint16 number within the specified range
func (rnd Random) Nextu16n(minVal uint16, maxVal uint16) uint16 {
	return (minVal + (rnd.Nextu16() % getRangeu16(minVal, maxVal)))
}

// Nexti32 - Get the next random int32 number
func (rnd Random) Nexti32() int32 {
	return int32(rnd.Next())
}

// Nexti32n - Get the next random int32 number within the specified range
func (rnd Random) Nexti32n(minVal int32, maxVal int32) int32 {
	return minVal + (rnd.Nexti32() % getRangei32(minVal, maxVal))
}

// Nextu32 - Get the next random unt16 number
func (rnd Random) Nextu32() uint32 {
	return uint32(rnd.Next())
}

// Nextu32n - Get the next random unt16 number within the specified range
func (rnd Random) Nextu32n(minVal uint32, maxVal uint32) uint32 {
	return minVal + (rnd.Nextu32() % getRangeu32(minVal, maxVal))
}

// Nexti64 - Get the next random int64 number
func (rnd Random) Nexti64() int64 {
	return int64(rnd.Next())
}

// Nexti64n - Get the next random int64 number within the specified range
func (rnd Random) Nexti64n(minVal int64, maxVal int64) int64 {
	return minVal + (rnd.Nexti64() % getRangei64(minVal, maxVal))
}

// Nextu64 - Get the next random uint64 number
func (rnd Random) Nextu64() uint64 {
	return uint64(rnd.Next())
}

// Nextu64n - Get the next random uint64 number within the specified range
func (rnd Random) Nextu64n(minVal uint64, maxVal uint64) uint64 {
	return minVal + (rnd.Nextu64() % getRangeu64(minVal, maxVal))
}

// Nextf32 - Get the next random float32 number
func (rnd Random) Nextf32() float32 {
	return float32(Makef32(int32(rnd.Next()), int32(rnd.Next())))
}

// Nextf32n - Get the next random float32 number within the specified range
func (rnd Random) Nextf32n(minVal float32, maxVal float32) float32 {
	return float32(minVal + float32((math.Mod(float64(rnd.Nextf32()), float64(getRangef32(minVal, maxVal))))))
}

// Nextf64 - Get the next random float64 number
func (rnd Random) Nextf64() float64 {
	return float64(Makef64(int64(rnd.Next()), int64(rnd.Next())))
}

// Nextf64n - Get the next random float64 number within the specified range
func (rnd Random) Nextf64n(minVal float64, maxVal float64) float64 {
	return minVal + (math.Mod(rnd.Nextf64(), getRangef64(minVal, maxVal)))
}
