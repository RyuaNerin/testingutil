package testingutil

import "testing"

func Test_randomInt(t *testing.T) {
	for maxValue := 2; maxValue < 16*1024; maxValue++ {
		m := maxValue
		if m > 16*16 {
			m = 16 * 16
		}
		for fitValue := 1; fitValue < m; fitValue++ {
			for iter := 0; iter < 100; iter++ {
				v := randomInt(maxValue, fitValue)
				if v <= 0 || v >= maxValue {
					t.Errorf("randomInt(%d, %d) = %d", maxValue, fitValue, v)
				}
				if fitValue > 0 && v%fitValue != 0 {
					t.Errorf("randomInt(%d, %d) = %d", maxValue, fitValue, v)
				}
			}
		}
	}
}
