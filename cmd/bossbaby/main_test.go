package main

import (
	"testing"
)

func TestBossBabyCheck_expected_manygood_boy(t *testing.T) {
	shot := "sSRSSRRRX"
	want := "Good boy"
	if got := checkBehavior(shot); got != want {
		t.Errorf("checkBehavior() = %v, want %v", got, want)
	}
}

func TestBossBabyCheck_expected_many(t *testing.T) {
	tests := map[string]string{
		"SRSSRRR":    "Good boy", // General case where Boss Baby has revenged every shot
		"RSSRR":      "Bad boy",  // Case where Boss Baby shoots back before being shot at
		"SSSRRRRS":   "Bad boy",  // Case where the last shot is not revenged
		"SRRSSR":     "Bad boy",  // Case where some shots are not revenged
		"SSRSRRR":    "Good boy", // Case where Boss Baby has revenged every shot
		"RRRR":       "Bad boy",  // Case with only revenge shots without any prior shots
		"SSSS":       "Bad boy",  // Case with only shots without any revenge
		"SRRS":       "Bad boy",  // Case where not all shots are revenged
		"SRRSRR":     "Good boy", // Case where Boss Baby has revenged every shot
		"SR":         "Good boy", // Case with only one shot and one revenge
		"R":          "Bad boy",  // Case with a revenge shot without any prior shots
		"S":          "Bad boy",  // Case with a shot without any revenge
		"SRRSRRRSS":  "Bad boy",  // Case where the last shot is not revenged
		"SSSSRRRRRR": "Good boy", // Case with multiple shots and sufficient revenge
		"XEEESS":     "Bad boy",
		"XEEE":       "Good boy",
	}
	for s, index := range tests {
		if got := checkBehavior(s); got != index {
			t.Errorf("checkBehavior() = %v, want %v", s, index)
		}
	}

}
