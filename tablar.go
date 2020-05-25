package main

import (
	"fmt"
	"strconv"
)
import "math/rand"
import "time"

func main() {

	notes := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

	// Customize params for variable results
	key := "G"
	mode := "Phrygian"
	notesToGenerate := 12

	var scale = getScale(notes, key, mode)
	fmt.Println(key,mode,":",scale)

	var fretBoard = createFretBoard(notes)
	var playableNotes = getPlayableNotes(fretBoard, scale)

	var riff = generateRiff(playableNotes, notesToGenerate)
	printRiff(riff)
}

// Generate a random riff with the number of playable notes defined
// TODO: Provide constraints on how far away the notes can be from one another
func generateRiff(playableNotes [][]int, count int) [][]string {

	var riff [][]string
	for i := 0; i < count; i++ {
		var stringNum = random(len(playableNotes)-1)
		var noteInScale = random(len(playableNotes[0])-1)
		var fret = playableNotes[stringNum][noteInScale]

		var note = []string{"-","-","-","-","-","-"}
		note[stringNum] = strconv.Itoa(fret)
		riff = append(riff, note)
	}
	return riff
}

// Print the generated tablature
func printRiff(tab [][]string) {

	for i := len(tab[0])-1; i >= 0; i-- {
		for j := 0; j < len(tab); j++ {
			if j > len(tab) || i > len(tab[0]) {
				continue
			}
			fmt.Print(tab[j][i],"-")
		}
		fmt.Println()
	}
}

// Generate a Standard E Tuning fretboard with notes for open strings up to fret 11
// TODO: Support drop tunings
func createFretBoard(notes []string) [][]string {

	var fretBoard [][]string
	fretBoard = append(fretBoard, wrapNotesForScale(notes, "E"))
	fretBoard = append(fretBoard, wrapNotesForScale(notes, "A"))
	fretBoard = append(fretBoard, wrapNotesForScale(notes, "D"))
	fretBoard = append(fretBoard, wrapNotesForScale(notes, "G"))
	fretBoard = append(fretBoard, wrapNotesForScale(notes, "B"))
	fretBoard = append(fretBoard, wrapNotesForScale(notes, "E"))
	return fretBoard
}

// Generate a list of frets that are valid to play given the scale
func getPlayableNotes(fretBoard [][]string, scale []string) [][]int {

	var playableNotes [][]int
	for i := 0; i < len(fretBoard); i++ {
		var stringNotes []int
		for j := 0; j < len(fretBoard[i]); j++ {
			// If note is in scale, add to playable notes on string
			if getIndexForElement(scale, fretBoard[i][j]) != -1 {
				stringNotes = append(stringNotes, j)
			}
		}
		playableNotes = append(playableNotes, stringNotes)
	}
	return playableNotes
}

// Get the list of notes, in order, that form the scale for the key and the mode
func getScale(notes []string, key string, mode string) []string {

	var wrappedNotesForKey = wrapNotesForScale(notes, key)
	var modeFormula = getModeFormula(mode)

	var scale []string
	for _, i := range modeFormula {
		scale = append(scale, wrappedNotesForKey[i])
	}
	return scale
}

// Return the list of all notes rotated to start with the key
func wrapNotesForScale(notes []string, key string) []string {

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
func getIndexForElement(notes []string, key string) int {

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
	case "Major", "Ionian": return []int{0, 2, 4, 5, 7, 9, 11}
	case "Dorain": return []int{0, 2, 3, 5, 7, 9, 10}
	case "Phrygian": return []int{0, 1, 3, 5, 7, 8, 10}
	case "Lydian": return []int{0, 2, 4, 6, 7, 9, 11}
	case "Mixolydian": return []int{0, 2, 4, 5, 7, 9, 10}
	case "Minor", "Aeolian": return []int{0, 2, 3, 5, 7, 8, 10}
	case "Locrian": return []int{0, 1, 3, 5, 6, 8, 10}
	default: return []int{}
	}
}

// Return a random number between 0 and the passed-in limit
func random(limit int) int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	return rand.Intn(limit - min + 1) + min
}
