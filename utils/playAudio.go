package utils

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
	"log"
	"os"
	"time"
)

func PlayAudio(filePath, musicType string) {
	var (
		streamer beep.StreamSeekCloser
		format   beep.Format
		err      error
	)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	switch musicType {
	case "mp3":
		streamer, format, err = mp3.Decode(f)
	case "wav":
		streamer, format, err = wav.Decode(f)
	case "ogg":
		streamer, format, err = vorbis.Decode(f)
	case "flac":
		streamer, format, err = flac.Decode(f)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer func(streamer beep.StreamSeekCloser) {
		err := streamer.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(streamer)

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
