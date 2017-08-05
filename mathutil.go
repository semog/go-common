package cmn

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
