package _9bottles

import (
	"fmt"
)

func Verse(index uint) (string, error) {
	current := index
	previous := index - 1
	return fmt.Sprintf(
			"%v bottles of beer on the wall, %v bottles of beer.\nTake one down and pass it around, %v bottles of beer on the wall.\n", current, current, previous),
		nil
}
