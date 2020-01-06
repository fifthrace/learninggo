package main

import (
	"log"
	"os"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"time"
)

func main() {
	f, err := os.Open("Chad_Crouch_-_Shipping_Lanes.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)
}