package homework

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	slice := []int64{4, 8, 2, 1, 4, 88}
	expected := []int64{88, 4, 1, 2, 8, 4}

	actual := reverse(slice)

	if !reflect.DeepEqual(expected, actual) {
		fmt.Println(expected)
		fmt.Println(actual)
		t.Errorf("error, man")
	}
}
