package rohit

import "testing"

func TestMama(t *testing.T) {
	expected := "నమస్కారం"
	actual := Mama()
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
