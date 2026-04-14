package main

import (
	"bytes"
	"io"
	"os"
	//"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

var otoCtx *oto.Context
var readyChan chan struct{}
var contextCreated bool = false

var decodedMp3s [24][]byte

func createContext(op oto.NewContextOptions) (*oto.Context, chan struct{}) {
	otoCtx, readyChan, err := oto.NewContext(&op)
	if err != nil {
		panic("error in oto.NewContext" + err.Error())
	}

	return otoCtx, readyChan
}

func loadSounds(files [24]string) {
	for i, file := range files {
		// read mp3 into mem
		fileBytes, err := os.ReadFile(file)

		if err != nil {
			panic("error reading " + file + err.Error())
		}

		// decode
		fileBytesReader := bytes.NewReader(fileBytes)

		decoder, err := mp3.NewDecoder(fileBytesReader)
		if err != nil {
			panic("error decoding mp3" + err.Error())
		}

		pcm, err := io.ReadAll(decoder)

		decodedMp3s[i] = pcm
	}
}

func playSound(note int) {
	if !contextCreated {
		// Options
		op := &oto.NewContextOptions{}
		op.SampleRate = 44100
		op.ChannelCount = 2
		op.Format = oto.FormatSignedInt16LE
		op.BufferSize = 4096

		otoCtx, readyChan = createContext(*op)
		contextCreated = true

		// wait on the channel
		<-readyChan
	}

	reader := bytes.NewReader(decodedMp3s[note-1])
	player := otoCtx.NewPlayer(reader)

	player.Play()
}

// list
func playChord(notes []int) {

}
