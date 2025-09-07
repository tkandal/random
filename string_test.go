package random

import "testing"

func TestString16(t *testing.T) {
	l := 16
	s := String(l)
	if len(s) != l {
		t.Errorf("String(16) = %s, want length %d", s, l)
	}
}

func TestString15(t *testing.T) {
	l := 15
	s := String(l)
	if len(s) != (l + 1) {
		t.Errorf("String(16) = %s, want length 16", s)
	}
}

func TestStringMax(t *testing.T) {
	l := 257
	s := String(l)
	if len(s) != MaxStringLen {
		t.Errorf("String(257) = %s, want length %d", s, MaxStringLen)
	}
}

func TestSecureString16(t *testing.T) {
	l := 16
	s, err := SecureString(l)
	if err != nil {
		t.Errorf("SecureString(16) returned error: %v", err)
	}
	if len(s) != l {
		t.Errorf("SecureString(%d) = %s, want length %d", l, s, l)
	}
}
func TestSecureString15(t *testing.T) {
	l := 15
	s, err := SecureString(l)
	if err != nil {
		t.Errorf("SecureString(16) returned error: %v", err)
	}
	if len(s) != (l + 1) {
		t.Errorf("SecureString(15) = %s, want length 16", s)
	}
}

func TestSecureStringMax(t *testing.T) {
	l := 257
	s, err := SecureString(l)
	if err != nil {
		t.Errorf("SecureString(2048) returned error: %v", err)
	}
	if len(s) != MaxStringLen {
		t.Errorf("SecureString(2048) = %s, want length %d", s, MaxStringLen)
	}
}

func TestBytes16(t *testing.T) {
	l := 16
	b := Bytes(l)
	if len(b) != l {
		t.Errorf("Bytes(%d) = %v, want length %d", l, b, l)
	}
}

func TestBytesMax(t *testing.T) {
	l := 257
	b := Bytes(l)
	if len(b) != MaxStringLen {
		t.Errorf("Bytes(257) = %v, want length %d", b, MaxStringLen)
	}
}

func TestSecureBytes16(t *testing.T) {
	l := 16
	b, err := SecureBytes(l)
	if err != nil {
		t.Errorf("SecureBytes(16) returned error: %v", err)
	}
	if len(b) != l {
		t.Errorf("SecureBytes(%d) = %v, want length %d", l, b, l)
	}
}

func TestSecureBytesMax(t *testing.T) {
	l := 257
	b, err := SecureBytes(l)
	if err != nil {
		t.Errorf("SecureBytes(16) returned error: %v", err)
	}
	if len(b) != MaxStringLen {
		t.Errorf("SecureBytes(257) = %v, want length %d", b, MaxStringLen)
	}
}
