package greet

import "testing"

func TestSayHi(t *testing.T) {
	logger = func(s string) {
		if s != "Hello, Daniel!" {
			t.Fail()
		}
	}

	SayHi("Daniel")
}
