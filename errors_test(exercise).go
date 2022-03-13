package timeparse

import "testing"

func TestParseTime (tyme *testing.T) {
	table := []struct {
		time string
		ok bool
	} {
		{"19:00:14", true},
		{"1:00:04", true},
		{"a word", false},
		{"9:-4:57", false},
		{"", false},
		{"19:00:", false},
		{"a1:b2:x3", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			tyme.Errorf("%v: %v, error should be nil", data.time, err)
		}
	}
}