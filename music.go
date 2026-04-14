package main

import ()

type Mode int

const (
	Ionian  = iota // major
	Aeolian        // natural minor
)

type Chord int

const (
	Major = iota
	Minor
	Dominant
)

// just use + or - 1 for more obscure extensions
type Extension int

const (
	MinorSeventh = iota
	MajorSeventh
	Ninth
	Eleventh
)

//var map[Chord][]notes

var ionianDiatonic [7]string = [7]string{"I", "ii", "iii", "IV", "V", "vi", "vii°"}
var aeolianDiatonic [7]string = [7]string{"i", "ii°", "III", "iv", "v", "VI", "VII"}

var currentNote = "C"
var currentKeyType = "Major"

func getChord(name Chord, extensions []Extension) []int {
	return []int{}
}

func getCurrentKey() string {
	return currentNote + " " + currentKeyType
}
