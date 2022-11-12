package iteration

import "testing"

// func TestRepeat(t *testing.T) {
func BenchmarkRepeat(t *testing.B) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
