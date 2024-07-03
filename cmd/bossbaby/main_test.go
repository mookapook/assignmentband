package main

import (
	"testing"
)

func TestBossBabyRevenge_boss_expected_good_boy(t *testing.T) {
	shot := "sSRSSRRRX"
	want := "Good boy"
	if got := checkBehavior(shot); got != want {
		t.Errorf("checkBehavior() = %v, want %v", got, want)
	}
}

func TestBossBabyRevenge_boss_expected_many(t *testing.T) {
	tests := map[string]string{
		"SRSSRRR":    "Good boy", // กรณีทั่วไปที่ Boss Baby ได้แก้แค้นครบทุกการยิง
		"RSSRR":      "Bad boy",  // กรณีที่ Boss Baby ยิงกลับก่อนจะถูกยิง
		"SSSRRRRS":   "Bad boy",  // กรณีที่การยิงสุดท้ายไม่ได้ถูกแก้แค้น
		"SRRSSR":     "Bad boy",  // กรณีที่มีการยิงที่ไม่ถูกแก้แค้น
		"SSRSRRR":    "Good boy", // กรณีที่ Boss Baby ได้แก้แค้นครบทุกการยิง
		"RRRR":       "Bad boy",  // กรณีที่มีแต่การยิงกลับโดยไม่มีการยิงใส่ก่อน
		"SSSS":       "Bad boy",  // กรณีที่มีแต่การยิงใส่โดยไม่มีการยิงกลับ
		"SRRS":       "Bad boy",  // กรณีที่มีการยิงแต่ไม่ได้แก้แค้นครบ
		"SRRSRR":     "Good boy", // กรณีที่ Boss Baby ได้แก้แค้นครบทุกการยิง
		"SR":         "Good boy", // กรณีที่มีการยิงและการแก้แค้นเพียงครั้งเดียว
		"R":          "Bad boy",  // กรณีที่มีการยิงกลับโดยไม่มีการยิงใส่ก่อน
		"S":          "Bad boy",  // กรณีที่มีการยิงใส่โดยไม่มีการยิงกลับ
		"SRRSRRRSS":  "Bad boy",  // กรณีที่มีการยิงสุดท้ายไม่ได้ถูกแก้แค้น
		"SSSSRRRRRR": "Good boy", // กรณีที่มีการยิงใส่หลายครั้งและการแก้แค้นครบทุกการยิง
		"":           "",
	}
	for s, index := range tests {
		if got := checkBehavior(s); got != index {
			t.Errorf("checkBehavior() = %v, want %v", s, index)
		}
	}

}
