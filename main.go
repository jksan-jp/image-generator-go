package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	ogImgPath := "./res/bg/b_0.png"
	mainImagePath := "./res/main/m_0.png"
	itemImagePath := "./res/item/i_0.png"
	saveImagePath := "./output/result.png"

	bgImg, err := loadImage(ogImgPath)
	if err != nil {
		fmt.Println(err)
	}

	mainImg, err := loadImage(mainImagePath)
	if err != nil {
		fmt.Println(err)
	}

	itemImg, err := loadImage(itemImagePath)
	if err != nil {
		fmt.Println(err)
	}

	startPointLogo := image.Point{0, 0}
	logoRectangle := image.Rectangle{startPointLogo, startPointLogo.Add(mainImg.Bounds().Size())}
	originRectangle := image.Rectangle{image.Point{0, 0}, bgImg.Bounds().Size()}

	rgba := image.NewRGBA(originRectangle)
	draw.Draw(rgba, originRectangle, bgImg, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, logoRectangle, mainImg, image.Point{0, 0}, draw.Over)
	draw.Draw(rgba, logoRectangle, itemImg, image.Point{0, 0}, draw.Over)

	out, err := os.Create(saveImagePath)
	if err != nil {
			fmt.Println(err)
	}

	png.Encode(out, rgba)
}

func loadImage(imgPath string) (image.Image,error) {
	imgFile, err := os.Open(imgPath)
	if err != nil {
			fmt.Println(err)
	}

	loadImg, _, err := image.Decode(imgFile)
	if err != nil {
			fmt.Println(err)
	}

	return loadImg, err
}