package main

import (
	"fmt"
	"image"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tshowimage [path to image]")
		return
	}

	// parse args
	src := os.Args[1]

	//Create a Window
	window := gocv.NewWindow("Show Blurred Image")
	//Read Image into Matrix
	mat := gocv.IMRead(src, gocv.IMReadColor)
	dst := gocv.NewMatWithSize(mat.Rows(), mat.Cols(), mat.Type())
	gocv.GaussianBlur(mat, &dst, image.Pt(75, 75), 5.0, 5.0, gocv.BorderDefault)

	//Show Image in Window
	window.IMShow(dst)

	//Hit Escape
	gocv.WaitKey(0)

	defer mat.Close()
	defer window.Close()
}
