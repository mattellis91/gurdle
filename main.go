package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)


type Game struct {
	runes []rune
}

const (
	title string = "Gurdle"
	width int = 435
	height int = 600
	rows int = 6
	cols int = 5
)

var (
	fontSize = 24
	mplusNormalFont font.Face
	bkg = color.White
	lightGray = color.RGBA{0xc2, 0xc5, 0xc6, 0xff}
	gray = color.RGBA{0x77, 0x7c, 0x7e, 0xff}
	yellow = color.RGBA{0xcd, 0xb3, 0x5d, 0xff}
	green = color.RGBA{0x60, 0xa6, 0x65, 0xff}
	fontColor = color.Black
	edge = false
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	grid [cols*rows]string
	dict []string
	check [cols*rows]int
	loc int = 0
	won = false
	answer string
)

func (g *Game) repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bkg)
	for w := 0; w < cols; w++ { // draw grid
		for h := 0; h < rows; h++ {
			rect := ebiten.NewImage(75, 75)
			rect.Fill(lightGray)
			fontColor = color.Black
			if check[w+(h*cols)] != 0 {
				if check[w+(h*cols)] == 1 { rect.Fill(green)}
				if check[w+(h*cols)] == 2 { rect.Fill(yellow)}
				if check[w+(h*cols)] == 3 { rect.Fill(gray)}
				fontColor = color.White
			}
			if w+cols*h == loc  && check[w+(h*cols)] == 0 { 
				rect.Fill(gray)
			}
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(w*85+10), float64(h*85+10))
			screen.DrawImage(rect, op)
			if check[w+(h*cols)] == 0 {
				rect2 := ebiten.NewImage(73, 73)
				rect2.Fill(color.White)
				op2 := &ebiten.DrawImageOptions{}
				op2.GeoM.Translate(float64(w*85+10)+1, float64(h*85+10)+1)
				screen.DrawImage(rect2, op2)
			}
			if grid[w+(h*cols)] != "" {
				msg := fmt.Sprintf("%s", strings.ToUpper(grid[w+(h*cols)]))
				text.Draw(screen, msg, mplusNormalFont, w*85+38, h*85+55, fontColor)
			}
		}
	}
}

func (g *Game)	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	game := &Game{}
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(fontSize),
		DPI:    72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gurdle")

	content, err := ioutil.ReadFile("words.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		dict = strings.Split(string(content), "\n")
	}
	rand.Seed(time.Now().UnixNano())
	answer = dict[rand.Intn(len(dict))]
	fmt.Println(answer)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}