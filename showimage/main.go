package main

import (
	"fmt"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tshowimage [path to image]")
		return
	}

	// parse args
	image := os.Args[1]

	window := gocv.NewWindow("Show Image")
	mat := gocv.IMRead(image, gocv.IMReadColor)
	window.IMShow(mat)
	gocv.WaitKey(0)
	defer mat.Close()
	defer window.Close()
}
