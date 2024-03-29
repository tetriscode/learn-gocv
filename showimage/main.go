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

	//Create a Window
	window := gocv.NewWindow("Show Image")

	//Read Image into Matrix
	mat := gocv.IMRead(image, gocv.IMReadColor)

	//Show Image in Window
	window.IMShow(mat)

	//Hit Escape
	gocv.WaitKey(0)

	defer mat.Close()
	defer window.Close()
}
