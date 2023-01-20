package _9bottles

import (
	"testing"
)

func TestFirstVerse(t *testing.T) {
	const verseNum uint = 99

	expected := "99 bottles of beer on the wall, " +
		"99 bottles of beer.\n" +
		"Take one down and pass it around, " +
		"98 bottles of beer on the wall.\n"

	verse, err := Verse(verseNum)

	if verse == "" || err != nil {
		t.Errorf(
			"Can't retrieve verse num: %v",
			verseNum)
	}

	if verse != expected {
		t.Errorf(
			"Incorrect output>> \nExpected:\n%v\nActual:\n%v\n",
			expected,
			verse)
	}
}

func TestAnotherVerse(t *testing.T) {
	const verseNum uint = 3

	expected := "3 bottles of beer on the wall, " +
		"3 bottles of beer.\n" +
		"Take one down and pass it around, " +
		"2 bottles of beer on the wall.\n"

	verse, err := Verse(verseNum)
	if verse == "" || err != nil {
		t.Errorf(
			"Can't retrieve verse num: %v",
			verseNum)
	}

	if verse != expected {
		t.Errorf(
			"Incorrect output>> \nExpected:\n%v\nActual:\n%v\n",
			expected,
			verse)
	}
}
