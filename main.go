package main

import (
	"fmt"
	"image/png"
	"os"
	"qrudicon/lib"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	img := lib.NewSimpleQrudicon(`leck mich am arsch1 ich hasse dich ... asciiiiii-art :P`, 1024)
	fmt.Printf("%d\n", (time.Now().UnixNano()-t)/1000000)

	out, err := os.Create("/tmp/qrudicon.png")
	if err != nil {
		panic(err)
	}

	err = png.Encode(out, img)
}
