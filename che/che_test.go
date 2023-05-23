package che

import (
	"reflect"
	"testing"
)

func TestChe(t *testing.T) {
	res := Che()
	f := "Привет, мир!"

	if !reflect.DeepEqual(res, f) {
		t.Errorf("error")
		return
	}
}
