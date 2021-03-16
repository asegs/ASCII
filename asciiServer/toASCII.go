package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func createPixels(filename string)Picture {
	image.RegisterFormat("jpeg","jpeg",jpeg.Decode,jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error: File could not be opened")
		fmt.Println(err.Error())
	}

	defer file.Close()

	pixels, err := getPixels(file)

	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		fmt.Println(err.Error())

	}
	return Picture{ImageData: pixels}
}

// Get the bi-dimensional pixel array
func getPixels(file io.Reader) ([][][4]int, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][][4]int
	for y := 0; y < height; y++ {
		var row [][4]int
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) [4]int {
	pixel := [4]int{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
	return pixel
}


type Picture struct {
	ImageData [][][4] int
	heightIdx int
	widthIdx int
}


func handler(filename string, inverse bool,extension string)string{
	start := time.Now()
	p := createPixels("/home/arctic/Documents/Programming/Git/ASCII/asciiServer/images/"+filename+extension)
	chunkHeight := int(len(p.ImageData)/(getPropLength(p, 600)/2 +1))
	chunkWidth := int(len(p.ImageData[0])/600)
	chars := make([][]rune,roundingHandler(len(p.ImageData),chunkHeight))
	for x:=0;x<len(chars);x++{
		chars[x] = make([]rune,roundingHandler(len(p.ImageData[0]),chunkWidth))
	}
	chars = vaporize(p,600,int(getPropLength(p,600)/2),chars)
	var sb strings.Builder
	for i:=0;i<len(chars);i++{
		for b:=0;b<len(chars[0]);b++{
			sb.WriteRune(chars[i][b])
		}
		sb.WriteRune('\n')
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
	Write("/home/arctic/Documents/Programming/Git/ASCII/asciiServer/textfiles/knight.txt",sb.String())
	return sb.String()
}

func roundingHandler(dim int,chunkDim int)int{
	var iters int
	if dim % chunkDim ==0{
		iters = dim/chunkDim
	}else{
		fDim := float64(dim)
		fChunkDim := float64(chunkDim)
		iters = (int)(fDim/fChunkDim+1)
	}
	return iters
}


func vaporize(picture Picture,width int,height int,chars [][]rune)[][]rune{
	complete := make(chan bool,1)
	chunkHeight := int(len(picture.ImageData)/height)
	chunkWidth := int(len(picture.ImageData[0])/width)
	heightIters := roundingHandler(len(picture.ImageData),chunkHeight)
	widthIters := roundingHandler(len(picture.ImageData[0]),chunkWidth)
	size := heightIters*widthIters
	toFill := make(chan bool,size)
	readyToConvert := make(chan Picture,size)
	heightIdx := 0
	widthIdx := 0
	for row := 0;row<len(picture.ImageData);row+=chunkHeight{
		widthIdx = 0
		for col := 0;col<len(picture.ImageData[0]);col+=chunkWidth{
			go createPhoto(chunkHeight,chunkWidth,row,&picture,col,heightIdx,widthIdx,&readyToConvert)
			go func() {
				picture := <-readyToConvert
				chars[picture.heightIdx][picture.widthIdx] = ascii(picture,false)
				complete<-true
			}()
			widthIdx++
		}
		heightIdx++
	}
	go func() {
		for true{
			if len(toFill)==cap(toFill){
				break
			}
		}
		complete <- true
	}()
	<-complete
	return chars
}

func createPhoto(chunkHeight int,chunkWidth int,row int,picture *Picture,col int,heightIdx int,widthIdx int,needParsing *chan Picture){
	//initialize new width x height chunk
	chunk := make([][][4]int,chunkHeight)
	for i := 0;i<chunkHeight;i++{
		chunk[i] = make([][4]int,chunkWidth)
	}

	for i:= row;i<row+chunkHeight&&i<len(picture.ImageData);i++{
		chunk[i-row] = picture.ImageData[i][col:col+chunkWidth]
	}
	*needParsing<-Picture{ImageData: chunk,heightIdx: heightIdx,widthIdx: widthIdx}

}

func intAbs(i int)int{
	if i<0{
		return i*-1
	}
	return i
}

func Write(filename string,body string) {

	f, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(body)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Printf("Wrote to file: %s\n",filename)
}

func ascii(picture Picture,inverse bool)rune{
	if len(picture.ImageData)==0{
		return ' '
	}
	returns := [12]rune{'M','N','m','/','d','y','s','+',':','-','`',' '}
	imageData := picture.ImageData
	boxCount := len(imageData)*len(imageData[0])
	totalColorNum := 0
	toSubtract := 11
	if inverse{
		toSubtract = 0
	}
	for row := 0;row<len(imageData);row++{
		for col := 0;col<len(imageData[0]);col++{
			for i:=0;i<3;i++{
				totalColorNum+=imageData[row][col][i]
			}
		}
	}
	avgDarkness := totalColorNum/(boxCount*3)
	if avgDarkness<=30{
		return returns[intAbs(toSubtract-11)]
	}else if avgDarkness<=45{
		return returns[intAbs(toSubtract-10)]
	}else if avgDarkness<=60{
		return returns[intAbs(toSubtract-9)]
	}else if avgDarkness<=80{
		return returns[intAbs(toSubtract-8)]
	}else if avgDarkness<=100{
		return returns[intAbs(toSubtract-7)]
	}else if avgDarkness<=120{
		return returns[intAbs(toSubtract-6)]
	}else if avgDarkness<=140{
		return returns[intAbs(toSubtract-5)]
	}else if avgDarkness<=160{
		return returns[intAbs(toSubtract-4)]
	}else if avgDarkness<=180{
		return returns[intAbs(toSubtract-3)]
	}else if avgDarkness<=200{
		return returns[intAbs(toSubtract-2)]
	}else if avgDarkness<=220{
		return returns[intAbs(toSubtract-1)]
	}else {
		return returns[toSubtract]
	}
}

func getPropLength(picture Picture,height int) int{
	return len(picture.ImageData[0])*height/len(picture.ImageData)
}