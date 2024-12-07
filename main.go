package main

import (
	"fmt"
	"os/exec"
)

func main() {
	inputMP4 := "input.mp4"
	outputHLS := "output.m3u8"

	cmd := exec.Command("ffmpeg", "-i", inputMP4, "-codec: copy", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", outputHLS)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("HLS conversion complete.")
	}
}
