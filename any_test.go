package any

import (
	"testing"
)

func TestAnyType(t *testing.T) {
	var a Type
	a = 1
	a = "aa"
	a = true
	a = struct{ Name string }{"foo"}
	a = map[string]string{"foo": "bar"}
	_ = a
}

func TestAnyMap(t *testing.T) {
	m := make(Map)
	m["a"] = 1
	m[2] = false
	m[true] = struct{}{}
	_ = m
}

func TestAnySlice(t *testing.T) {
	l := make(Slice, 0)
	l = append(l, 1, true, "foo", struct{}{})
	_ = l
}
