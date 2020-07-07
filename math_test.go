package math

import (
	"strconv"
	"testing"
)

func TestCeilAfterDiv(t *testing.T) {
	val := CeilAfterDiv(10, 5)
	if val != 2 {
		t.Error("Failing for 10 / 5 = " + strconv.Itoa(val))
	}
	val = CeilAfterDiv(10, 3)
	if val != 4 {
		t.Error("Failing for 10 / 3 = " + strconv.Itoa(val))
	}
}
