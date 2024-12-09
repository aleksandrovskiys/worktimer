package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func getTimeRepresentation(remaining time.Duration) string {
	remainingHours := math.Floor(remaining.Seconds() / 3600)
	remainingMinutes := math.Floor((remaining.Seconds() - remainingHours*3600) / 60)
	remainingSeconds := math.Ceil(remaining.Seconds() - (remaining.Seconds() - remainingHours*3600) + (remaining.Seconds() - remainingMinutes*60))
	t := time.Date(1, 1, 1, int(remainingHours), int(remainingMinutes), int(remainingSeconds), 0, time.Now().Location())
	s := t.Format(time.TimeOnly)
	return s
}

func playBellSound() {
	bellBytes, err := Asset("bell.wav")
	if err != nil {
		log.Fatalf("Can't open file: %v", err)
	}

	reader := bytes.NewReader(bellBytes)

	streamer, format, err := wav.Decode(reader)
	if err != nil {
		log.Fatalf("Can't decode file: %v", err)
	}

	defer streamer.Close()

	done := make(chan bool)

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
