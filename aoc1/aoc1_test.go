package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	vals := []string{"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet"}
	assert.Equal(t, 142, Calibrate(vals))
}

func TestUpdatedCalibration(t *testing.T) {
	vals := []string{"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen"}
	assert.Equal(t, 281, Calibrate(vals))
}

func TestRevised(t *testing.T) {
	vals := []string{"1abcone"}
	assert.Equal(t, 11, Calibrate(vals))
}

func TestLine19(t *testing.T) {
	vals := []string{"ntvhxqzsixxcrfpgstwo915onevxz"}
	assert.Equal(t, 61, Calibrate(vals))
}

func TestLine80Digits(t *testing.T) {
	s := "kbkclv3onelmf4ntxhxbrppsixsix"
	assert.Equal(t, []int{3, 1, 4, 6, 6}, digits(s))
}

func TestTricky(t *testing.T) {
	s := "8oneight9"
	assert.Equal(t, []int{8, 1, 9}, digits(s))
	assert.Equal(t, []int{2, 2, 5}, digits("225"))                                          //928
	assert.Equal(t, []int{6, 1, 1, 2}, digits("sixone12"))                                  // 927
	assert.Equal(t, []int{7, 7, 4}, digits("cvhlpzsbmknkqpgsevenlkzvm7hnznjsbszgvxrmdnn4")) //858
}

func TestLine5Digits(t *testing.T) {
	s := "lkdbjd5"
	assert.Equal(t, []int{5}, digits(s))
	assert.Equal(t, 55, Calibrate([]string{s}))
}

func TestLine61Digits(t *testing.T) {
	s := "dvxjblhdjqttxdfourhhrgdpmvvone83"
	assert.Equal(t, []int{4, 1, 8, 3}, digits(s))
}
