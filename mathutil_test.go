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
