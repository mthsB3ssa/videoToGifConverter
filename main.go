package main

import (
    ffmpeg "github.com/u2takey/ffmpeg-go"
)

func VideoToGIF(inputPath, outputPath string) error {
    return ffmpeg.Input(inputPath).
        Output(outputPath, ffmpeg.KwArgs{
            "vf":   "fps=10,scale=320:-1:flags=lanczos",
            "loop": 0,
            "ss":   "0",
            "t":    "10",
        }).
        OverWriteOutput().
        Run()
}

func main() {
    err := VideoToGIF("input.mp4", "output.gif")
    if err != nil {
        panic(err)
    }
}
