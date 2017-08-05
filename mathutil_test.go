package cmn

import "testing"

func TestRotl_i32(t *testing.T) {
	rot := Rotl_i32(1, 1)
	if rot != 2 {
		t.Fail()
	}
}

func TestRotl_u32(t *testing.T) {
	rot := Rotl_u32(1, 1)
	if rot != 2 {
		t.Fail()
	}
}

func TestRotl_i64(t *testing.T) {
	rot := Rotl_i64(1, 1)
	if rot != 2 {
		t.Fail()
	}
}

func TestRotl_u64(t *testing.T) {
	rot := Rotl_u64(1, 1)
	if rot != 2 {
		t.Fail()
	}
}

func TestRotr_i32(t *testing.T) {
	rot := Rotr_i32(2, 1)
	if rot != 1 {
		t.Fail()
	}
}

func TestRotr_u32(t *testing.T) {
	rot := Rotr_u32(2, 1)
	if rot != 1 {
		t.Fail()
	}
}

func TestRotr_i64(t *testing.T) {
	rot := Rotr_i64(2, 1)
	if rot != 1 {
		t.Fail()
	}
}

func TestRotr_64(t *testing.T) {
	rot := Rotr_u64(2, 1)
	if rot != 1 {
		t.Fail()
	}
}
