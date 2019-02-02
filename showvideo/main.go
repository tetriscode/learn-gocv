package main

import (
	"fmt"
	"log"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tshowvideo [path to video]")
		return
	}

	imgFile := os.Args[1]
	//Create a Window
	window := gocv.NewWindow("Show Video")

	//Load the file
	capture, err := gocv.OpenVideoCapture(imgFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer capture.Close()
	defer window.Close()

	//Create empty image to load video file into
	img := gocv.NewMat()

	for {
		capture.Read(&img)
		window.IMShow(img)
		c := window.WaitKey(1)
		//break if escape key is hit
		if c == 27 {
			break
		}
	}
}
