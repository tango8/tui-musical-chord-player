package main

import ()

type Mode int

const (
	Ionian  = iota // major
	Aeolian        // natural minor
)

var ionianDiatonic [7]string = [7]string{"I", "ii", "iii", "IV", "V", "vi", "vii°"}
var aeolianDiatonic [7]string = [7]string{"i", "ii°", "III", "iv", "v", "VI", "VII"}
