package main

import (
	"fmt"
	"gocv/cvt_color"
	"gocv/image_processing/thresh"
	"gocv/out"
	"gocv/read"

	"gocv/image_processing/blur"
	"gocv/transform/resize"
	"gocv/transform/rotate"

	"time"
)



func main(){
	// img_path := "test_folder/images/image1.png"
	// img_path := "E:/Pictures/Images/558918.jpg"
	img_path := "test_folder/images/download5.png"
	// img_path := "E:/dataset/Face/Bolly/Faces/Akshay Kumar/Akshay Kumar_0.jpg"
	// img_path := "E:/Pictures/white.jpg"

	

	img, err := read.ReadImage(img_path)
	if err != nil{
		fmt.Println("Err", err)
	}
	start := time.Now()



	img = resize.Resize(img, 0, 1000)
	img = rotate.RotateImageDegree(img, 360)
	img, _ = blur.GaussianBlur(img, 9)
	gray := cvt_color.RGBToGray(img)
	
	thresh := thresh.Thresholding(gray, 120, 255, thresh.THRESH_BINARY)

	timeElapsed := time.Since(start)
	fmt.Println("This function took", timeElapsed, "time")


	out.ImWriteGray("./output.png", thresh)

}