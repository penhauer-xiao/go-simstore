package simstore

import (
	"math/rand"
	"testing"
)

func TestAdd3(t *testing.T) {
	s := New3(1000000)
	testAdd(t, s, 3)
}

func TestAdd3Small(t *testing.T) {
	s := New3Small(1000000)
	testAdd(t, s, 3)
}

func TestAdd6(t *testing.T) {
	s := New6(1000000)
	testAdd(t, s, 6)
}

func testAdd(t *testing.T, s Storage, d int) {

	for i := 0; i < 1000000; i++ {
		s.Add(uint64(rand.Int63()), uint64(i))
	}

	sig := uint64(0x001122334455667788)
	s.Add(sig, 0xdeadbeef)

	s.Finish()

	var fails int

	const queries = 1000000

	for j := 0; j < queries; j++ {

		q := sig

		// bits := rand.Intn(7)
		bits := d

		for i := 0; i < bits; i++ {
			q ^= 1 << uint(rand.Intn(64))
		}

		found := s.Find(q)
		if len(found) != 1 || found[0] != 0xdeadbeef {
			t.Errorf("sig = %016x (%064b)\n", sig, sig^q)
			fails++
		}
	}

	if fails != 0 {
		t.Logf("fails = %f", 100*float64(fails)/queries)
	}
}