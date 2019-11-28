package payswitch

import "testing"

func TestCreate12DigitTransactionID(t *testing.T) {
	t.Log("TESTING VALID CASE")
	testValid := 3

	s, err := ConvT12DigitTransactionID(testValid)
	if err != nil {
		t.Fatal(s)
	}

	if len(s) != 12 {
		t.Fatal(s)
	}

	t.Log("TESTING INVALID CASE")
	testInvalid := 19900088800234
	s, err = ConvT12DigitTransactionID(testInvalid)
	if err == nil {
		t.Fatal(s)
	}
	t.Log(err)
}

func TestCreate12DigitAmount(t *testing.T) {
	t.Log("TESTING VALID CASE")
	testValid := 3.25

	s, err := ConvT12DigitAmount(testValid)

	if err != nil {
		t.Fatal(s)
	}

	if len(s) != 12 {
		t.Fatal(s)
	}

	t.Log("TESTING INVALID CASE")
	testInvalid := 199000888000222000555.214131
	s, err = ConvT12DigitAmount(testInvalid)
	if err == nil {
		t.Fatal(s)
	}
	t.Log(err)

}
