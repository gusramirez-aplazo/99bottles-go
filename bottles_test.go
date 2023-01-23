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

func TestVerse2(t *testing.T) {
	const verseNum uint = 2

	expected := "2 bottles of beer on the wall, " +
		"2 bottles of beer.\n" +
		"Take one down and pass it around, " +
		"1 bottle of beer on the wall.\n"

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

func TestVerse1(t *testing.T) {
	const verseNum uint = 1

	expected := "1 bottle of beer on the wall, " +
		"1 bottle of beer.\n" +
		"Take it down and pass it around, " +
		"no more bottles of beer on the wall.\n"

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

func TestVerse0(t *testing.T) {
	const verseNum uint = 0

	expected := "No more bottles of beer on the wall, " +
		"no more bottles of beer.\n" +
		"Go to the store and buy some more, " +
		"99 bottles of beer on the wall.\n"

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
