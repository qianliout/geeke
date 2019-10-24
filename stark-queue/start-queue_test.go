package starkqueue

import (
	"fmt"
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	s := ")"
	r := IsValidParentheses(s)
	fmt.Println(r)
}
