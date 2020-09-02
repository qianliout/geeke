package main

import (
	"testing"
)

func BenchmarkRemoveInvalidParentheses(t *testing.B) {
	for i := 0; i < t.N; i++ {
		removeInvalidParentheses("((()((s((((()")
	}
}
