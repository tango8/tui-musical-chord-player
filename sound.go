package main

import (
	"bytes"
	//"io"
	"os"
	//"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

var otoCtx *oto.Context
var readyChan chan struct{}
var contextCreated bool = false

var player *oto.Player

func createContext(op oto.NewContextOptions) (*oto.Context, chan struct{}) {
	otoCtx, readyChan, err := oto.NewContext(&op)
	if err != nil {
		panic("error in oto.NewContext" + err.Error())
	}

	return otoCtx, readyChan
}

func playSound(filename string) {
	// read mp3 into mem
	fileBytes, err := os.ReadFile(filename)

	if err != nil {
		panic("error reading " + filename + err.Error())
	}

	// decode into reader obj
	fileBytesReader := bytes.NewReader(fileBytes)

	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic("error decoding mp3" + err.Error())
	}

	// Options
	op := &oto.NewContextOptions{}
	op.SampleRate = 44100
	op.ChannelCount = 2
	op.Format = oto.FormatSignedInt16LE
	op.BufferSize = 16384

	if !contextCreated {
		otoCtx, readyChan = createContext(*op)
		contextCreated = true
	}

	player = otoCtx.NewPlayer(decodedMp3) // paused by default

	player.Play() // Play() is async

	// wait on the channel
	<-readyChan

	/*for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}*/
	//
	// Now that the sound finished playing, we can restart from the beginning (or go to any location in the sound) using seek
	// newPos, err := player.(io.Seeker).Seek(0, io.SeekStart)
	// if err != nil{
	//     panic("player.Seek failed: " + err.Error())
	// }
	// println("Player is now at position:", newPos)
	// player.Play()
}
