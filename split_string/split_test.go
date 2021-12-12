package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected:%#v, got:%#v\n", want, got)
	}
}

func Test2Split(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c", "d"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected:%#v, got:%#v\n", want, got)
	}
}
