package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("load bg")
	bgImgs := loadImgs("./res/bg/")
	fmt.Println("main bg")
	mainImgs := loadImgs("./res/main/")
	fmt.Println("item bg")
	itemImgs := loadImgs("./res/item/")

	for i, bg := range bgImgs {
		for j, main := range mainImgs {
			for k, item := range itemImgs {
				createImage(bg, main, item, i+"_"+j+"_"+k)
			}
		}
	}
}

func createImage(bgImg, mainImg, itemImg image.Image, imgName string) {
	startPointLogo := image.Point{0, 0}
	logoRectangle := image.Rectangle{startPointLogo, startPointLogo.Add(mainImg.Bounds().Size())}
	originRectangle := image.Rectangle{image.Point{0, 0}, bgImg.Bounds().Size()}

	rgba := image.NewRGBA(originRectangle)
	draw.Draw(rgba, originRectangle, bgImg, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, logoRectangle, mainImg, image.Point{0, 0}, draw.Over)
	draw.Draw(rgba, logoRectangle, itemImg, image.Point{0, 0}, draw.Over)

	outputPath := "./output/" + imgName + ".png"
	out, err := os.Create(outputPath)
	if err != nil {
		fmt.Println(err)
	}

	png.Encode(out, rgba)
}

func loadImgs(dirPath string) map[string]image.Image {
	imgMap := map[string]image.Image{}
	files, _ := ioutil.ReadDir(dirPath)
	for _, f := range files {
		img, err := loadImage(dirPath + f.Name())
		if err != nil {
			continue
		}
		replaced1 := strings.Replace(f.Name(), ".png", "", 1)
		fmt.Println(replaced1)
		imgMap[replaced1] = img
	}
	return imgMap
}

func loadImage(imgPath string) (image.Image, error) {
	fmt.Println("- loadImage : " + imgPath)
	imgFile, err := os.Open(imgPath)
	if err != nil {
		fmt.Println(err)
	}

	loadImg, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	return loadImg, err
}
