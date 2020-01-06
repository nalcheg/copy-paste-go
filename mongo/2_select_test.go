package main

import "testing"

func TestMongo_SelectExampleTwoStraight(t *testing.T) {
	m := Mongo{}
	if err := m.Connect("mongodb://localhost:27017"); err != nil {
		t.Error(err)
	}

	pairs, err := m.SelectExampleOne("rur")
	if err != nil {
		t.Error(err)
	}

	p, err := m.SelectExampleTwoStraight(pairs)
	if err != nil {
		t.Error(err)
	}
	if len(p) != 94 {
		t.Log(len(p))
		t.Error()
	}
}

func BenchmarkMongo_SelectExampleTwoStraight(b *testing.B) {
	m := Mongo{}
	if err := m.Connect("mongodb://localhost:27017"); err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		pairs, err := m.SelectExampleOne("rur")
		if err != nil {
			b.Error(err)
		}

		p, err := m.SelectExampleTwoStraight(pairs)
		if err != nil {
			b.Error(err)
		}
		if len(p) != 94 {
			b.Log(len(p))
			b.Error()
		}
	}
}

func BenchmarkMongo_SelectExampleThreeConcurent(b *testing.B) {
	m := Mongo{}
	if err := m.Connect("mongodb://localhost:27017"); err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		pairs, err := m.SelectExampleOne("rur")
		if err != nil {
			b.Error(err)
		}

		p, err := m.SelectExampleThreeConcurent(pairs)
		if err != nil {
			b.Error(err)
		}
		if len(p) != 94 {
			b.Log(len(p))
			b.Error()
		}
	}
}
