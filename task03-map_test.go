package homework

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortMapValues(t *testing.T) {
	m := map[int]string{
		222: "low",
		333: "world",
		111: "hell",
	}

	expected := []string{
		"hell",
		"low",
		"world",
	}

	actual := sortMapValues(m)

	if !reflect.DeepEqual(expected, actual) {
		fmt.Println(expected)
		fmt.Println(actual)
		t.Errorf("error, man")
	}
}
