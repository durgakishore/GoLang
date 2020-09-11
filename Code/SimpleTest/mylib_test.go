package simpletest

import "testing"

func TestMessageWriter(t *testing.T) {

	if messageWriter("Hello", "Kishore") != "Hello Kishore" {
		t.Fail()
	}
}
