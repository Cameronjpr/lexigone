package main

import "testing"

func TestFindableKeyAndGoodPath(t *testing.T) {
	found := search("./src", "key_2")

	if found != true {
		t.Errorf("search('./src', 'key_2') = %v; want %v", found, true)
	}
}

func TestUnfindableKeyAndGoodPath(t *testing.T) {
	found := search("./src", "key_1")

	if found != false {
		t.Errorf("search('./src', 'key_1') = %v; want %v", found, false)
	}
}
