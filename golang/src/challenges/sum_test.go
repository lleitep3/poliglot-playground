package challenges

import (
	"testing"
)

func TestSumInt(t *testing.T) {
	total := Sum[int64](1, 2, 3)

	if total != 6 {
		t.Errorf("Expected Total to be (%d) but receive (%d)", 6, total)
	}
}

func TestSumFloat(t *testing.T) {
	total := Sum(11.23, 2.34, 3.22)

	if total != 16.79 {
		t.Errorf("Expected Total to be (%f) but receive (%f)", 16.79, total)
	}
}
