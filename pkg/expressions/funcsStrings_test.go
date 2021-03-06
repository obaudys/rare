package expressions

import "testing"

func TestSubstring(t *testing.T) {
	testExpression(t,
		mockContext("abcd"),
		"{substr {0} 0 2} {substr {0} 0 10} {substr {0} 3 2} {substr {0} 3 1} {substr 0}",
		"ab abcd d d <ARGN>")
}

func TestSelect(t *testing.T) {
	testExpression(t,
		mockContext("ab c d", "ab\tq"),
		"{select {0} 0} {select {0} 1} {select {0} 2} {select {0} 3} {select 0} {select {1} 1}",
		"ab c d  <ARGN> q")
	testExpression(t, mockContext(), `{select "ab cd ef" 1}`, "cd")
}
