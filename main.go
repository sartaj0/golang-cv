package main

import (
	"fmt"

	"gocv/cvt_color"

	"gocv/io"

	// "gocv/transform/resize"
	// "gocv/transform/rotate"

	"gocv/image_processing/blur"
	"gocv/image_processing/edge"
	"gocv/image_processing/morph"

	// "gocv/image_processing/thresh"

	"time"
)



func main(){
	// img_path := "test_folder/images/image1.png"
	img_path := "test_folder/images/558918.jpg"
	// img_path := "test_folder/images/download5.png"
	// img_path := "test_folder/images/kitten.png"

	

	img, err := io.ReadImage(img_path)
	if err != nil{
		fmt.Println("Err", err)
	}
	start := time.Now()

	// img = resize.Resize(img, 0, 800)
	// img = rotate.RotateImageDegree(img, 360)
	img, _ = blur.GaussianBlur(img, 7, 5)
	// gray, _ = blur.AverageBlur(gray, 7)
	gray := cvt_color.RGBToGray(img)
	// gray, _ = blur.MedianBlur(gray, 7)
	// img, _ = blur.GaussianBlur(gray, 7, 5)

	
	// thresh := thresh.Thresholding(gray, 225, 255, thresh.THRESH_BINARY_INV)
	// edge_img := edge.Sobel(gray, 0, 1)
	edge_img := edge.Canny(gray, 10, 30)
	// edge_img, _ := edge.Laplacian(gray)
	erosion_img := morph.Erosion(edge_img)
	erosion_img = morph.Erosion(erosion_img)
	dilated_img := morph.Dilation(edge_img)

	timeElapsed := time.Since(start)
	fmt.Println("This function took", timeElapsed, "time")


	// out.ImWrite("test_folder/output/thresh.png", thresh)
	io.ImWrite("test_folder/output/erosion_img.png", erosion_img)
	io.ImWrite("test_folder/output/dilated_img.png", dilated_img)
	io.ImWrite("test_folder/output/edge.png", edge_img)
	io.ImWrite("test_folder/output/gray.png", gray)
	io.ImWrite("test_folder/output/rgb.png", img)


}