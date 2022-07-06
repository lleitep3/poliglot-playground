package challenges

import "testing"

func TestReverse(t *testing.T) {
	reverseWord := Reverse("world")

	if reverseWord != "dlrow" {
		t.Errorf("Expected to be \"%s\" but received \"%s\"", "dlrow", reverseWord)
	}
}
