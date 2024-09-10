package main

import (
	"fmt"

	"gocv/cvt_color"

	"gocv/out"
	"gocv/read"

	// "gocv/transform/resize"
	// "gocv/transform/rotate"

	"gocv/image_processing/blur"
	"gocv/image_processing/edge"

	// "gocv/image_processing/thresh"

	"time"
)



func main(){
	// img_path := "test_folder/images/image1.png"
	// img_path := "test_folder/images/558918.jpg"
	img_path := "test_folder/images/download5.png"
	// img_path := "test_folder/images/kitten.png"

	

	img, err := read.ReadImage(img_path)
	if err != nil{
		fmt.Println("Err", err)
	}
	start := time.Now()

	// img = resize.Resize(img, 0, 500)
	// img = rotate.RotateImageDegree(img, 360)
	img, _ = blur.GaussianBlur(img, 7, 5)
	// gray, _ = blur.AverageBlur(gray, 7)
	gray := cvt_color.RGBToGray(img)
	// gray, _ = blur.MedianBlur(gray, 7)
	// img, _ = blur.GaussianBlur(gray, 7, 5)

	
	// thresh := thresh.Thresholding(gray, 225, 255, thresh.THRESH_BINARY_INV)
	// edge_img := edge.Sobel(gray, 0, 1)
	edge_img := edge.Canny(gray, 60, 120)
	// edge_img, _ := edge.Laplacian(gray)

	timeElapsed := time.Since(start)
	fmt.Println("This function took", timeElapsed, "time")


	// out.ImWrite("test_folder/output/thresh.png", thresh)
	out.ImWrite("test_folder/output/edge.png", edge_img)
	out.ImWrite("test_folder/output/gray.png", gray)
	out.ImWrite("test_folder/output/rgb.png", img)


}