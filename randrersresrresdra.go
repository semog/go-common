// Generate random numbers with a very high random repeat period.
// Based on the article "Fast, High-Quality, Parallel Random Number Generators"
// by Mark A. Overton, published in Dr. Dobb's Journal on 24 May, 2011.
// http://www.drdobbs.com/tools/fast-high-quality-parallel-random-number/229625477

package gocommon

import (
	"sync"
)

var mutex = &sync.Mutex{}

// RandRersResrResdra - Use the RERS, RESR, and RESDRA generators
type RandRersResrResdra struct {
	xx uint64
	yy uint64
	zz uint64
}

// Next - Generate random numbers with a very high random repeat period.
// Based on the article "Fast, High-Quality, Parallel Random Number Generators"
// by Mark A. Overton, published in Dr. Dobb's Journal on 24 May, 2011.
// http://drdobbs.com/go-parallel/article/229625477
func (rnd RandRersResrResdra) Next() uint64 {
	mutex.Lock()
	//RERS,   period = 4758085248529 (prime)
	rnd.xx = Rotlu64(rnd.xx, 8) - Rotlu64(rnd.xx, 29)
	//RESR,   period = 3841428396121 (prime)
	rnd.yy = Rotlu64(rnd.yy, 21) - rnd.yy
	rnd.yy = Rotlu64(rnd.yy, 20)
	//RESDRA, period = 5345004409 (prime)
	rnd.zz = Rotlu64(rnd.zz, 42) - rnd.zz
	rnd.zz = rnd.zz + Rotlu64(rnd.zz, 14)
	retVal := (rnd.xx ^ rnd.yy ^ rnd.zz)
	mutex.Unlock()
	return retVal
}
