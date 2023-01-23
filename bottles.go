package _9bottles

import (
	"fmt"
)

func Verse(index uint) (string, error) {
	current := index
	previous := index - 1

	switch index {

	case 0:
		return "No more bottles of beer on the wall, " +
				"no more bottles of beer.\n" +
				"Go to the store and buy some more, " +
				"99 bottles of beer on the wall.\n",
			nil
	case 1:
		return "1 bottle of beer on the wall, " +
				"1 bottle of beer.\n" +
				"Take it down and pass it around, " +
				"no more bottles of beer on the wall.\n",
			nil
	case 2:
		return "2 bottles of beer on the wall, " +
				"2 bottles of beer.\n" +
				"Take one down and pass it around, " +
				"1 bottle of beer on the wall.\n",
			nil
	default:
		return fmt.Sprintf(
				"%v bottles of beer on the wall, ", current) +
				fmt.Sprintf(
					"%v bottles of beer.\n", current) +
				"Take one down and pass it around, " +
				fmt.Sprintf(
					"%v bottles of beer on the wall.\n", previous),
			nil

	}

}
