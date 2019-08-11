package gocommon

import "testing"

func TestRotli32(t *testing.T) {
	rot := Rotli32(1, 1)
	if rot != 2 {
		t.Fail()
	}
}

func TestRotlu32(t *testing.T) {
	rot := Rotlu32(1, 1)
	if rot != 2 {
		t.Fail()
	}
}

func TestRotli64(t *testing.T) {
	rot := Rotli64(1, 1)
	if rot != 2 {
		t.Fail()
	}
}

func TestRotlu64(t *testing.T) {
	rot := Rotlu64(1, 1)
	if rot != 2 {
		t.Fail()
	}
}

func TestRotri32(t *testing.T) {
	rot := Rotri32(2, 1)
	if rot != 1 {
		t.Fail()
	}
}

func TestRotru32(t *testing.T) {
	rot := Rotru32(2, 1)
	if rot != 1 {
		t.Fail()
	}
}

func TestRotri64(t *testing.T) {
	rot := Rotri64(2, 1)
	if rot != 1 {
		t.Fail()
	}
}

func TestRotr64(t *testing.T) {
	rot := Rotru64(2, 1)
	if rot != 1 {
		t.Fail()
	}
}

func TestMini32(t *testing.T) {
	if Mini32(5, 7) != 5 {
		t.Fail()
	}
	if Mini32(7, 5) != 5 {
		t.Fail()
	}
	if Mini32(5, -7) != -7 {
		t.Fail()
	}
	if Mini32(-7, 5) != -7 {
		t.Fail()
	}
}

func TestMinu32(t *testing.T) {
	if Minu32(5, 7) != 5 {
		t.Fail()
	}
	if Minu32(7, 5) != 5 {
		t.Fail()
	}
}

func TestMini64(t *testing.T) {
	if Mini64(5, 7) != 5 {
		t.Fail()
	}
	if Mini64(7, 5) != 5 {
		t.Fail()
	}
	if Mini64(5, -7) != -7 {
		t.Fail()
	}
	if Mini64(-7, 5) != -7 {
		t.Fail()
	}
}

func TestMinu64(t *testing.T) {
	if Minu64(5, 7) != 5 {
		t.Fail()
	}
	if Minu64(7, 5) != 5 {
		t.Fail()
	}
}

func TestMaxi32(t *testing.T) {
	if Maxi32(5, 7) != 7 {
		t.Fail()
	}
	if Maxi32(7, 5) != 7 {
		t.Fail()
	}
	if Maxi32(5, -7) != 5 {
		t.Fail()
	}
	if Maxi32(-7, 5) != 5 {
		t.Fail()
	}
}

func TestMaxu32(t *testing.T) {
	if Maxu32(5, 7) != 7 {
		t.Fail()
	}
	if Maxu32(7, 5) != 7 {
		t.Fail()
	}
}

func TestMaxi64(t *testing.T) {
	if Maxi64(5, 7) != 7 {
		t.Fail()
	}
	if Maxi64(7, 5) != 7 {
		t.Fail()
	}
	if Maxi64(5, -7) != 5 {
		t.Fail()
	}
	if Maxi64(-7, 5) != 5 {
		t.Fail()
	}
}

func TestMaxu64(t *testing.T) {
	if Maxu64(5, 7) != 7 {
		t.Fail()
	}
	if Maxu64(7, 5) != 7 {
		t.Fail()
	}
}
