package main

import (
	"headfirstgo/interfaces/gadget"
)

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
}

func playListRevised(device gadget.Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
}

func main() {
	// Before interface
	// Cant change to TapeRecorder without causing issues
	player := gadget.TapePlayer{}
	mixTape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixTape)

	var tPlayer gadget.Player = gadget.TapePlayer{}
	playListRevised(tPlayer, mixTape)

	var rPlayer gadget.Player = gadget.TapeRecorder{}
	playListRevised(rPlayer, mixTape)

	// Type Assertions
	// We want to access the Record method on TapeRecorder
	r2Player := rPlayer.(gadget.TapeRecorder)
	r2Player.Record()

}