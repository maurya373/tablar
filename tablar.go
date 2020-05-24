package main

import "fmt"

func main() {

	notes := [12]string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	key := "G"
	mode := "Phrygian"

	fmt.Println(getScale(notes, key, mode))
}

// Get the list of notes, in order, that form the scale for the key and the mode
func getScale(notes [12]string, key string, mode string) []string {

	var wrappedNotesForKey = wrapNotesForScale(notes, key)
	var modeFormula = getModeFormula(mode)

	var scale []string
	for _, i := range modeFormula {
		scale = append(scale, wrappedNotesForKey[i])
	}
	return scale
}

// Return the list of all notes rotated to start with the key
func wrapNotesForScale(notes [12]string, key string) []string {

	var noteList []string
	index := getIndexForElement(notes, key)

	for _, s := range notes[index:] {
		noteList = append(noteList, s)
	}
	for _, s := range notes[0: index] {
		noteList = append(noteList, s)
	}
	return noteList
}

// Get the index where the key is located in the notes list
func getIndexForElement(notes [12]string, key string) int {

	for i := range notes {
		if notes[i] == key {
			return i
		}
	}
	return -1
}

// Get the formula for a given mode, in order to create the full scale from a root note (key)
func getModeFormula(mode string) []int {

	switch mode {
	case "Major", "Ionian": 	return []int{0, 2, 4, 5, 7, 9, 11}
	case "Dorain": 			return []int{0, 2, 3, 5, 7, 9, 10}
	case "Phrygian":		return []int{0, 1, 3, 5, 7, 8, 10}
	case "Lydian": 			return []int{0, 2, 4, 6, 7, 9, 11}
	case "Mixolydian": 		return []int{0, 2, 4, 5, 7, 9, 10}
	case "Minor", "Aeolian": 	return []int{0, 2, 3, 5, 7, 8, 10}
	case "Locrian": 		return []int{0, 1, 3, 5, 6, 8, 10}
	default: 			return []int{}
	}
}
