package homework

import (
	"fmt"
	"testing"
)

func TestAverage(t *testing.T) {
	a := [15]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 1395}

	var expected float32 = 100

	actual := average(a)

	if expected != actual {
		fmt.Println("expected:", expected)
		fmt.Println("actual:", actual)
		t.Errorf("error, man")
	}
}
