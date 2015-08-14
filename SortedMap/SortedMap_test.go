package SortedMap_test

import (
	"reflect"
	"testing"

	"bitbucket.org/prvn30/polstalker/SortedMap"
)

func Test_SortedMap(t *testing.T) {
	values := map[string]int{
		"arch":      784,
		"manjaro":   922,
		"freebsd":   400,
		"fedora":    1145,
		"dragonfly": 209}
	expected := []string{"fedora", "manjaro", "arch", "freebsd", "dragonfly"}

	actual := SortedMap.SortedKeys(values)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test failed, expected: '%v:%v:%v', got: '%v:%v:%v'", expected, expected)
	}
}
