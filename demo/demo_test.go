package demo

import "testing"

func TestSayHi(t *testing.T) {
	actualResult := SayHi()
	if actualResult != "Say hi from demo" {
		t.Error("Error")
	}
}
