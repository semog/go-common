// Copyright 2015 APCS Hackware, Inc.
// Generate random numbers with a very high random repeat period.
// Based on the article "Fast, High-Quality, Parallel Random Number Generators"
// by Mark A. Overton, published in Dr. Dobb's Journal on 24 May, 2011.
// http://drdobbs.com/go-parallel/article/229625477

package mathutil

import (
	"sync"
)

type IRandom interface {
	Next() uint64
	Next_i16(minVal int16, maxVal int16) int16
	Next_u16(minVal uint16, maxVal uint16) uint16
	Next_i32(minVal int32, maxVal int32) int32
	Next_u32(minVal uint32, maxVal uint32) uint32
	Next_i64(minVal int64, maxVal int64) int64
	Next_u64(minVal uint64, maxVal uint64) uint64
}

func Rotl_i32(value int32, count uint32) int32 {
	return (value << count) | (value >> (32 - count))
}

func Rotl_u32(value uint32, count uint32) uint32 {
	return (value << count) | (value >> (32 - count))
}

func Rotl_i64(value int64, count uint32) int64 {
	return (value << count) | (value >> (64 - count))
}

func Rotl_u64(value uint64, count uint32) uint64 {
	return (value << count) | (value >> (64 - count))
}

func Rotr_i32(value int32, count uint32) int32 {
	return (value >> count) | (value << (32 - count))
}

func Rotr_u32(value uint32, count uint32) uint32 {
	return (value >> count) | (value << (32 - count))
}

func Rotr_i64(value int64, count uint32) int64 {
	return (value >> count) | (value << (64 - count))
}

func Rotr_u64(value uint64, count uint32) uint64 {
	return (value >> count) | (value << (64 - count))
}

func getRange_i16(minVal int16, maxVal int16) int16 {
	rng := (maxVal - minVal)
	if 0 == rng {
		return 1
	} else {
		return rng
	}
}

func getRange_u16(minVal uint16, maxVal uint16) uint16 {
	rng := (maxVal - minVal)
	if 0 == rng {
		return 1
	} else {
		return rng
	}
}

func getRange_i32(minVal int32, maxVal int32) int32 {
	rng := (maxVal - minVal)
	if 0 == rng {
		return 1
	} else {
		return rng
	}
}

func getRange_u32(minVal uint32, maxVal uint32) uint32 {
	rng := (maxVal - minVal)
	if 0 == rng {
		return 1
	} else {
		return rng
	}
}

func getRange_i64(minVal int64, maxVal int64) int64 {
	rng := (maxVal - minVal)
	if 0 == rng {
		return 1
	} else {
		return rng
	}
}

func getRange_u64(minVal uint64, maxVal uint64) uint64 {
	rng := (maxVal - minVal)
	if 0 == rng {
		return 1
	} else {
		return rng
	}
}

type Random struct {
}

type MyRandom Random

type RandRersResrResdra struct {
	xx uint64
	yy uint64
	zz uint64
}

func (this Random) Next() uint64 {
	// TODO: Implement
	return 1
}

func (this Random) Next_i16(minVal int16, maxVal int16) int16 {
	return (minVal + (int16(this.Next()) % getRange_i16(minVal, maxVal)))
}

func (this Random) Next_u16(minVal uint16, maxVal uint16) uint16 {
	return (minVal + (uint16(this.Next()) % getRange_u16(minVal, maxVal)))
}

func (this Random) Next_i32(minVal int32, maxVal int32) int32 {
	return (minVal + (int32(this.Next()) % getRange_i32(minVal, maxVal)))
}

func (this Random) Next_u32(minVal uint32, maxVal uint32) uint32 {
	return (minVal + (uint32(this.Next()) % getRange_u32(minVal, maxVal)))
}

func (this Random) Next_i64(minVal int64, maxVal int64) int64 {
	return (minVal + (int64(this.Next()) % getRange_i64(minVal, maxVal)))
}

func (this Random) Next_u64(minVal uint64, maxVal uint64) uint64 {
	return (minVal + (uint64(this.Next()) % getRange_u64(minVal, maxVal)))
}

var mutex = &sync.Mutex{}

/// <summary>
/// Generate random numbers with a very high random repeat period.
/// Based on the article "Fast, High-Quality, Parallel Random Number Generators"
/// by Mark A. Overton, published in Dr. Dobb's Journal on 24 May, 2011.
/// http://drdobbs.com/go-parallel/article/229625477
/// </summary>

func (this RandRersResrResdra) Next() uint64 {
	mutex.Lock()
	//RERS,   period = 4758085248529 (prime)
	this.xx = Rotl_u64(this.xx, 8) - Rotl_u64(this.xx, 29)
	//RESR,   period = 3841428396121 (prime)
	this.yy = Rotl_u64(this.yy, 21) - this.yy
	this.yy = Rotl_u64(this.yy, 20)
	//RESDRA, period = 5345004409 (prime)
	this.zz = Rotl_u64(this.zz, 42) - this.zz
	this.zz = this.zz + Rotl_u64(this.zz, 14)
	retVal := (this.xx ^ this.yy ^ this.zz)
	mutex.Unlock()
	return retVal
}
