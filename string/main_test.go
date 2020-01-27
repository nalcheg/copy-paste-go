package main

import "testing"

func BenchmarkAddCondition(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var where WhereCondition
		where = where.AddCondition("to_account IS NULL ")
		_ = where.String()
	}
}

func BenchmarkAddConditionPointer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var where WhereCondition
		where.AddConditionPointer("to_account IS NULL ")
		_ = where.String()
	}
}

func TestWhereCondition_AddCondition(t *testing.T) {
	var where WhereCondition
	where = where.AddCondition("to_account IS NULL ")
	s := where.String()
	if " WHERE to_account IS NULL " != s {
		t.Error()
	}

	where.AddConditionPointer("from_account IS NULL ")
	if " WHERE to_account IS NULL  AND from_account IS NULL " != where.String() {
		t.Error()
	}

	ss := where.AddCondition("from_account IS NULL ")
	s += ss.String()
	if " WHERE to_account IS NULL  AND from_account IS NULL " != where.String() {
		t.Error()
	}
}
