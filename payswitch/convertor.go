package payswitch

import (
	"errors"
	"strconv"
)

func ConvT12DigitTransactionID(k int) (string, error) {
	n := strconv.Itoa(k)
	z := 12 - len(n)
	if z < 0 {
		return "", errors.New("int k provided out of bound: length of k < 12")
	}
	t := ""
	for i := 0; i < z; i++ {
		t += "0"
	}
	t += n
	return t, nil
}

func ConvT12DigitAmount(k float64) (string, error) {
	n := strconv.Itoa(int(k * 100))
	z := 12 - len(n)

	if z < 0 {
		return "", errors.New("float k provided out of bound: length of k*100 < 12")
	}

	t := ""
	for i := 0; i < z; i++ {
		t += "0"
	}
	t += n
	return t, nil
}
