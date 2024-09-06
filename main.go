package main

import (
	"fmt"

	"gocv/cvt_color"

	"gocv/out"
	"gocv/read"

	"gocv/transform/resize"
	"gocv/transform/rotate"

	"gocv/image_processing/blur"
	"gocv/image_processing/edge"
	"gocv/image_processing/thresh"

	"time"
)



func main(){
	// img_path := "test_folder/images/image1.png"
	// img_path := "test_folder/images/558918.jpg"
	img_path := "test_folder/images/download5.png"

	

	img, err := read.ReadImage(img_path)
	if err != nil{
		fmt.Println("Err", err)
	}
	start := time.Now()



	img = resize.Resize(img, 0, 400)
	img = rotate.RotateImageDegree(img, 360)
	img, _ = blur.GaussianBlur(img, 7)
	gray := cvt_color.RGBToGray(img)

	
	
	thresh := thresh.Thresholding(gray, 210, 255, thresh.THRESH_BINARY_INV)
	edge_img := edge.Sobel(thresh, 1, 1)

	timeElapsed := time.Since(start)
	fmt.Println("This function took", timeElapsed, "time")


	out.ImWriteGray("test_folder/output/thresh.png", thresh)
	out.ImWriteGray("test_folder/output/edge.png", edge_img)
	out.ImWriteGray("test_folder/output/gray.png", gray)
	out.ImWrite("test_folder/output/rgb.png", img)

}