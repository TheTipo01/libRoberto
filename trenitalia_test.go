package libroberto

import (
	"testing"
)

func TestSearchAndGetTrain(t *testing.T) {
	str := SearchAndGetTrain("9508")
	if str == "" {
		t.Fatal("Empty train response!")
	} else {
		t.Log(str)
	}
}
