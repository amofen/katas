package main

import (
	"testing"
)

func TestBase64Encode(t *testing.T) {
	encodedString, err := Base64Encode("This Is My Clear String")
	expected := "XYZNEWENCODEDSTHING"
	if expected != encodedString || err != nil {
		t.Errorf("Am learning from my failures, don't fix me, i'll fix myself.Or ... whatever !\nExpected : %v \nActual: %v", expected, encodedString)
	}
}
