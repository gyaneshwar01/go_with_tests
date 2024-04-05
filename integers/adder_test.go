package integers

import "testing"

func Test(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected %d got %d", expected, sum)
	}

}
