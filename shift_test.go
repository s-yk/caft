package main

import "testing"

func TestShift(t *testing.T) {
	a := shift("abcxyz", 1, false)
	e := "bcdyza"
	if e != a {
		t.Errorf("expected: %s, but: %s", e, a)
	}

	a = shift("abcxyz", 1, true)
	e = "zabwxy"
	if e != a {
		t.Errorf("expected: %s, but: %s", e, a)
	}

	a = shift("abcxyz", 2, false)
	e = "cdezab"
	if e != a {
		t.Errorf("expected: %s, but: %s", e, a)
	}

	a = shift("abcxyz", 2, true)
	e = "yzavwx"
	if e != a {
		t.Errorf("expected: %s, but: %s", e, a)
	}

	a = shift("ABCXYZ", 1, false)
	e = "BCDYZA"
	if e != a {
		t.Errorf("expected: %s, but: %s", e, a)
	}

	a = shift("ABCXYZ", 1, true)
	e = "ZABWXY"
	if e != a {
		t.Errorf("expected: %s, but: %s", e, a)
	}

	a = shift("ABCXYZ", 2, false)
	e = "CDEZAB"
	if e != a {
		t.Errorf("expected: %s, but: %s", e, a)
	}

	a = shift("ABCXYZ", 2, true)
	e = "YZAVWX"
	if e != a {
		t.Errorf("expected: %s, but: %s", e, a)
	}

	a = shift("012789", 1, false)
	e = "123890"
	if e != a {
		t.Errorf("expected: %s, but: %s", e, a)
	}

	a = shift("012789", 1, true)
	e = "901678"
	if e != a {
		t.Errorf("expected: %s, but: %s", e, a)
	}

	a = shift("012789", 2, false)
	e = "234901"
	if e != a {
		t.Errorf("expected: %s, but: %s", e, a)
	}

	a = shift("012789", 2, true)
	e = "890567"
	if e != a {
		t.Errorf("expected: %s, but: %s", e, a)
	}

	a = shift("あいうえお", 2, true)
	e = "あいうえお"
	if e != a {
		t.Errorf("expected: %s, but: %s", e, a)
	}

	a = shift("abc 123 あいうえお xyz", 1, false)
	e = "bcd 234 あいうえお yza"
	if e != a {
		t.Errorf("expected: %s, but: %s", e, a)
	}
}
